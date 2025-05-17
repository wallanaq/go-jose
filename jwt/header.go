package jwt

// Header represents the JOSE header of a JWT.
// It includes metadata about the cryptographic operations to be performed, such as the algorithm used, key identifiers, and certificate references.
// Relevant specifications: RFC 7515 (JWS), RFC 7516 (JWE), RFC 7517 (JWK).
type Header struct {
	Algorithm      string   `json:"alg,omitempty"`      // Signing or encryption algorithm
	Encryption     string   `json:"enc,omitempty"`      // Content encryption algorithm
	Compression    string   `json:"zip,omitempty"`      // Compression algorithm (e.g., "DEF")
	JWKSetURL      string   `json:"jku,omitempty"`      // URL for the JWK Set
	JWK            string   `json:"jwk,omitempty"`      // Inline JWK
	KeyID          string   `json:"kid,omitempty"`      // Key ID
	X509URL        string   `json:"x5u,omitempty"`      // URL for X.509 certificate
	X509Cert       string   `json:"x5c,omitempty"`      // X.509 certificate chain (base64 DER)
	X509Thumb      string   `json:"x5t,omitempty"`      // SHA-1 thumbprint of the certificate
	X509CertSHA256 string   `json:"x5t#S256,omitempty"` // SHA-256 thumbprint of the certificate
	Type           string   `json:"typ,omitempty"`      // Token type (e.g., "JWT")
	ContentType    string   `json:"cty,omitempty"`      // Embedded content type
	Critical       string   `json:"crit,omitempty"`     // Indicates critical header parameters
	KeyType        string   `json:"kty,omitempty"`      // Key type (e.g., "RSA", "EC")
	PublicKeyUse   string   `json:"use,omitempty"`      // Intended public key use ("sig" or "enc")
	KeyOps         []string `json:"key_ops,omitempty"`  // Permitted key operations
}
