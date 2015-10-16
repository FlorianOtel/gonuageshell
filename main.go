package main

import (
	"encoding/json"
	"fmt"
	"net"

	nuage "github.com/FlorianOtel/nuage"

	nuage_v3_2 "github.com/FlorianOtel/nuage_v3_2"

	"strings"

	log "github.com/Sirupsen/logrus"

	"github.com/abiosoft/ishell"
)

var (
	// myconn *nuage.Connection

	// Initialize the connection with some harmless (?) defaults
	myconn = &nuage.Connection{
		Url:     "https://127.0.0.1:8443",
		Apivers: "v3_2",
	}

	// Nuage API connection defaults. We need to keep them as global vars since commands can be invoked in whatever order.
	org  = "org"
	user = "user"
	pass = "pass"
)

func main() {

	// create new shell.
	// by default, new shell includes 'exit', 'help' and 'clear' commands.
	shell := ishell.NewShell()

	shell.Println("Nuage API Interactive Shell")

	shell.Register("greet", mygreet)

	shell.Register("debuglevel", debuglevel)

	// API connection handling

	shell.Register("displayconn", displayconn)

	shell.Register("setconn", setconn)

	shell.Register("makeconn", makeconn)

	// Enterprise CRUD operations
	shell.Register("GET", Get)

	shell.Register("CREATE", Create)

	shell.Register("DELETE", Delete)

	// shell.Register("EnterprisesList", EnterprisesList)

	// shell.Register("EnterpriseGet", EnterpriseGet)

	// start shell
	shell.Start()
}

func Delete(args ...string) (string, error) {
	// Format: <entity> <ID>
	if len(args) != 2 {
		return "Format:\n    DELETE <entity> <ID>", nil
	}
	entity := args[0]
	id := args[1]
	switch entity {
	case "enterprise": // DELETE enterprise <ID>
		org := new(nuage_v3_2.Enterprise)
		org.ID = id
		err := org.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err
	case "domaintemplate": // DELETE domaintemplate <ID>
		dt := new(nuage_v3_2.Domaintemplate)
		dt.ID = id
		err := dt.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err
	case "domain": // DELETE domain <ID>
		domain := new(nuage_v3_2.Domain)
		domain.ID = id
		err := domain.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err
	case "zonetemplate": // DELETE zonetemplate <ID>
		zt := new(nuage_v3_2.Zonetemplate)
		zt.ID = id
		err := zt.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	case "zone": // DELETE zone <ID>
		zone := new(nuage_v3_2.Zone)
		zone.ID = id
		err := zone.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	case "subnet": // DELETE subnet <ID>
		subnet := new(nuage_v3_2.Subnet)
		subnet.ID = id
		err := subnet.Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	case "vport": // DELETE vport <ID>
		var vp nuage_v3_2.VPort
		vp.ID = id
		err := (&vp).Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	case "vminterface": // DELETE vminterface <ID>
		var vmi nuage_v3_2.VMInterface
		vmi.ID = id
		err := (&vmi).Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	case "vm": // DELETE vm <ID>
		var vm nuage_v3_2.VirtualMachine
		vm.ID = id
		err := (&vm).Delete(myconn)
		if err != nil {
			return "", err
		}
		return "", err

	default:
		return "Don't know how to DELETE entity: " + entity, nil
	}
	return "", nil
}

