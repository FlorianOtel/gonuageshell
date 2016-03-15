package nuage

import (
	log "github.com/FlorianOtel/gonuageshell/Godeps/_workspace/src/github.com/Sirupsen/logrus"

	"bytes"
	"crypto/tls"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

////////
//////// API version independent operations. Handle objects as []byte. Up to the callers to JSON encode / decode those  in an API version specific format.
//////// These are low-level functions, to be consumed by / via version sepcific wrappers in other  packages (e.g. "nuage_v3_2", "nuage_v4_0", etc)
////////

//// Generic "GET" entity. Endpoint format (ensured by the caller):
// <entity>
// <entity> <ID>
// <entity> <ID> <children>

func GetEntity(c *Connection, endpoint string) ([]byte, error) {

	// if len(args) < 1 || len(args) > 3 {
	// 	// log.Debugf("Nuage generic GET: len(args): %d", len(args))
	// 	log.Debugf("Nuage GET entity: Malformed requst, invalid entity: %s", strings.Join(args, "/"))
	// 	err := fmt.Errorf("Nuage GET entity: Malformed requst, invalid entity: %s", strings.Join(args, "/"))
	// 	return nil, err
	// }

	reply, statuscode, err := nuagetransaction(c, "GET", c.Url+"/nuage/api/"+c.Apivers+"/"+endpoint, []byte(""))

	if err != nil {
		log.Debugf("Nuage GET entity: Error: %s", err)
		return nil, err
	}

	if statuscode != 200 {
		log.Debugf("Nuage GET entity: HTTP status code: %d", statuscode)
		err = fmt.Errorf("HTTP status code: %d", statuscode)
		return nil, err
	}

	return reply, nil

}

// Create Entity. Up to the caller to encode it as a valid "payload []byte" and select an appropriate API entity -- e.g. "enterprises"
func CreateEntity(c *Connection, entity string, payload []byte) ([]byte, error) {
	reply, statuscode, err := nuagetransaction(c, "POST", c.Url+"/nuage/api/"+c.Apivers+"/"+entity, payload)

	if err != nil {
		log.Debugf("Nuage CREATE entity: Unable to create entity. Error: %s", err)
		return nil, err
	}

	if statuscode != 201 {
		log.Debugf("Nuage CREATE entity: Unable to create entity. HTTP status code: %d", statuscode)
		err = fmt.Errorf("HTTP status code: %d", statuscode)
		return nil, err
	}

	return reply, nil
}

// Delete Entity. Up to the caller to provide a correct ID for the given API entity -- e.g. "enterprises"
func DeleteEntity(c *Connection, entity string, id string) ([]byte, error) {
	reply, statuscode, err := nuagetransaction(c, "DELETE", c.Url+"/nuage/api/"+c.Apivers+"/"+entity+"/"+id, []byte(""))

Reeval:
	if err != nil {
		log.Debugf("Nuage DELETE: Unable to delete: %s with ID: %s", entity, id)
		return nil, err
	}

	log.Debugf("Nuage DELETE: Assessing HTTP status code: %d", statuscode)
	switch statuscode {
	case 300: // Used for Enterprise delete, must confirm deletion
		// XXX -- This works when "entity" is "enterprises"
		// XXX -- Check if there are other delete methods (i.e. "entity") for which the HTTP status code is "300"
		reply, statuscode, err = nuagetransaction(c, "DELETE", c.Url+"/nuage/api/"+c.Apivers+"/"+entity+"/"+id+"/?responseChoice=1", []byte(""))
		// The reply from this should be a "204" with No content. Need to check again.
		goto Reeval
	case 204: // Deleted
		return reply, nil
	default:
		log.Debugf("Nuage DELETE: Unable to delete: %s with ID: %s. HTTP status code: %d", entity, id, statuscode)
		err = fmt.Errorf("HTTP status code: %d", statuscode)
		return nil, err

	}
	return reply, err
}

////////
//////// Auxiliary functions
////////

// Return string representation of a Nuage API connection
func (c *Connection) String() string {
	// str := fmt.Sprint("Nuage API connection:\n")
	str := fmt.Sprintf("\n    Endpoint URL: [%s]\n", c.Url)
	str = str + fmt.Sprintf("    API version: [%s]\n", c.Apivers)

	if c.token != nil {
		str = str + fmt.Sprintf("    Connection established as User: [%s], Enterprise: [%s] \n", c.token.UserName, c.token.EnterpriseName)
	} else {
		str = str + fmt.Sprint("    Not connected\n")
	}

	return str
}

// Initialize Nuage API connection using username & password. Stores a valid Authtoken upon success.
func (c *Connection) Connect(org, user, pass string) error {
	// var err error

	var auth []Authtoken

	// log.Debugf("Base64 encoding of %s is: %s", user+":"+pass, base64.URLEncoding.EncodeToString([]byte(user+":"+pass)))

	// Get APIkey + its timestamp from "/nuage/api/v1_0/me"
	req, err := http.NewRequest("GET", c.Url+"/nuage/api/v1_0/me", nil)
	req.Header.Set("X-Nuage-Organization", org)
	req.Header.Set("Authorization", "XREST "+base64.URLEncoding.EncodeToString([]byte(user+":"+pass)))
	req.Header.Set("Content-Type", "application/json")

	// Skip TLS security check
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	log.Debugf("Attempting to make connection to: %s", c.Url+"/nuage/api/v1_0/me")
	resp, err := client.Do(req)

	if err != nil {
		return err
	}

	defer resp.Body.Close()

	log.Debugf("Response Status: %s", resp.Status)
	log.Debugf("Response Headers: %s", resp.Header)

	if resp.StatusCode != 200 {
		log.Debugf("VSD authentication to ["+c.Url+"/nuage/api/v1_0/me"+"] failed with status: %s", resp.Status)
		err = fmt.Errorf("HTTP status: %s", resp.Status)
		return err
	}

	// body, err := ioutil.ReadAll(resp.Body)
	// fmt.Println("===> response Body:", string(body))

	err = json.NewDecoder(resp.Body).Decode(&auth)

	if err != nil {
		log.Debugf("Unable to decode JSON payload: %s", err.Error())
		return err
	}

	// json.Unmarshal(body, &auth)

	log.Debugf("Response body: %#v \n", auth[0])

	// Keep a pointer to this connection so we can reuse the credentials
	c.token = &auth[0]

	return nil
}

// Basic Nuage API transaction. Returns the response body (empty), HTTP response code and any errors. Up to the caller to check HTTP error codes. Unexported.
func nuagetransaction(c *Connection, method string, url string, jsonpayload []byte) ([]byte, int, error) {
	if c.token == nil {
		log.Debugf("Invalid connection: %s", c)
		return []byte(""), -1, errors.New("Invalid Nuage API connection")
	}
	// Still TBD: Additional sanity checks for the connection e.g. is token still valid ?

	req, err := http.NewRequest(method, url, nil)
	req.Header.Set("X-Nuage-Organization", c.token.EnterpriseName)
	req.Header.Set("Authorization", "XREST "+base64.URLEncoding.EncodeToString([]byte(c.token.UserName+":"+c.token.Apikey)))
	req.Header.Set("Content-Type", "application/json")

	// "POST" methods require a valid payload.
	if method == "POST" && len(jsonpayload) != 0 {
		// If we are passed a payload, encode that
		req.Body = ioutil.NopCloser(bytes.NewBuffer(jsonpayload))
		// log.Debugf("Request payload: %s", string(jsonpayload))
		defer req.Body.Close()
	}

	// Skip TLS security check
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client := &http.Client{Transport: tr}

	log.Debugf("Nuage API connection: %s to/from: %s with payload: %s", method, url, string(jsonpayload))
	resp, err := client.Do(req)
	if err != nil {
		return []byte(""), -1, err
	}

	log.Debugf("Response Status: %s", resp.Status)
	log.Debugf("Response Headers: %s", resp.Header)

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	log.Debugf("Response Body: %s", string(body))

	return body, resp.StatusCode, nil
}
