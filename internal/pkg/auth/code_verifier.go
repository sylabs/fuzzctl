// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
)

// base64URLEncode returns a base64url encoding of octets, per RFC 7636 Appendix A.
func base64URLEncode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}

// CodeVerifier represents a code verifier used to implement the Proof Key for Code Exchange by
// OAuth Public Clients specification (RFC 7636).
type CodeVerifier struct {
	v string
}

// NewCodeVerifier creates a new CodeVerifier as per the Proof Key for Code Exchange by OAuth
// Public Clients specification (RFC 7636).
//
// Use String() to obtain the code verifier, and Challenge() to obtain the code challenge.
func NewCodeVerifier(s rand.Source) (*CodeVerifier, error) {
	b, err := readRandom(s, 32)
	if err != nil {
		return nil, err
	}
	cv := CodeVerifier{
		v: base64URLEncode(b),
	}
	return &cv, nil
}

// String returns the code verifier as described in RFC 7636 § 4.1.
func (cv *CodeVerifier) String() string {
	return cv.v
}

// Challenge returns the code challenge as described in RFC 7636 § 4.2.
func (cv *CodeVerifier) Challenge() string {
	b := sha256.Sum256([]byte(cv.v))
	return base64URLEncode(b[:])
}

// ChallengeMethod returns the method used to encode the code challenge.
func (cv *CodeVerifier) ChallengeMethod() string {
	return "S256"
}