func Create(args ...string) (string, error) {

	// At least 2 arguments: entity <Name>

	if len(args) < 2 {
		return "Format:\n    CREATE <entity> <Name> [options]", nil
	}

	entity := args[0]

	switch entity {
	case "enterprise":
		if len(args) != 2 {
			return "Format:\n    CREATE enterprise <Name>", nil
		}

		// CREATE enterprise <Name>
		org := new(nuage_v3_2.Enterprise)
		org.Name = args[1]
		err := org.Create(myconn)
		if err != nil {
			return "", err
		}

		// JSON pretty-print the org
		jsonorg, _ := json.MarshalIndent(org, "", "\t")
		fmt.Printf("\n ===> Org: [%s] <=== \n%#s\n", org.Name, string(jsonorg))
		return "", err

	case "domaintemplate":
		if len(args) != 3 {
			return "Format:\n    CREATE domaintemplate <Name> <Parent Enterprise ID> ", nil
		}

		// CREATE domaintemplate <Name> <Parent Enterprise ID>
		dt := new(nuage_v3_2.Domaintemplate)
		dt.Name = args[1]
		dt.ParentID = args[2]
		err := dt.Create(myconn)
		if err != nil {
			return "", err
		}
		// JSON pretty-print the domain template
		jsondt, _ := json.MarshalIndent(dt, "", "\t")
		fmt.Printf("\n ===> Domain Template: Name [%s] <=== \n%#s\n", dt.Name, string(jsondt))
		return "Domain Template Create -- done", err

	case "domain":
		if len(args) != 4 {
			return "Format:\n    CREATE domain <Name> <Parent Enterprise ID> <Domain template ID>", nil
		}
		// CREATE domain <Name> <Parent Enterprise ID> <Domain template ID>
		domain := new(nuage_v3_2.Domain)
		domain.Name = args[1]
		domain.ParentID = args[2]
		domain.TemplateID = args[3]
		err := domain.Create(myconn)
		if err != nil {
			return "", err
		}
		jsondomain, _ := json.MarshalIndent(domain, "", "\t")
		fmt.Printf("\n ===> Domain Name [%s] <=== \n%#s\n", domain.Name, string(jsondomain))
		return "Domain Create -- done", err

	case "zonetemplate":
		if len(args) != 3 {
			return "Format:\n    CREATE zonetemplate <Name> <Parent domain template ID>", nil
		}
		// CREATE zonetemplate <Name> <Parent domain template ID>
		zt := new(nuage_v3_2.Zonetemplate)
		zt.Name = args[1]
		zt.ParentID = args[2]
		err := zt.Create(myconn)
		if err != nil {
			return "", err
		}
		jsonzt, _ := json.MarshalIndent(zt, "", "\t")
		fmt.Printf("\n ===> Zone template: Name [%s] <=== \n%#s\n", zt.Name, string(jsonzt))
		return "Zone Template Create -- done", err
	case "zone":
		if len(args) < 3 {
			return "Format:\n    CREATE zone <Name> <Parent Domain ID> [ <Zone template ID> ]", nil
		}
		// CREATE zone <Name> <Parent Domain ID> [ <Zone template ID> ]
		zone := new(nuage_v3_2.Zone)
		zone.Name = args[1]
		zone.ParentID = args[2]
		if len(args) >= 4 {
			zone.TemplateID = args[3]
		}
		err := zone.Create(myconn)
		if err != nil {
			return "", err
		}
		jsonzone, _ := json.MarshalIndent(zone, "", "\t")
		fmt.Printf("\n ===> Zone Name [%s] <=== \n%#s\n", zone.Name, string(jsonzone))
		return "Zone Create -- done", err

	case "subnet":
		if len(args) < 4 {
			return "Format:\n    CREATE subnet <Name> <Parent Zone ID> <Subnet template ID> \n or:\n     CREATE subnet <Name> <Parent Zone ID> <Subnet address> <Subnet mask>\n", nil
		}
		switch len(args) {
		case 4:
			// CREATE subnet <Name> <Parent Subnet ID> <Subnet template ID>
			subnet := new(nuage_v3_2.Subnet)
			subnet.Name = args[1]
			subnet.ParentID = args[2]
			subnet.TemplateID = args[3]
			err := subnet.Create(myconn)
			if err != nil {
				return "", err
			}
			jsonsubnet, _ := json.MarshalIndent(subnet, "", "\t")
			fmt.Printf("\n ===> Subnet Name [%s] <=== \n%#s\n", subnet.Name, string(jsonsubnet))
			return "Subnet Create -- done", err
		case 5:
			// CREATE subnet <Name> <Parent Subnet ID> <Subnet address> <Subnet mask>
			subnet := new(nuage_v3_2.Subnet)
			subnet.Name = args[1]
			subnet.ParentID = args[2]
			// TBD -- make sure these are proper dot notation...
			subnet.Address = args[3]
			subnet.Netmask = args[4]
			err := subnet.Create(myconn)
			if err != nil {
				return "", err
			}
			jsonsubnet, _ := json.MarshalIndent(subnet, "", "\t")
			fmt.Printf("\n ===> Subnet Name [%s] <=== \n%#s\n", subnet.Name, string(jsonsubnet))
			return "Subnet Create -- done", err
		}

	case "vport":
		if len(args) < 3 {
			return "Format:\n    CREATE vport <Name> <Parent Subnet ID> [ options ...]", nil
		}
		// CREATE vport <Name> <Parent Subnet ID> [ options ...]
		subnet := new(nuage_v3_2.Subnet)
		subnet.ID = args[2]

		var vport nuage_v3_2.VPort
		vport.Name = args[1]
		// ???
		vport.ID = vport.Name
		vport.Type = "VM"
		vport.AddressSpoofing = "INHERITED"

		vport.Active = true
		// ???? Not needed but still....
		vport.ParentID = subnet.ID
		vport.ParentType = "subnet"

		jsonvport, err := json.MarshalIndent(vport, "", "\t")
		fmt.Printf("\n ===> Created VPort: Name [%s] <=== \n%#s\n", vport.Name, string(jsonvport))

		vp, err := subnet.AddVPort(myconn, vport)

		if err != nil {
			return "", err
		}

		jsonvp, err := json.MarshalIndent(vp, "", "\t")
		fmt.Printf("\n ===> Created VPort: Name [%s] <=== \n%#s\n", vp.Name, string(jsonvp))

		return "VPort Create -- done", err

	case "vm":
		if len(args) != 5 {
			return "Format:\n    CREATE vm <Name> <UUID> <Interface0-MAC> <Interface0-VPortID>", nil
		}
		// CREATE vm <Name> <UUID> <Interface0-MAC> <Interface0-VPortID>
		var vm nuage_v3_2.VirtualMachine
		vm.Name = args[1]
		vm.UUID = args[2]

		var vmi nuage_v3_2.VMInterface
		vmi.MAC = args[3]
		vmi.VPortID = args[4]

		vm.Interfaces = append(vm.Interfaces, vmi)

		err := (&vm).Create(myconn)
		if err != nil {
			return "", err
		}
		jsonvm, _ := json.MarshalIndent(vm, "", "\t")
		fmt.Printf("\n ===> Virtual Machine: Name [%s] <=== \n%#s\n", vm.Name, string(jsonvm))
		return "Virtual Machine Create -- done", err

	default:
		return "Don't know how to create Nuage API entity: [" + entity + "]" + " with Name [" + args[1] + "]" + " and options: " + strings.Join(args[2:], " "), nil
	}
	return "", nil
}

