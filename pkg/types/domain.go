package types

type (
	DNSStatus string
)

const (
	DNSStatusVerified DNSStatus = "verified"
	DNSStatusPending  DNSStatus = "pending"
	DNSStatusAborted  DNSStatus = "aborted"
	DNSStatusFailed   DNSStatus = "failed"
)

type DomainStatusDNS struct {
	Value             DNSStatus `json:"value" mirror:"type:'verified' | 'pending' | 'aborted' | 'failed'"`
	CurrentRetryCount int       `json:"current_retry_count"`
	LastUpdated       int       `json:"last_updated"`
}

type DomainStatus struct {
	// DNS is the DNS status for the domain
	DNS DomainStatusDNS `json:"dns" mirror:"type:DomainStatusDNS"`
}

type Domain struct {
	// Name is the domain name e.g example.com, sub.example.com
	DomainName string `json:"name"`

	// Status is the certificate and DNS status for the domain
	Status DomainStatus `json:"status,omitempty"`
}
