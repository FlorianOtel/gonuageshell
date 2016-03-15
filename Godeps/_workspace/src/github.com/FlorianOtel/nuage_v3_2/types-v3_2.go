package nuage_v3_2

////////
//////// VirtualMachine and related
////////

type VirtualMachine struct {
	AppName         string        `json:"appName,omitempty"`
	DeleteExpiry    int64         `json:"deleteExpiry,omitempty"`
	DeleteMode      string        `json:"deleteMode,omitempty"`
	DomainIDs       []string      `json:"domainIDs,omitempty"`
	EnterpriseID    string        `json:"enterpriseID,omitempty"`
	EnterpriseName  string        `json:"enterpriseName,omitempty"`
	HypervisorIP    string        `json:"hypervisorIP,omitempty"`
	Interfaces      []VMInterface `json:"interfaces,omitempty"`
	L2DomainIDs     []string      `json:"l2DomainIDs,omitempty"`
	Name            string        `json:"name"`
	ResyncInfo      VMResync      `json:"resyncInfo,omitempty"`
	SiteIdentifier  string        `json:"siteIdentifier,omitempty"`
	SubnetIDs       []string      `json:"subnetIDs,omitempty"`
	UserID          string        `json:"userID,omitempty"`
	UserName        string        `json:"userName,omitempty"`
	UUID            string        `json:"UUID"`
	Status          string        `json:"status,omitempty"`
	ReasonType      string        `json:"reasonType,omitempty"`
	VRSID           string        `json:"VRSID,omitempty"`
	ZoneIDs         []string      `json:"zoneIDs,omitempty"`
	CreationDate    int64         `json:"creationDate,omitempty"`
	LastUpdatedBy   string        `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate int64         `json:"lastUpdatedDate,omitempty"`
	Owner           string        `json:"owner,omitempty"`
	EntityScope     string        `json:"entityScope,omitempty"`
	ExternalID      string        `json:"externalID,omitempty"`
	ID              string        `json:"ID,omitempty"`
	ParentID        string        `json:"parentID,omitempty"`
	ParentType      string        `json:"parentType,omitempty"`
}

type VirtualMachineslice []VirtualMachine

// XXX -- TBD -- the REST listing shows "children" as a field as well, but missing from the 3.2 API description (???). Omitting...
type VMResync struct {
	LastRequestTimestamp    int64  `json:"lastRequestTimestamp,omitempty"`
	LastTimeResyncInitiated int64  `json:"lastTimeResyncInitiated,omitempty"`
	Status                  string `json:"status,omitempty"`
	CreationDate            int64  `json:"creationDate,omitempty"`
	LastUpdatedBy           string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate         int64  `json:"lastUpdatedDate,omitempty"`
	Owner                   string `json:"owner,omitempty"`
	EntityScope             string `json:"entityScope,omitempty"`
	ExternalID              string `json:"externalID,omitempty"`
	ID                      string `json:"ID,omitempty"`
	ParentID                string `json:"parentID,omitempty"`
	ParentType              string `json:"parentType,omitempty"`
}

////////
//////// VMInterface
////////

type VMInterface struct {
	IPAddress                   string `json:"IPAddress,omitempty"`
	MAC                         string `json:"MAC"`
	MultiNICVPortName           string `json:"multiNICVPortName,omitempty"`
	Name                        string `json:"name,omitempty"`
	VMUUID                      string `json:"VMUUID,omitempty"`
	AssociatedFloatingIPAddress string `json:"associatedFloatingIPAddress,omitempty"`
	AttachedNetworkID           string `json:"attachedNetworkID,omitempty"`
	AttachedNetworkType         string `json:"attachedNetworkType,omitempty"`
	DomainID                    string `json:"domainID,omitempty"`
	DomainName                  string `json:"domainName,omitempty"`
	Gateway                     string `json:"gateway,omitempty"`
	Netmask                     string `json:"netmask,omitempty"`
	NetworkName                 string `json:"networkName,omitempty"`
	PolicyDecisionID            string `json:"policyDecisionID,omitempty"`
	TierID                      string `json:"tierID,omitempty"`
	VPortID                     string `json:"VPortID,omitempty"`
	VPortName                   string `json:"VPortName,omitempty"`
	ZoneID                      string `json:"zoneID,omitempty"`
	ZoneName                    string `json:"zoneName,omitempty"`
	CreationDate                int64  `json:"creationDate,omitempty"`
	LastUpdatedBy               string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate             int64  `json:"lastUpdatedDate,omitempty"`
	Owner                       string `json:"owner,omitempty"`
	EntityScope                 string `json:"entityScope,omitempty"`
	ExternalID                  string `json:"externalID,omitempty"`
	ID                          string `json:"ID,omitempty"`
	ParentID                    string `json:"parentID,omitempty"`
	ParentType                  string `json:"parentType,omitempty"`
}

type VMInterfaceslice []VMInterface

////////
//////// vPort
////////

type VPort struct {
	Active                              bool   `json:"active"`
	AddressSpoofing                     string `json:"addressSpoofing"`
	VLANID                              string `json:"VLANID,omitempty"`
	Description                         string `json:"description,omitempty"`
	DomainID                            string `json:"domainID,omitempty"`
	Multicast                           string `json:"multicast,omitempty"`
	AssociatedFloatingIPID              string `json:"associatedFloatingIPID,omitempty"`
	HasAttachedInterfaces               bool   `json:"hasAttachedInterfaces,omitempty"`
	AssociatedMulticastChannelMapID     string `json:"associatedMulticastChannelMapID,omitempty"`
	MultiNICVPortID                     string `json:"multiNICVPortID,omitempty"`
	Name                                string `json:"name"`
	OperationalState                    string `json:"operationalState,omitempty"`
	AssociatedSendMulticastChannelMapID string `json:"associatedSendMulticastChannelMapID,omitempty"`
	SystemType                          string `json:"systemType,omitempty"`
	Type                                string `json:"type"`
	ZoneID                              string `json:"zoneID,omitempty"`
	CreationDate                        int64  `json:"creationDate,omitempty"`
	LastUpdatedBy                       string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                     int64  `json:"lastUpdatedDate,omitempty"`
	Owner                               string `json:"owner,omitempty"`
	EntityScope                         string `json:"entityScope,omitempty"`
	ExternalID                          string `json:"externalID,omitempty"`
	ID                                  string `json:"ID,omitempty"`
	ParentID                            string `json:"parentID,omitempty"`
	ParentType                          string `json:"parentType,omitempty"`
}

// type VPortslice []VPort

////////
//////// Subnet
////////

type Subnet struct {
	Address                           string `json:"address"`
	AssociatedApplicationID           string `json:"associatedApplicationID,omitempty"`
	AssociatedApplicationObjectID     string `json:"associatedApplicationObjectID,omitempty"`
	AssociatedApplicationObjectType   string `json:"associatedApplicationObjectType,omitempty"`
	AssociatedSharedNetworkResourceID string `json:"associatedSharedNetworkResourceID,omitempty"`
	TemplateID                        string `json:"templateID,omitempty"`
	DefaultAction                     string `json:"defaultAction,omitempty"`
	Description                       string `json:"description,omitempty"`
	Gateway                           string `json:"gateway,omitempty"`
	GatewayMACAddress                 string `json:"gatewayMACAddress,omitempty"`
	IPType                            string `json:"IPType,omitempty"`
	MaintenanceMode                   string `json:"maintenanceMode,omitempty"`
	Name                              string `json:"name"`
	Netmask                           string `json:"netmask"`
	PATEnabled                        string `json:"PATEnabled,omitempty"`
	PolicyGroupID                     int32  `json:"policyGroupID,omitempty"`
	Public                            bool   `json:"public,omitempty"`
	RouteDistinguisher                string `json:"routeDistinguisher,omitempty"`
	RouteTarget                       string `json:"routeTarget,omitempty"`
	ServiceID                         int32  `json:"serviceID,omitempty"`
	VnId                              int32  `json:"vnId,omitempty"`
	Multicast                         string `json:"multicast,omitempty"`
	Encryption                        string `json:"encryption,omitempty"`
	AssociatedMulticastChannelMapID   string `json:"associatedMulticastChannelMapID,omitempty"`
	ProxyARP                          bool   `json:"proxyARP,omitempty"`
	SplitSubnet                       bool   `json:"splitSubnet,omitempty"`
	CreationDate                      int64  `json:"creationDate,omitempty"`
	LastUpdatedBy                     string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                   int64  `json:"lastUpdatedDate,omitempty"`
	Owner                             string `json:"owner,omitempty"`
	EntityScope                       string `json:"entityScope,omitempty"`
	ExternalID                        string `json:"externalID,omitempty"`
	ID                                string `json:"ID,omitempty"`
	ParentID                          string `json:"parentID,omitempty"`
	ParentType                        string `json:"parentType,omitempty"`
}

type Subnetslice []Subnet

////////
//////// Zone
////////

type Zone struct {
	Address                         string `json:"address,omitempty"`
	AssociatedApplicationID         string `json:"associatedApplicationID,omitempty"`
	AssociatedApplicationObjectID   string `json:"associatedApplicationObjectID,omitempty"`
	AssociatedApplicationObjectType string `json:"associatedApplicationObjectType,omitempty"`
	TemplateID                      string `json:"templateID,omitempty"`
	Description                     string `json:"description,omitempty"`
	IPType                          string `json:"IPType,omitempty"`
	MaintenanceMode                 string `json:"maintenanceMode,omitempty"`
	Name                            string `json:"name"`
	Netmask                         string `json:"netmask,omitempty"`
	NumberOfHostsInSubnets          int    `json:"numberOfHostsInSubnets,omitempty"`
	PolicyGroupID                   int32  `json:"policyGroupID,omitempty"`
	PublicZone                      bool   `json:"publicZone,omitempty"`
	Multicast                       string `json:"multicast,omitempty"`
	Encryption                      string `json:"encryption,omitempty"`
	AssociatedMulticastChannelMapID string `json:"associatedMulticastChannelMapID,omitempty"`
	CreationDate                    int64  `json:"creationDate,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                 int64  `json:"lastUpdatedDate,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID"`
	ParentType                      string `json:"parentType,omitempty"`
}

