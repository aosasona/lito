package errtype

const (
	UNABLE_TO_CONNECT_TO_DB   = "unable to establish database connection"
	NO_HOST_PROVIDED          = "no host provided"
	NO_HOST_PROTOCOL_PROVIDED = "host protocol not defined (add http:// or https://)"
	PORT_NOT_AVAILABLE        = "port is not available"
	PORT_ZERO                 = "port cannot be zero"
)
