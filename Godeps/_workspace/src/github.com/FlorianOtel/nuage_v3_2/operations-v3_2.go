package nuage_v3_2

import (
	"encoding/json"
	"fmt"

	// "reflect"

	nuage "github.com/FlorianOtel/gonuageshell/Godeps/_workspace/src/github.com/FlorianOtel/nuage"

	log "github.com/FlorianOtel/gonuageshell/Godeps/_workspace/src/github.com/Sirupsen/logrus"
)

////////
//////// VirtualMachine methods and operations. OBS: Some methods have other entities as method receivers (e.g. Subnet / Domain / ... VMInterface List).
////////

// Caller MUST initialize:
// - UUID (as e.g. returned upon "ovs-appctl vm/send-event define <XML definition file>" )
// - Name

// Caller SHOULD initialize at least one interface
// - vm.Interfaces[0].MAC  -- MAC addr of the first interface
// - vm.Interfaces[0].VPortID  -- the ID of the VPort where this interface should be connected

// VirtualMachine Create
func (vm *VirtualMachine) Create(c *nuage.Connection) error {
	if vm == nil {
		err := fmt.Errorf("VirtualMachine Create: Empty method receiver, nothing to do")
		return err
	}

	if vm.Name == "" {
		err := fmt.Errorf("VirtualMachine Create: Empty Name, nothing to do")
		return err
	}

	if vm.UUID == "" {
		err := fmt.Errorf("VirtualMachine Create: Empty UUID, nothing to do")
		return err
	}

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var vma [1]VirtualMachine
	// Copies the method receiver & all initialized fields
	vma[0] = *vm

	jsonvm, _ := json.MarshalIndent(vma[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "vms", jsonvm)

	if err != nil {
		log.Debugf("VirtualMachine Create: Unable to create VirtualMachine with name: [%s] . Error: %s ", vm.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &vma)

	if err != nil {
		log.Debugf("VirtualMachine Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*vm = vma[0]
	log.Debugf("VirtualMachine Create: Created VirtualMachine with ID: [%s]", vm.ID)
	return nil
}

// VirtualMachine Delete.  Caller must initialize the VirtualMachine ID (vp.ID)
func (vm *VirtualMachine) Delete(c *nuage.Connection) error {

	if vm.ID == "" {
		err := fmt.Errorf("VirtualMachine Delete: Empty VirtualMachine ID, nothing to do")
		return err
	}

	_, err := nuage.DeleteEntity(c, "vms", vm.ID)

	if err != nil {
		log.Debugf("VirtualMachine Delete: Unable to delete VirtualMachine with ID: [%s] . Error: %s ", vm.ID, err)
		return err
	}

	log.Debugf("VirtualMachine Delete: Deleted VirtualMachine with ID: [%s]", vm.ID)
	return nil

}

// Virtual Machhine  Get.  Caller must initialize the VirtualMachine ID (vm.ID)
func (vm *VirtualMachine) Get(c *nuage.Connection) error {

	if vm.ID == "" {
		err := fmt.Errorf("VirtualMachine Get: Empty VirtualMachine ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "vms/"+vm.ID)

	if err != nil {
		log.Debugf("VirtualMachine Get. Unable to find VirtualMachine with ID: [%s] . Error: %s ", vm.ID, err)
		return err
	}

	var vma [1]VirtualMachine

	err = json.Unmarshal(reply, &vma)
	if err != nil {
		log.Debugf("VirtualMachine Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*vm = vma[0]
	log.Debugf("VirtualMachine Get: Found VirtualMachine with Name: [%s] and ID: [%s]", vm.Name, vm.ID)
	return nil

}

// Global list (all VirtualMachines in the Data Center)
func (vms *VirtualMachineslice) List(c *nuage.Connection) error {

	reply, err := nuage.GetEntity(c, "vms")

	if err != nil {
		log.Debugf("VirtualMachine List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("VirtualMachine List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, vms)

	if err != nil {
		log.Debugf("VirtualMachine List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("VirtualMachine List: done")
	return nil
}

////////
//////// VMInterface operations. OBS: Some are methods for other entities (e.g. Subnet / Domain / ... VMInterface List).
////////

// VMInterfaces list for a Domain.  Caller must initialize the Domain ID (d.ID)
func (d *Domain) VMInterfacesList(c *nuage.Connection) ([]VMInterface, error) {

	if d.ID == "" {
		err := fmt.Errorf("Domain VMInterfaces List: Empty Domain ID, nothing to do")
		return nil, err
	}

	reply, err := nuage.GetEntity(c, "domains/"+d.ID+"/vminterfaces")

	if err != nil {
		log.Debugf("Domain VMInterfaces List: Error %s ", err)
		return nil, err
	}

	if len(reply) == 0 {
		log.Debugf("Domain VMInterfaces List: Empty list")
		return nil, nil
	}

	var vmis []VMInterface

	err = json.Unmarshal(reply, &vmis)
	if err != nil {
		log.Debugf("Domain VMInterfaces List:  Unable to decode JSON payload: %s ", err)
		return nil, err
	}

	log.Debug("Domain VMInterfaces List: done")
	return vmis, nil

}

// VMInterfaces list for a Subnet.  Caller must initialize the Subnet ID (s.ID)
func (s *Subnet) VMInterfacesList(c *nuage.Connection) ([]VMInterface, error) {

	if s.ID == "" {
		err := fmt.Errorf("Subnet VMInterfaces List: Empty Subnet ID, nothing to do")
		return nil, err
	}

	reply, err := nuage.GetEntity(c, "subnets/"+s.ID+"/vminterfaces")

	if err != nil {
		log.Debugf("Subnet VMInterfaces List: Error %s ", err)
		return nil, err
	}

	if len(reply) == 0 {
		log.Debugf("Subnet VMInterfaces List: Empty list")
		return nil, nil
	}

	var vmis []VMInterface

	err = json.Unmarshal(reply, &vmis)
	if err != nil {
		log.Debugf("Subnet VMInterfaces List:  Unable to decode JSON payload: %s ", err)
		return nil, err
	}

	log.Debug("Subnet VMInterfaces List: done")
	return vmis, nil

}

// VMInterface Delete.  Caller must initialize the VMInterface ID (vmi.ID)
func (vmi *VMInterface) Delete(c *nuage.Connection) error {

	if vmi.ID == "" {
		err := fmt.Errorf("VMInterface Delete: Empty VMInterface ID, nothing to do")
		return err
	}

	_, err := nuage.DeleteEntity(c, "vminterfaces", vmi.ID)

	if err != nil {
		log.Debugf("VMInterface Delete: Unable to delete VMInterface with ID: [%s] . Error: %s ", vmi.ID, err)
		return err
	}

	log.Debugf("VMInterface Delete: Deleted VMInterface with ID: [%s]", vmi.ID)
	return nil

}

// VMInterface Get.  Caller must initialize the VMInterface ID (vm.ID)
func (vmi *VMInterface) Get(c *nuage.Connection) error {

	if vmi.ID == "" {
		err := fmt.Errorf("VMInterface Get: Empty VMInterface ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "vminterfaces/"+vmi.ID)

	if err != nil {
		log.Debugf("VMInterface Get. Unable to find VMInterface with ID: [%s] . Error: %s ", vmi.ID, err)
		return err
	}

	var vmia [1]VMInterface

	err = json.Unmarshal(reply, &vmia)
	if err != nil {
		log.Debugf("VMInterface Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*vmi = vmia[0]
	log.Debugf("VMInterface Get: Found VMInterface with Name: [%s] and ID: [%s]", vmi.Name, vmi.ID)
	return nil

}

// Global list (all VMInterfaces in the Data Center)
func (vmis *VMInterfaceslice) List(c *nuage.Connection) error {

	reply, err := nuage.GetEntity(c, "vminterfaces")

	if err != nil {
		log.Debugf("VMInterface List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("VMInterface List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, vmis)

	if err != nil {
		log.Debugf("VMInterface List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("VMInterface List: done")
	return nil
}

////////
//////// VPort operations. OBS: Some are methods for other entities (e.g. Subnet / Domain / ... VPort List).
////////

// Add VPort to Subnet / Create VPort. Caller must initialize the Subnet ID (s.ID). Additionally, the virtual Port must have:
// - A Name (vp.Name)
// - A Type (vp.Type)
func (s *Subnet) AddVPort(c *nuage.Connection, vp VPort) (VPort, error) {
	var vpa [1]VPort

	// In the worst case we return what we received
	vpa[0] = vp
	if s.ID == "" {
		err := fmt.Errorf("Subnet Add VPort: Empty Subnet ID, nothing to do")
		return vpa[0], err
	}

	//XXX -- TBD: Better sanity checks
	if vp.Name == "" || vp.Type == "" || vp.AddressSpoofing == "" {
		err := fmt.Errorf("Subnet Add VPort: Invalid VPort initialization, nothing to do")
		return vpa[0], err
	}

	jsonvport, _ := json.MarshalIndent(vp, "", "\t")
	reply, err := nuage.CreateEntity(c, "subnets/"+s.ID+"/vports", jsonvport)

	if err != nil {
		log.Debugf("Subnet Add VPort: Error: %s ", err)
		return vpa[0], err
	}

	err = json.Unmarshal(reply, &vpa)
	if err != nil {
		log.Debugf("Subnet Add VPort:  Unable to decode JSON payload: %s ", err)
		return vpa[0], err
	}
	log.Debug("Subnet Add VPort: done")
	return vpa[0], nil

}

// VPorts list for a Domain.  Caller must initialize the Domain ID (s.ID)
func (d *Domain) VPortsList(c *nuage.Connection) ([]VPort, error) {

	if d.ID == "" {
		err := fmt.Errorf("Domain VPorts List: Empty Domain ID, nothing to do")
		return nil, err
	}

	reply, err := nuage.GetEntity(c, "domains/"+d.ID+"/vports")

	if err != nil {
		log.Debugf("Domain VPorts List: Error %s ", err)
		return nil, err
	}

	if len(reply) == 0 {
		log.Debugf("Domain VPorts List: Empty list")
		return nil, nil
	}

	var vports []VPort

	err = json.Unmarshal(reply, &vports)
	if err != nil {
		log.Debugf("Domain VPorts List:  Unable to decode JSON payload: %s ", err)
		return nil, err
	}
	log.Debug("Domain VPorts List: done")
	return vports, nil

}

// VPort list for a Subnet.  Caller must initialize the Subnet ID (s.ID)
func (s *Subnet) VPortsList(c *nuage.Connection) ([]VPort, error) {

	if s.ID == "" {
		err := fmt.Errorf("Subnet VPorts List: Empty Subnet ID, nothing to do")
		return nil, err
	}

	reply, err := nuage.GetEntity(c, "subnets/"+s.ID+"/vports")

	if err != nil {
		log.Debugf("Subnet VPorts List: Error %s ", err)
		return nil, err
	}

	if len(reply) == 0 {
		log.Debugf("Subnet VPorts List: Empty list")
		return nil, nil
	}

	var vports []VPort

	err = json.Unmarshal(reply, &vports)
	if err != nil {
		log.Debugf("Subnet VPorts List:  Unable to decode JSON payload: %s ", err)
		return nil, err
	}

	log.Debug("Subnet VPorts List: done")
	return vports, nil

}

// VPort Delete.  Caller must initialize the VPort ID (vp.ID)
func (vp *VPort) Delete(c *nuage.Connection) error {

	if vp.ID == "" {
		err := fmt.Errorf("VPort Delete: Empty VPort ID, nothing to do")
		return err
	}

	_, err := nuage.DeleteEntity(c, "vports", vp.ID)

	if err != nil {
		log.Debugf("VPort Delete: Unable to delete VPort with ID: [%s] . Error: %s ", vp.ID, err)
		return err
	}

	log.Debugf("VPort Delete: Deleted VPort with ID: [%s]", vp.ID)
	return nil

}

// VPort Get.  Caller must initialize the VPort ID (vp.ID)
func (vp *VPort) Get(c *nuage.Connection) error {

	if vp.ID == "" {
		err := fmt.Errorf("VPort Get: Empty VPort ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "vports/"+vp.ID)

	if err != nil {
		log.Debugf("VPort Get: Error %s ", err)
		return err
	}

	var vpa [1]VPort

	err = json.Unmarshal(reply, &vpa)
	if err != nil {
		log.Debugf("VPort Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*vp = vpa[0]
	log.Debugf("VPort Get: Found VPort with Name: [%s] and ID: [%s]", vp.Name, vp.ID)
	return nil

}

////////
//////// Subnet methods
////////

// Caller must populate Subnet ID (s.ID)
func (s *Subnet) Delete(c *nuage.Connection) error {
	if s == nil {
		err := fmt.Errorf("Subnet Delete: Empty method receiver, nothing to do")
		return err
	}

	if s.ID == "" {
		err := fmt.Errorf("Subnet Delete: Empty ID, nothing to do")
		return err
	}
	_, err := nuage.DeleteEntity(c, "subnets", s.ID)

	if err != nil {
		log.Debugf("Subnet Delete: Unable to delete Subnet with ID: [%s] . Error: %s ", s.ID, err)
		return err
	}

	log.Debugf("Subnet Delete: Deleted subnet with ID: [%s] ", s.ID)
	return nil
}

// Assumes the method receiver was allocated using "new(Subnet)"
// Caller must populate:
// - Name (s.Name)
// - Parent Zone ID (s.ParentID)
// - Subnet Tempate ID (s.TemplateID)
// - Address (s.Address) -- e.g. "10.24.24.0"
// - Netmask (s.Netmask) -- e.g. "255.255.255.0"
// - Optionally:  Subnet Template ID (s.TemplateID)
func (s *Subnet) Create(c *nuage.Connection) error {
	if s == nil {
		err := fmt.Errorf("Subnet Create: Empty method receiver, nothing to do")
		return err
	}

	if s.Name == "" {
		err := fmt.Errorf("Subnet Create: Empty Name, nothing to do")
		return err
	}

	if s.ParentID == "" {
		err := fmt.Errorf("Subnet Create: Empty ParentID, nothing to do")
		return err
	}

	if s.TemplateID == "" && (s.Address == "" || s.Netmask == "") {
		err := fmt.Errorf("Subnet Create: Need either Subnet Template ID or Subnet Address & Netmask. Nothing to do")
		return err
	}

	// XXX - Do not insist on it being present
	// if s.TemplateID == "" {
	// 	err := fmt.Errorf("Subnet Create: Empty subnet template ID, nothing to do")
	// 	return err
	// }

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var subneta [1]Subnet
	// XXX - This copies the supplied Name, ParentID and TemplateID
	subneta[0] = *s

	jsonsubnet, _ := json.MarshalIndent(subneta[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "zones/"+s.ParentID+"/subnets", jsonsubnet)
	if err != nil {
		log.Debugf("Subnet Create: Unable to create Subnet with name: [%s] . Error: %s ", s.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &subneta)

	if err != nil {
		log.Debugf("Subnet Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*s = subneta[0]
	log.Debugf("Subnet Create: Created Subnet with ID: [%s]", s.ID)
	return nil
}

// Get by Subnet ID (s.ID)
func (s *Subnet) Get(c *nuage.Connection) error {
	if s.ID == "" {
		err := fmt.Errorf("Subnet Get: Empty ID, nothing to do")
		return err
	}
	reply, err := nuage.GetEntity(c, "subnets/"+s.ID)

	if err != nil {
		log.Debugf("Subnet Get: Unable to get subnet with ID: [%s]. Error: %s ", s.ID, err)
		return err
	}

	var subneta [1]Subnet
	err = json.Unmarshal(reply, &subneta)
	if err != nil {
		log.Debugf("Subnet Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*s = subneta[0]
	log.Debugf("Subnet Get: Found Subnet with Name: [%s] and ID: [%s]", s.Name, s.ID)
	return nil
}

func (ss *Subnetslice) List(c *nuage.Connection, parentid string) error {
	var reply []byte
	var err error

	if parentid == "" { // get global list of subnets
		reply, err = nuage.GetEntity(c, "subnets")
	} else {
		// get the list of subnets for a given Zone ID
		reply, err = nuage.GetEntity(c, "zones/"+parentid+"/subnets")
	}

	if err != nil {
		log.Debugf("Subnet List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Subnet List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, ss)

	if err != nil {
		log.Debugf("Subnet List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("Subnet List: done")
	return nil
}

////////
//////// Zone methods
////////

// Caller must populate Zone ID (z.ID)
func (z *Zone) Delete(c *nuage.Connection) error {
	if z == nil {
		err := fmt.Errorf("Zone Delete: Empty method receiver, nothing to do")
		return err
	}

	if z.ID == "" {
		err := fmt.Errorf("Zone Delete: Empty ID, nothing to do")
		return err
	}
	_, err := nuage.DeleteEntity(c, "zones", z.ID)

	if err != nil {
		log.Debugf("Zone Delete: Unable to delete Zone with ID: [%s] . Error: %s ", z.ID, err)
		return err
	}

	log.Debugf("Zone Delete: Deleted zone with ID: [%s] ", z.ID)
	return nil
}

// Assumes the method receiver was allocated using "new(Zone)"
// Caller must populate:
// - Name (z.Name)
// - Parent Domain ID (z.ParentID)
// - Optionally:  Zone Template ID (z.TemplateID)
func (z *Zone) Create(c *nuage.Connection) error {
	if z == nil {
		err := fmt.Errorf("Zone Create: Empty method receiver, nothing to do")
		return err
	}

	if z.Name == "" {
		err := fmt.Errorf("Zone Create: Empty Name, nothing to do")
		return err
	}

	if z.ParentID == "" {
		err := fmt.Errorf("Zone Create: Empty ParentID, nothing to do")
		return err
	}

	// XXX - Do not insist on it being present
	// if z.TemplateID == "" {
	// 	err := fmt.Errorf("Zone Create: Empty zone template ID, nothing to do")
	// 	return err
	// }

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var za [1]Zone
	// XXX - This copies the supplied Name, ParentID and TemplateID
	za[0] = *z

	jsonzone, _ := json.MarshalIndent(za[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "domains/"+z.ParentID+"/zones", jsonzone)

	if err != nil {
		log.Debugf("Zone Create: Unable to create Zone with name: [%s] . Error: %s ", z.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &za)

	if err != nil {
		log.Debugf("Zone Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*z = za[0]
	log.Debugf("Zone Create: Created Zone with ID: [%s]", z.ID)
	return nil
}

// Get by Zone ID (z.ID)
func (z *Zone) Get(c *nuage.Connection) error {
	if z.ID == "" {
		err := fmt.Errorf("Zone template Get: Empty ID, nothing to do")
		return err
	}
	reply, err := nuage.GetEntity(c, "zones/"+z.ID)

	if err != nil {
		log.Debugf("Zone Get: Unable to get domain with ID: [%s]. Error: %s ", z.ID, err)
		return err
	}

	var za [1]Zone
	err = json.Unmarshal(reply, &za)
	if err != nil {
		log.Debugf("Zone Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*z = za[0]
	log.Debugf("Zone Get: Found Zone with Name: [%s] and ID: [%s]", z.Name, z.ID)
	return nil
}

func (zs *Zoneslice) List(c *nuage.Connection, parentid string) error {
	var reply []byte
	var err error

	if parentid == "" { // get global list of zones
		reply, err = nuage.GetEntity(c, "zones")
	} else {
		// get the list of domains for a given domain ID
		reply, err = nuage.GetEntity(c, "domains/"+parentid+"/zones")
	}

	if err != nil {
		log.Debugf("Zone List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Zone List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, zs)

	if err != nil {
		log.Debugf("Zone List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("Zone List: done")
	return nil
}

////////
//////// Zone template  methods
////////

// Assumes the method receiver was allocated using "new(Zonetemplate)"
// Caller must populate the ID (zt.ID)
func (zt *Zonetemplate) Delete(c *nuage.Connection) error {
	if zt == nil {
		err := fmt.Errorf("Zone template Delete: Empty method receiver, nothing to do")
		return err
	}

	if zt.ID == "" {
		err := fmt.Errorf("Zone template Delete: Empty ID, nothing to do")
		return err
	}
	_, err := nuage.DeleteEntity(c, "zonetemplates", zt.ID)

	if err != nil {
		log.Debugf("Zone template Delete: Unable to delete Zone template with ID: [%s] . Error: %s ", zt.ID, err)
		return err
	}

	log.Debugf("Zone template Delete: Deleted zone template with ID: [%s] ", zt.ID)
	return nil
}

// Assumes the method receiver was allocated using "new(Zonetemplate)"
// Caller must populate Name (zt.Name) and ParentID (zt.ParentID)
func (zt *Zonetemplate) Create(c *nuage.Connection) error {
	if zt == nil {
		err := fmt.Errorf("Zone template Create: Empty method receiver, nothing to do")
		return err
	}

	if zt.Name == "" {
		err := fmt.Errorf("Zone template Create: Empty Name, nothing to do")
		return err
	}

	if zt.ParentID == "" {
		err := fmt.Errorf("Zone template Create: Empty ParentID, nothing to do")
		return err
	}

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var zta [1]Zonetemplate

	// XXX - This copies the supplied Name and ParentID
	zta[0] = *zt

	jsonzt, _ := json.MarshalIndent(zta[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "domaintemplates/"+zt.ParentID+"/zonetemplates", jsonzt)

	if err != nil {
		log.Debugf("Zone template Create: Unable to create Zone template with name: [%s] . Error: %s ", zt.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &zta)

	if err != nil {
		log.Debugf("Zone template Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*zt = zta[0]
	log.Debugf("Zone template Create: Created Zone template with ID: [%s]", zt.ID)
	return nil
}

// GET by ID (zt.ID)
func (zt *Zonetemplate) Get(c *nuage.Connection) error {
	if zt.ID == "" {
		err := fmt.Errorf("Zone template Get: Empty ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "zonetemplates/"+zt.ID)

	if err != nil {
		log.Debugf("Zone template Get: Unable to get zone template with ID: [%s]. Error: %s ", zt.ID, err)
		return err
	}

	var zta [1]Zonetemplate
	err = json.Unmarshal(reply, &zta)
	if err != nil {
		log.Debugf("Zone template Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*zt = zta[0]
	log.Debugf("Zone template Get: Found Zone template with Name: [%s] and ID: [%s]", zt.Name, zt.ID)
	return nil
}

func (zts *Zonetemplateslice) List(c *nuage.Connection, parentid string) error {
	if parentid == "" {
		err := fmt.Errorf("Zone template List: Empty ParentID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "domaintemplates/"+parentid+"/zonetemplates")

	if err != nil {
		log.Debugf("Zone templates List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Zone templates List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, zts)

	if err != nil {
		log.Debugf("Zone template List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("Zone template List: done")
	return nil
}

////////
//////// Domain methods
////////

// Caller must populate domain ID (d.ID)
func (d *Domain) Delete(c *nuage.Connection) error {
	if d == nil {
		err := fmt.Errorf("Domain Delete: Empty method receiver, nothing to do")
		return err
	}

	if d.ID == "" {
		err := fmt.Errorf("Domain Delete: Empty ID, nothing to do")
		return err
	}
	_, err := nuage.DeleteEntity(c, "domains", d.ID)

	if err != nil {
		log.Debugf("Domain Delete: Unable to delete Domain with ID: [%s] . Error: %s ", d.ID, err)
		return err
	}

	log.Debugf("Domain Delete: Deleted domain with ID: [%s] ", d.ID)
	return nil
}

// Assumes the method receiver was allocated using "new(Domain)"
// Caller must populate:
// - Name (d.Name)
// - Parent Enterprise ID (d.ParentID)
// - Domain Template ID (d.TemplateID)
func (d *Domain) Create(c *nuage.Connection) error {
	if d == nil {
		err := fmt.Errorf("Domain Create: Empty method receiver, nothing to do")
		return err
	}

	if d.Name == "" {
		err := fmt.Errorf("Domain Create: Empty Name, nothing to do")
		return err
	}

	if d.ParentID == "" {
		err := fmt.Errorf("Domain Create: Empty ParentID, nothing to do")
		return err
	}

	if d.TemplateID == "" {
		err := fmt.Errorf("Domain Create: Empty domain template ID, nothing to do")
		return err
	}

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var da [1]Domain
	// XXX - This copies the supplied Name, ParentID and TemplateID
	da[0] = *d

	jsondomain, _ := json.MarshalIndent(da[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "enterprises/"+d.ParentID+"/domains", jsondomain)

	if err != nil {
		log.Debugf("Domain Create: Unable to create Domain with name: [%s] . Error: %s ", d.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &da)

	if err != nil {
		log.Debugf("Domain Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*d = da[0]
	log.Debugf("Domain Create: Created Domain with ID: [%s]", d.ID)
	return nil
}

// Get by Domain ID (d.ID)
func (d *Domain) Get(c *nuage.Connection) error {
	if d.ID == "" {
		err := fmt.Errorf("Domain template Get: Empty ID, nothing to do")
		return err
	}
	reply, err := nuage.GetEntity(c, "domains/"+d.ID)

	if err != nil {
		log.Debugf("Domain Get: Unable to get domain with ID: [%s] . Error: %s ", d.ID, err)
		return err
	}

	var da [1]Domain
	err = json.Unmarshal(reply, &da)
	if err != nil {
		log.Debugf("Domain Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*d = da[0]
	log.Debugf("Domain Get: Found Domain with Name: [%s] and ID: [%s]", d.Name, d.ID)
	return nil
}

func (ds *Domainslice) List(c *nuage.Connection, parentid string) error {
	var reply []byte
	var err error

	if parentid == "" { // get global list of domains
		reply, err = nuage.GetEntity(c, "domains")
	} else {
		// get the list of domains for a given enterprise ID
		reply, err = nuage.GetEntity(c, "enterprises/"+parentid+"/domains")
	}

	if err != nil {
		log.Debugf("Domain List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Domain List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, ds)

	if err != nil {
		log.Debugf("Domain List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("Domain List: done")
	return nil
}

////////
//////// Domaintemplate methods
////////

// Assumes the method receiver was allocated using "new(Domaintemplate)"
// Caller must populate the ID (dt.ID)
func (dt *Domaintemplate) Delete(c *nuage.Connection) error {
	if dt == nil {
		err := fmt.Errorf("Domain template Delete: Empty method receiver, nothing to do")
		return err
	}

	if dt.ID == "" {
		err := fmt.Errorf("Domain template Delete: Empty ID, nothing to do")
		return err
	}
	_, err := nuage.DeleteEntity(c, "domaintemplates", dt.ID)

	if err != nil {
		log.Debugf("Domain template Delete: Unable to delete Domain template with ID: [%s] . Error: %s ", dt.ID, err)
		return err
	}

	log.Debugf("Domain template Delete: Deleted domain template with ID: [%s] ", dt.ID)
	return nil
}

// Assumes the method receiver was allocated using "new(Domaintemplate)"
// Caller must populate Name (dt.Name) and ParentID (dt.ParentID)
func (dt *Domaintemplate) Create(c *nuage.Connection) error {
	if dt == nil {
		err := fmt.Errorf("Domain template Create: Empty method receiver, nothing to do")
		return err
	}

	if dt.Name == "" {
		err := fmt.Errorf("Domain template Create: Empty Name, nothing to do")
		return err
	}

	if dt.ParentID == "" {
		err := fmt.Errorf("Domain template Create: Empty ParentID, nothing to do")
		return err
	}

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var dta [1]Domaintemplate

	// XXX - This copies the supplied Name and ParentID
	dta[0] = *dt

	jsondt, _ := json.MarshalIndent(dta[0], "", "\t")
	reply, err := nuage.CreateEntity(c, "enterprises/"+dt.ParentID+"/domaintemplates", jsondt)

	if err != nil {
		log.Debugf("Domain template Create: Unable to create Domain template with name: [%s] . Error: %s ", dt.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &dta)

	if err != nil {
		log.Debugf("Domain template Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*dt = dta[0]
	log.Debugf("Domain template Create: Created Domain template with ID: [%s]", dt.ID)
	return nil
}

// GET by ID (dt.ID)
func (dt *Domaintemplate) Get(c *nuage.Connection) error {
	if dt.ID == "" {
		err := fmt.Errorf("Domain template Get: Empty ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "domaintemplates/"+dt.ID)

	if err != nil {
		log.Debugf("Domain template Get: Unable to get domain template with ID: [%s] . Error: %s ", dt.ID, err)
		return err
	}

	var dta [1]Domaintemplate
	err = json.Unmarshal(reply, &dta)
	if err != nil {
		log.Debugf("Domain template Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*dt = dta[0]
	log.Debugf("Domain template Get: Found Domain template with Name: [%s] and ID: [%s]", dt.Name, dt.ID)
	return nil
}

func (dts *Domaintemplateslice) List(c *nuage.Connection, parentid string) error {
	if parentid == "" {
		err := fmt.Errorf("Domain template List: Empty ParentID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "enterprises/"+parentid+"/domaintemplates")

	if err != nil {
		log.Debugf("Domain templates List: Unable to obtain list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Domain templates List: Empty list")
		return nil
	}

	err = json.Unmarshal(reply, dts)

	if err != nil {
		log.Debugf("Domain template List: Unable to decode JSON payload: %s ", err)
		return err
	}
	log.Debug("Domain template List: done")
	return nil
}

////////
//////// Enterprise methods
////////

// Must have a valid ID (org.ID)
func (org *Enterprise) Delete(c *nuage.Connection) error {
	if org == nil {
		err := fmt.Errorf("Enterprise Delete: Empty method receiver, nothing to do")
		return err
	}

	if org.ID == "" {
		err := fmt.Errorf("Enterprise Delete: Empty Enterprise ID, nothing to do")
		return err
	}

	_, err := nuage.DeleteEntity(c, "enterprises", org.ID)

	if err != nil {
		log.Debugf("Enterprise Delete: Unable to delete Enterprise with name %s . Error: %s ", org.Name, err)
		return err
	}

	log.Debugf("Enterprise Delete: Deleted Enterprise [%s] with ID: [%s] ", org.Name, org.ID)
	return nil

}

// Assumes that the method receiver was allocated using "new(Enterprise)", initialized accordingly (name + description).
func (org *Enterprise) Create(c *nuage.Connection) error {
	if org == nil {
		err := fmt.Errorf("Enterprise Create: Empty method receiver, nothing to do")
		return err
	}

	// Check that the method receiver was allocated properly
	// Disabled. Overkill ?

	// if reflect.TypeOf(*org).String() != "nuage_v3_2.Enterprise" {
	// 	err := fmt.Errorf("Enterprise Create: Invalid method receiver type")
	// 	return err
	// }

	// It has to be an array since the reply from the server is as an array of JSON objects, and we use it for decoding as well
	var orga [1]Enterprise
	orga[0] = *org

	if org.Description == "" {
		// Default Enterpise Description unless one is specified
		orga[0].Description = "Created by Golang API driver"
	} else {
		orga[0].Description = org.Description
	}

	jsonorg, _ := json.MarshalIndent(orga[0], "", "\t")

	// Quick and dirty alternative: Just build a JSON object with "name" and "description" fields
	// jsonorg := "      {\"name\":\"" + name + "\",\"description\":\"Created by Golang API client\"}      "

	reply, err := nuage.CreateEntity(c, "enterprises", jsonorg)

	if err != nil {
		log.Debugf("Enterprise Create: Unable to create Enterprise with name: %s . Error: %s ", org.Name, err)
		return err
	}

	err = json.Unmarshal(reply, &orga)

	if err != nil {
		log.Debugf("Enterprise Create: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX -- Mutate the method receiver
	*org = orga[0]
	log.Debugf("Enterprise Create: Created Enterprise with Name: [%s] and ID: [%s]", org.Name, org.ID)
	return nil

}

// GET enterprise by ID (org.ID)
func (org *Enterprise) Get(c *nuage.Connection) error {
	if org.ID == "" {
		err := fmt.Errorf("Enterprise Get: Empty ID, nothing to do")
		return err
	}

	reply, err := nuage.GetEntity(c, "enterprises/"+org.ID)

	if err != nil {
		log.Debugf("Enterprise Get: Unable to find Enterprise with ID: [%s] . Error: %s ", org.ID, err)
		return err
	}

	var orga [1]Enterprise
	err = json.Unmarshal(reply, &orga)

	if err != nil {
		log.Debugf("Enterprise Get: Unable to decode JSON payload: %s ", err)
		return err
	}

	// XXX - Mutate the receiver
	*org = orga[0]

	log.Debugf("Enterprise Get: Found Enterprise with Name: [%s] and ID: [%s]", org.Name, org.ID)
	return nil

}

// enterprises list
func (orglist *EnterpriseSlice) List(c *nuage.Connection) error {

	// XXX - Alternative
	// var orgs []Enterprise

	reply, err := nuage.GetEntity(c, "enterprises")
	if err != nil {
		log.Debugf("Enterprise List: Unable to obtain Enterprise list: %s ", err)
		return err
	}

	if len(reply) == 0 {
		log.Debugf("Enterprise List: Empty list")
		return nil
	}

	// XXX - Alternative
	// err = json.Unmarshal(reply, &orgs)
	err = json.Unmarshal(reply, orglist)

	if err != nil {
		log.Debugf("Enterprise List: Unable to decode JSON payload: %s ", err)
		return err
	}

	//// var orgs []Enterprise
	////
	//// orgs = *orglist
	////
	//// for i, v := range orgs {
	//// 	jsonorg, _ := json.MarshalIndent(v, "", "\t")
	//// 	fmt.Printf("\n\n ===> Org nr [%d]: [%s] <=== \n%#s\n", i, orgs[i].Name, string(jsonorg))
	//// }

	// XXX - Alternative: This effeticvely converts "EnterpriseSlice" to "[]Enterprise"
	// *orglist = orgs

	// log.Fatal("\n\n KABOOM ?? Yes Rico, KABOOM..\n\n")

	log.Debug("Enterprise List: done")
	return nil
}