func Get(args ...string) (string, error) {

	// 1 argument:  <entity>
	// 2 arguments: <entity> <ID>
	// 3 arguments: <entity> <ID> <children>

	if len(args) < 1 || len(args) > 3 {
		return "GET <entity> [ <ID> [ <children> ] ]  ", nil
	}

	entity := args[0]

	switch entity {
	case "enterprises":
		switch len(args) {
		case 1: // GET enterprises
			var orglist nuage_v3_2.EnterpriseSlice
			err := orglist.List(myconn)

			if err != nil {
				return "", err
			}

			// Yuck -- We need to to this hacky type cast since we can iterate through "[]nuage_v3_2.Enterprise"  but not "nuage_v3_2.EnterpriseSlice"
			var orgs []nuage_v3_2.Enterprise
			orgs = orglist

			// Iterate through the list of org's and JSON pretty-print them
			for i, v := range orgs {
				org, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n\n ===> Org nr [%d]: Name [%s] <=== \n%#s\n", i, orgs[i].Name, string(org))
			}

			return "Enterprise list -- done", err

		case 2: // GET enterprises <ID>

			org := new(nuage_v3_2.Enterprise)
			org.ID = args[1]
			err := org.Get(myconn)
			if err != nil {
				return "", err
			}

			// JSON pretty-print the org
			jsonorg, _ := json.MarshalIndent(org, "", "\t")
			fmt.Printf("\n\n ===> Org: Name [%s] <=== \n%#s\n", org.Name, string(jsonorg))

			return "Enterprise Get ID -- done", err

		case 3:
			entityid := args[1]
			child := args[2]

			switch child {
			case "domaintemplates": // GET enterprises <ID> domaintemplates

				// Get list of domain templates for that org
				var dts nuage_v3_2.Domaintemplateslice
				err := dts.List(myconn, entityid)
				if err != nil {
					return "", err
				}
				// Yucky -- type cast from nuage_v3_2.Domaintemplateslice to []nuage_v3_2.Domaintemplate
				var dtl []nuage_v3_2.Domaintemplate
				dtl = dts
				// Iterate through the list of domain templates and JSON pretty-print them
				fmt.Printf("\n ######## Domain templates for Enterprise ID: [%s] ########\n", entityid)
				for i, v := range dtl {
					dt, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> Domain template nr [%d]: Name [%s] <=== \n%#s\n", i, dtl[i].Name, string(dt))
				}

				return "Domain template list -- done", err

			case "domains": // GET enterprises <ID> domains
				// Get list of domains for the Enterprise ID

				var ds nuage_v3_2.Domainslice
				err := ds.List(myconn, entityid)
				if err != nil {
					return "", err
				}
				// Yucky -- type cast from nuage_v3_2.Domainslice to []nuage_v3_2.Domain
				var dl []nuage_v3_2.Domain
				dl = ds
				// Iterate through the list of domains and JSON pretty-print them
				fmt.Printf("\n ######## Domains for Enterprise ID: [%s] ########\n", entityid)
				for i, v := range dl {
					domain, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> Domain nr [%d]: Name [%s] <=== \n%#s\n", i, dl[i].Name, string(domain))
				}

				return "Domain list -- done", err
			}
		}
	case "domaintemplates":
		switch len(args) {
		case 2: // GET domaintemplates <ID>
			dt := new(nuage_v3_2.Domaintemplate)
			dt.ID = args[1]
			err := dt.Get(myconn)

			if err != nil {
				return "", err
			}
			// JSON pretty-print the domain template
			jsondt, _ := json.MarshalIndent(dt, "", "\t")
			fmt.Printf("\n ===> Domain Template: Name [%s] <=== \n%#s\n", dt.Name, string(jsondt))
			return "Domain Template Get -- done", err
		case 3:
			dtid := args[1]
			child := args[2]
			switch child {
			case "zonetemplates": // GET domaintemplates <ID> zonetemplates
				var zts nuage_v3_2.Zonetemplateslice
				err := zts.List(myconn, dtid)
				if err != nil {
					return "", err
				}
				// Yucky -- type cast from nuage_v3_2.Zonetemplateslice to []nuage_v3_2.Zonetemplate
				var zta []nuage_v3_2.Zonetemplate
				zta = zts
				// Iterate through the list of zone templates and JSON pretty-print them
				fmt.Printf("\n ######## Zone templates for Domain template ID: [%s] ########\n", dtid)
				for i, v := range zta {
					zt, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> Zone template nr [%d]: Name [%s] <=== \n%#s\n", i, zta[i].Name, string(zt))
				}

				return "Zone template list -- done", err
			}

		}
	case "domains":
		switch len(args) {
		case 1: // GET domains
			// Get list of domains with "nil" as parent enterprise -- i.e. global list of all domains
			var ds nuage_v3_2.Domainslice
			err := ds.List(myconn, "")
			if err != nil {
				return "", err
			}
			// Yucky -- type cast from nuage_v3_2.Domainslice to []nuage_v3_2.Domain
			var dl []nuage_v3_2.Domain
			dl = ds
			for i, v := range dl {
				jsondomain, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n ===> Domain nr [%d]: Name [%s] <=== \n%#s\n", i, dl[i].Name, string(jsondomain))
			}
			return "Domain list -- done", err
		case 2: // GET domains <ID>
			// Get a specific Domain ID
			domain := new(nuage_v3_2.Domain)
			domain.ID = args[1]
			err := domain.Get(myconn)
			if err != nil {
				return "", err
			}
			jsondomain, _ := json.MarshalIndent(domain, "", "\t")
			fmt.Printf("\n ===> Domain Name [%s] <=== \n%#s\n", domain.Name, string(jsondomain))
			return "Domain Get -- done", err
		case 3:
			switch args[2] {
			case "vports": // GET domains <ID> vports
				var vports []nuage_v3_2.VPort
				var err error
				domain := new(nuage_v3_2.Domain)
				domain.ID = args[1]
				vports, err = domain.VPortsList(myconn)
				if err != nil {
					return "", err
				}
				for i, v := range vports {
					jsonvport, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> VPort nr [%d]: Name [%s] <=== \n%#s\n", i, vports[i].Name, string(jsonvport))
				}
				return "Domain VPorts list -- done", err

			case "vminterfaces": // GET domains <ID> vminterfaces
				// var vmis []nuage_v3_2.VMInterface
				// var err error
				var domain nuage_v3_2.Domain
				domain.ID = args[1]
				vmis, err := (&domain).VMInterfacesList(myconn)
				if err != nil {
					return "", err
				}
				for i, v := range vmis {
					jsonvmi, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> VMInterface nr [%d]: Name [%s] <=== \n%#s\n", i, vmis[i].Name, string(jsonvmi))
				}
				return "Subnet VMInterfaces list -- done", err

			}
		}
	case "zonetemplates":
		if len(args) < 2 {
			return "Format:\n    GET zonetemplates <ID>", nil
		}

		switch len(args) {
		case 2: // GET zonetemplates <ID>
			zt := new(nuage_v3_2.Zonetemplate)
			zt.ID = args[1]
			err := zt.Get(myconn)

			if err != nil {
				return "", err
			}
			// JSON pretty-print the zone template
			jsonzt, _ := json.MarshalIndent(zt, "", "\t")
			fmt.Printf("\n ===> Zone Template: Name [%s] <=== \n%#s\n", zt.Name, string(jsonzt))
			return "Zone Template Get -- done", err
		}
	case "zones":
		switch len(args) {
		case 1: // GET zones
			// Get list of zones with "nil" as parent domain -- i.e. global list of all zones
			var zs nuage_v3_2.Zoneslice
			err := zs.List(myconn, "")
			if err != nil {
				return "", err
			}
			// Yucky -- type cast from nuage_v3_2.Zoneslice to []nuage_v3_2.Zone
			var zl []nuage_v3_2.Zone
			zl = zs
			for i, v := range zl {
				jsonzone, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n ===> Zone nr [%d]: Name [%s] <=== \n%#s\n", i, zl[i].Name, string(jsonzone))
			}
			return "Zone list -- done", err
		case 2: // GET zones <ID>
			// Get a specific Zone ID
			zone := new(nuage_v3_2.Zone)
			zone.ID = args[1]
			err := zone.Get(myconn)
			if err != nil {
				return "", err
			}
			jsonzone, _ := json.MarshalIndent(zone, "", "\t")
			fmt.Printf("\n ===> Zone Name [%s] <=== \n%#s\n", zone.Name, string(jsonzone))
			return "Zone Get -- done", err
		}

	case "subnets":
		switch len(args) {
		case 1: // GET subnets
			// Get list of subnets with "nil" as parent domain -- i.e. global list of all subnets
			var ss nuage_v3_2.Subnetslice
			err := ss.List(myconn, "")
			if err != nil {
				return "", err
			}
			// Yucky -- type cast from nuage_v3_2.Subnetslice to []nuage_v3_2.Subnet
			var za []nuage_v3_2.Subnet
			za = ss
			for i, v := range za {
				jsonsubnet, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n ===> Subnet nr [%d]: Name [%s] <=== \n%#s\n", i, za[i].Name, string(jsonsubnet))
			}
			return "Subnet list -- done", err
		case 2: // GET subnets <ID>
			// Get a specific Subnet ID
			subnet := new(nuage_v3_2.Subnet)
			subnet.ID = args[1]
			err := subnet.Get(myconn)
			if err != nil {
				return "", err
			}
			jsonsubnet, _ := json.MarshalIndent(subnet, "", "\t")
			fmt.Printf("\n ===> Subnet Name [%s] <=== \n%#s\n", subnet.Name, string(jsonsubnet))
			return "Subnet Get -- done", err

		case 3:
			switch args[2] {
			case "vports": // GET subnets <ID> vports
				var vports []nuage_v3_2.VPort
				var err error
				subnet := new(nuage_v3_2.Subnet)
				subnet.ID = args[1]
				vports, err = subnet.VPortsList(myconn)
				if err != nil {
					return "", err
				}
				for i, v := range vports {
					jsonvport, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> VPort nr [%d]: Name [%s] <=== \n%#s\n", i, vports[i].Name, string(jsonvport))
				}
				return "Subnet VPorts list -- done", err
			case "vminterfaces": // GET subnets <ID> vminterfaces
				var vmis []nuage_v3_2.VMInterface
				var err error
				var subnet nuage_v3_2.Subnet
				subnet.ID = args[1]
				vmis, err = (&subnet).VMInterfacesList(myconn)
				if err != nil {
					return "", err
				}
				for i, v := range vmis {
					jsonvmi, _ := json.MarshalIndent(v, "", "\t")
					fmt.Printf("\n ===> VMInterface nr [%d]: Name [%s] <=== \n%#s\n", i, vmis[i].Name, string(jsonvmi))
				}
				return "Subnet VMInterfaces list -- done", err

			}
		}

	case "vports": // GET vports <ID>
		if len(args) != 2 {
			return "Format:\n    GET vports <ID>", nil
		}
		vport := new(nuage_v3_2.VPort)
		vport.ID = args[1]
		err := vport.Get(myconn)
		if err != nil {
			return "", err
		}
		jsonvport, _ := json.MarshalIndent(vport, "", "\t")
		fmt.Printf("\n ===> VPort Name [%s] <=== \n%#s\n", vport.Name, string(jsonvport))
		return "VPort Get -- done", err

	case "vminterfaces":
		switch len(args) {
		case 1: // GET vminterfaces
			var vmis nuage_v3_2.VMInterfaceslice
			err := vmis.List(myconn)
			if err != nil {
				return "", err
			}
			// Yucky -- type cast from nuage_v3_2.VMInterfaceslice to []nuage_v3_2.VMInterface
			var vmia []nuage_v3_2.VMInterface
			vmia = vmis
			for i, v := range vmia {
				jsonvminterface, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n ===> VMInterface nr [%d]: Name [%s] <=== \n%#s\n", i, vmia[i].Name, string(jsonvminterface))
			}
			return "VMinterfaces list -- done", err

		case 2: // GET vminterfaces <ID>
			var vminterface nuage_v3_2.VMInterface
			vminterface.ID = args[1]
			err := (&vminterface).Get(myconn)
			if err != nil {
				return "", err
			}
			jsonvminterface, _ := json.MarshalIndent(vminterface, "", "\t")
			fmt.Printf("\n ===> VMinterface Name [%s] <=== \n%#s\n", vminterface.Name, string(jsonvminterface))
			return "VMinterface Get -- done", err
		}

	case "vms":
		switch len(args) {
		case 1: // GET vms
			var vms nuage_v3_2.VirtualMachineslice
			err := vms.List(myconn)
			if err != nil {
				return "", err
			}
			// Yucky -- type cast from nuage_v3_2.VirtualMachineslice to []nuage_v3_2.VirtualMachine
			var vma []nuage_v3_2.VirtualMachine
			vma = vms
			for i, v := range vma {
				jsonvm, _ := json.MarshalIndent(v, "", "\t")
				fmt.Printf("\n ===> VirtualMachine nr [%d]: Name [%s] <=== \n%#s\n", i, vma[i].Name, string(jsonvm))
			}
			return "VirtualMachine list -- done", err

		case 2: // GET vms <ID>
			var vm nuage_v3_2.VirtualMachine
			vm.ID = args[1]
			err := (&vm).Get(myconn)
			if err != nil {
				return "", err
			}
			jsonvm, _ := json.MarshalIndent(vm, "", "\t")
			fmt.Printf("\n ===> VirtualMachine Name [%s] <=== \n%#s\n", vm.Name, string(jsonvm))
			return "Virtual Machine Get -- done", err
		}

	default:
		// Unknown entity request
		break
	}
	return "Don't know how to process Nuage API entity: " + strings.Join(args, " "), nil
}

