package types

// Will be handled by certmagic
type Retry struct {
	MaxTries     int `json:"max_tries"`
	Interval     int `json:"interval"`
	CurrentTries int `json:"current_tries"`
}

type (
	CertStatus string
	DNSStatus  string
)

const (
	CertStatusIssued  CertStatus = "issued"
	CertStatusPending CertStatus = "pending"
	CertStatusFailed  CertStatus = "failed"

	DNSStatusVerified DNSStatus = "verified"
	DNSStatusPending  DNSStatus = "pending"
	DNSStatusAborted  DNSStatus = "aborted"
)

type DomainStatusCert struct {
	Value       CertStatus `json:"value" ts:"type:'issued' | 'pending' | 'failed'"`
	LastUpdated int        `json:"last_updated"`
}

type DomainStatusDNS struct {
	Value       DNSStatus `json:"value" ts:"type:'verified' | 'pending' | 'aborted'"`
	Retry       Retry     `json:"retry"`
	LastUpdated int       `json:"last_updated"`
}

type DomainStatus struct {
	// Cert is the certificate status for the domain
	Cert DomainStatusCert `json:"cert" ts:"'issued' | 'pending' | 'failed'"`

	// DNS is the DNS status for the domain
	DNS DomainStatusDNS `json:"dns" ts:"'verified' | 'pending' | 'aborted'"`
}

type Domain struct {
	// Name is the domain name
	Name string `json:"name"`

	// Status is the certificate and DNS status for the domain
	Status DomainStatus `json:"status"`
}
