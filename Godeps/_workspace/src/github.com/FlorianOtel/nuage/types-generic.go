package nuage

type Connection struct {
	Url     string
	Apivers string
	token   *Authtoken
}

type Authtoken struct {
	Apikey         string
	APIKeyExpiry   int64
	Id             string
	AvatarData     string
	AvatarType     string
	Email          string
	EnterpriseID   string
	EnterpriseName string
	EntityScop     string
	ExternalID     string
	ExternalId     string
	FirstName      string
	LastName       string
	MobileNumber   string
	Password       string
	Role           string
	UserName       string
}
