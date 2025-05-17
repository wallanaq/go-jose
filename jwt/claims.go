package jwt

import "encoding/json"

// Claims defines the set of registered claims as per RFC 7519 (JWT).
// These are commonly used to control token validity and audience.
type Claims struct {
	Issuer         string   `json:"iss,omitempty"` // Token issuer
	Subject        string   `json:"sub,omitempty"` // Token subject
	Audience       []string `json:"aud,omitempty"` // Intended recipient(s)
	ExpirationTime int64    `json:"exp,omitempty"` // Expiration time (Unix timestamp)
	NotBefore      int64    `json:"nbf,omitempty"` // Time before which the token must not be accepted
	IssuedAt       int64    `json:"iat,omitempty"` // Token issuance time
	JWTID          string   `json:"jti,omitempty"` // Unique token identifier
	custom         map[string]any
}

// Get returns the raw value associated with the given key from the custom claims.
// It returns the value and a boolean indicating whether the key exists.
func (c Claims) Get(key string) (any, bool) {
	v, ok := c.custom[key]
	return v, ok
}

// GetString returns the value associated with the given key as a string.
// It returns the string value and a boolean indicating success.
// If the value is not a string or the key does not exist, it returns ("", false).
func (c Claims) GetString(key string) (string, bool) {
	v, ok := c.custom[key]
	if !ok {
		return "", ok
	}
	str, ok := v.(string)
	return str, ok
}

// GetInt64 returns the value associated with the given key as an int64.
// It supports int64 and json.Number types.
// It returns the int64 value and a boolean indicating success.
func (c Claims) GetInt64(key string) (int64, bool) {
	switch v := c.custom[key].(type) {
	case int64:
		return v, true
	case json.Number:
		i, err := v.Int64()
		return i, err == nil
	default:
		return 0, false
	}
}

// GetStringSlice returns the value associated with the given key as a slice of strings.
// It expects the underlying value to be a []interface{} where each element is a string.
// Returns the string slice and a boolean indicating success.
func (c Claims) GetStringSlice(key string) ([]string, bool) {
	v, ok := c.custom[key]
	if !ok {
		return nil, ok
	}
	slice, ok := v.([]any)
	if !ok {
		return nil, ok
	}
	out := make([]string, len(slice))
	for _, item := range slice {
		if s, ok := item.(string); ok {
			out = append(out, s)
		} else {
			return nil, ok
		}
	}
	return out, true
}

// CustomKeys returns a list of all keys present in the custom claims map.
func (c Claims) CustomKeys() []string {
	keys := make([]string, 0, len(c.custom))
	for k := range c.custom {
		keys = append(keys, k)
	}
	return keys
}
