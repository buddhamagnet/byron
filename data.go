package main

// AWS agent struct, contains all fields
// returned from JSON response.
type Agent struct {
	Id              string `json:"_id"`
	State           string
	CredentialsName string
	CreationMethod  string
	PrivateIPs      []string
	ImageId         string
	CloudType       string
	SizePreset      string
	AccountId       string
	PrivateDNS      string
	Name            string
	Zone            string
	PublicDNS       string
	Region          string
	ProviderId      string
	KeyPair         string
	Provider        string
	Type            string
	AgentKey        string
	StateUpdatedAt  StateUpdatedAt
	Group           string
	PublicIPs       []string
	UpdatedAt       string
	CreatedAt       string
	Deleted         bool
}

type StateUpdatedAt struct {
	Sec  int
	Usec int
}
