package jwt

// JWT represents a complete JSON Web Token as defined in RFC 7519.
// It consists of a header, claims (payload), and a signature.
type JWT struct {
	Header    Header // The JOSE header, as defined in RFC 7515 / RFC 7516
	Claims    Claims // Standard claims defined in RFC 7519 (JWT)
	Signature string // The signature string as per JWS (RFC 7515)
}
