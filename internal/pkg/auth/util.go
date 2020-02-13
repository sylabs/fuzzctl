// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"fmt"
	"math/rand"
)

// readRandom generates a slice of n bytes read from s.
func readRandom(s rand.Source, n int) ([]byte, error) {
	b := make([]byte, n)
	if n, err := rand.New(s).Read(b); err != nil {
		return nil, err
	} else if n != len(b) {
		return nil, fmt.Errorf("pkce: read unexpcted byte count (%d/%d)", len(b), n)
	}
	return b, nil
}

// newState reads from s and generates a state value for use in an OAuth request.
func newState(s rand.Source) (string, error) {
	b, err := readRandom(s, 32)
	if err != nil {
		return "", err
	}
	return base64URLEncode(b), nil
}