type Zoneslice []Zone

////////
//////// Zone template
////////

type Zonetemplate struct {
	Address                         string `json:"address,omitempty"`
	Description                     string `json:"description,omitempty"`
	IPType                          string `json:"IPType,omitempty"`
	Name                            string `json:"name"`
	Netmask                         string `json:"netmask,omitempty"`
	NumberOfHostsInSubnets          int    `json:"numberOfHostsInSubnets,omitempty"`
	PublicZone                      bool   `json:"publicZone,omitempty"`
	Multicast                       string `json:"multicast,omitempty"`
	Encryption                      string `json:"encryption,omitempty"`
	AssociatedMulticastChannelMapID string `json:"associatedMulticastChannelMapID,omitempty"`
	CreationDate                    int64  `json:"creationDate,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                 int64  `json:"lastUpdatedDate,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID"`
	ParentType                      string `json:"parentType,omitempty"`
}

type Zonetemplateslice []Zonetemplate

////////
//////// Domain
////////

type Domain struct {
	ApplicationDeploymentPolicy     string   `json:"applicationDeploymentPolicy,omitempty"`
	TemplateID                      string   `json:"templateID,omitempty"`
	BackHaulRouteDistinguisher      string   `json:"backHaulRouteDistinguisher,omitempty"`
	BackHaulRouteTarget             string   `json:"backHaulRouteTarget,omitempty"`
	BackHaulVNID                    int32    `json:"backHaulVNID,omitempty"`
	CustomerID                      int32    `json:"customerID,omitempty"`
	Description                     string   `json:"description,omitempty"`
	DHCPBehavior                    string   `json:"DHCPBehavior,omitempty"`
	DHCPServerAddress               string   `json:"DHCPServerAddress,omitempty"`
	DhcpServerAddresses             []string `json:"dhcpServerAddresses,omitempty"`
	TunnelType                      string   `json:"tunnelType,omitempty"`
	ECMPCount                       int      `json:"ECMPCount,omitempty"`
	Multicast                       string   `json:"multicast,omitempty"`
	Encryption                      string   `json:"encryption,omitempty"`
	ExportRouteTarget               string   `json:"exportRouteTarget,omitempty"`
	GlobalRoutingEnabled            bool     `json:"globalRoutingEnabled,omitempty"`
	ImportRouteTarget               string   `json:"importRouteTarget,omitempty"`
	LabelID                         int32    `json:"labelID,omitempty"`
	LeakingEnabled                  bool     `json:"leakingEnabled,omitempty"`
	MaintenanceMode                 string   `json:"maintenanceMode,omitempty"`
	AssociatedMulticastChannelMapID string   `json:"associatedMulticastChannelMapID,omitempty"`
	Name                            string   `json:"name,omitempty"`
	PATEnabled                      string   `json:"PATEnabled,omitempty"`
	PermittedAction                 string   `json:"permittedAction,omitempty"`
	PolicyChangeStatus              string   `json:"policyChangeStatus,omitempty"`
	RouteDistinguisher              string   `json:"routeDistinguisher,omitempty"`
	RouteTarget                     string   `json:"routeTarget,omitempty"`
	SecondaryDHCPServerAddress      string   `json:"secondaryDHCPServerAddress,omitempty"`
	ServiceID                       int32    `json:"serviceID,omitempty"`
	Stretched                       bool     `json:"stretched,omitempty"`
	UplinkPreference                string   `json:"uplinkPreference,omitempty"`
	CreationDate                    int64    `json:"creationDate,omitempty"`
	LastUpdatedBy                   string   `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                 int64    `json:"lastUpdatedDate,omitempty"`
	Owner                           string   `json:"owner,omitempty"`
	EntityScope                     string   `json:"entityScope,omitempty"`
	ExternalID                      string   `json:"externalID,omitempty"`
	ID                              string   `json:"ID,omitempty"`
	ParentID                        string   `json:"parentID"`
	ParentType                      string   `json:"parentType,omitempty"`
}