////////
////////
////////

// Establish Nuage API connection. Wrapper around Nuage.Connect()
func makeconn(args ...string) (string, error) {

	err := myconn.Connect(org, user, pass)

	if err != nil {
		fmt.Printf("Nuage API connection failed: ")
		return "", err
	} else {
		return "Nuage API connection established", nil
	}
}

// Displays Nuage connection details. Assumes top level var & relies on its string representation
func displayconn(args ...string) (string, error) {
	// fmt.Print(dummyconn)
	fmt.Print(myconn)
	return "", nil
}

// Set Nuage API connection details in top level var
func setconn(args ...string) (string, error) {
	var (
		err   error
		vsdip string
	)

	fmt.Println("Set Nuage API connection Details: Endpoint IP address ; User + Password ; Nuage API version")

	// Get VSD IP
	fmt.Print("  Enter your VSD IP address> ")
	_, err = fmt.Scanln(&vsdip)
	if err != nil {
		return "Error: ", err
	}

	if net.ParseIP(vsdip) == nil {
		return "'" + vsdip + "'" + " is not a valid IP address", nil
	}

	// We assume (hardocde) that the URL for the Nuage API has the form "https://<VSD_ip_addr>:8443"
	myconn.Url = "https://" + vsdip + ":8443"

	// Get Enterprise name.
	fmt.Print("  Enter your Enterprise (organization) name. Leave empty if default > ")
	_, err = fmt.Scanln(&org)

	if err != nil {
		if err.Error() != "unexpected newline" {
			return "Error: ", err
		}
	}

	// Get username
	fmt.Print("  Enter your username. Leave empty if default > ")
	_, err = fmt.Scanln(&user)
	if err != nil {
		if err.Error() != "unexpected newline" {
			return "Error: ", err
		}
	}

	// Get password
	fmt.Print("  Enter your password. Leave empty if default > ")
	_, err = fmt.Scanln(&pass)
	if err != nil {
		if err.Error() != "unexpected newline" {
			return "Error: ", err
		}
	}

	// TBD: Insert code for changing the Nuage API version here. Currently only 3_2 (hardcoded)
	return "", nil

}

// Set debug level
func debuglevel(args ...string) (string, error) {
	var loglevel string
	fmt.Print("Set debug level: Debug, Info (default): ")
	_, err := fmt.Scanln(&loglevel)
	if err != nil {
		if err.Error() == "unexpected newline" {
			log.SetLevel(log.InfoLevel)
			return "Debug level now set to Info", nil
		}
		return "Error", err
	}

	switch loglevel {
	case "Debug":
		log.SetLevel(log.DebugLevel)
		return "Debug level now set to: " + loglevel, nil
	case "Info":
		log.SetLevel(log.InfoLevel)
		return "Debug level now set to: " + loglevel, nil
	default:
		log.SetLevel(log.InfoLevel)
		return "Debug level now set to Info", nil
	}
}

// Test function -- dummy greet
func mygreet(args ...string) (string, error) {
	name := "Stranger"
	if len(args) > 0 {
		name = strings.Join(args, " ")
	}
	return "Hello " + name, nil
}