type Domainslice []Domain

////////
//////// Domain template
////////

type Domaintemplate struct {
	Description                     string `json:"description,omitempty"`
	Multicast                       string `json:"multicast,omitempty"`
	Encryption                      string `json:"encryption,omitempty"`
	AssociatedMulticastChannelMapID string `json:"associatedMulticastChannelMapID,omitempty"`
	Name                            string `json:"name"`
	PolicyChangeStatus              string `json:"policyChangeStatus,omitempty"`
	CreationDate                    int64  `json:"creationDate,omitempty"`
	LastUpdatedBy                   string `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                 int64  `json:"lastUpdatedDate,omitempty"`
	Owner                           string `json:"owner,omitempty"`
	EntityScope                     string `json:"entityScope,omitempty"`
	ExternalID                      string `json:"externalID,omitempty"`
	ID                              string `json:"ID,omitempty"`
	ParentID                        string `json:"parentID"`
	ParentType                      string `json:"parentType,omitempty"`
}

type Domaintemplateslice []Domaintemplate

////////
////////
////////

type Enterprise struct {
	AllowAdvancedQOSConfiguration         bool     `json:"allowAdvancedQOSConfiguration,omitempty"`
	AllowedForwardingClasses              []string `json:"allowedForwardingClasses,omitempty"`
	AllowGatewayManagement                bool     `json:"allowGatewayManagement,omitempty"`
	AllowTrustedForwardingClass           bool     `json:"allowTrustedForwardingClass,omitempty"`
	AssociatedEnterpriseSecurityID        string   `json:"associatedEnterpriseSecurityID,omitempty"`
	AssociatedGroupKeyEncryptionProfileID string   `json:"associatedGroupKeyEncryptionProfileID,omitempty"`
	AssociatedKeyServerMonitorID          string   `json:"associatedKeyServerMonitorID,omitempty"`
	AvatarData                            string   `json:"avatarData,omitempty"`
	AvatarType                            []string `json:"avatarType,omitempty"`
	CustomerID                            int64    `json:"customerID,omitempty"`
	Description                           string   `json:"description"`
	DHCPLeaseinterval                     int      `json:"DHCPLeaseinterval,omitempty"`
	EncryptionManagementMode              string   `json:"encryptionManagementMode,omitempty"`
	EnterpriseProfileID                   string   `json:"enterpriseProfileID,omitempty"`
	FloatingIPsQuota                      int      `json:"floatingIPsQuota,omitempty"`
	FloatingIPsUsed                       int      `json:"floatingIPsUsed,omitempty"`
	LDAPAuthorizationEnabled              bool     `json:"LDAPAuthorizationEnabled,omitempty"`
	LDAPEnabled                           bool     `json:"LDAPEnabled,omitempty"`
	Name                                  string   `json:"name"`
	ReceiveMultiCastListID                string   `json:"receiveMultiCastListID,omitempty"`
	SendMultiCastListID                   string   `json:"sendMultiCastListID,omitempty"`
	CreationDate                          int64    `json:"creationDate,omitempty"`
	LastUpdatedBy                         string   `json:"lastUpdatedBy,omitempty"`
	LastUpdatedDate                       int64    `json:"lastUpdatedDate,omitempty"`
	Owner                                 string   `json:"owner,omitempty"`
	EntityScope                           string   `json:"entityScope,omitempty"`
	ExternalID                            string   `json:"externalID,omitempty"`
	ID                                    string   `json:"ID,omitempty"`
	ParentID                              string   `json:"parentID,omitempty"`
	ParentType                            string   `json:"parentType,omitempty"`
}

type EnterpriseSlice []Enterprise
