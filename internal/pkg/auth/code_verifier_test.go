// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package auth

import (
	"math/rand"
	"testing"
)

func TestBase64URLEncode(t *testing.T) {
	tests := []struct {
		name string
		b    []byte
		want string
	}{
		{
			"RFC7636AppendixA",
			[]byte{3, 236, 255, 224, 193},
			"A-z_4ME",
		},
		{
			"RFC7636AppendixBChallenge",
			[]byte{116, 24, 223, 180, 151, 153, 224, 37,
				79, 250, 96, 125, 216, 173, 187, 186,
				22, 212, 37, 77, 105, 214, 191, 240,
				91, 88, 5, 88, 83, 132, 141, 121},
			"dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
		},
		{
			"RFC7636AppendixBVerifier",
			[]byte{19, 211, 30, 150, 26, 26, 216, 236,
				47, 22, 177, 12, 76, 152, 46, 8,
				118, 168, 120, 173, 109, 241, 68, 86,
				110, 225, 137, 74, 203, 112, 249, 195},
			"E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := base64URLEncode(tt.b); got != tt.want {
				t.Errorf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCodeVerifier(t *testing.T) {
	tests := []struct {
		name                string
		s                   rand.Source
		wantString          string
		wantChallenge       string
		wantChallengeMethod string
	}{
		{
			"Zero",
			rand.NewSource(0),
			"AZT9wvov_MBB0_8SBFtzyG5P-V_2YqXu6Cq99EotC3U",
			"TMR15BCkWXoSftFpiR7qdK2niDmEyIZyjqZeqDmJSmE",
			"S256",
		},
		{
			"One",
			rand.NewSource(1),
			"Uv38ByGCZU8WP18PmmIdcpVmx00QA3xNe7sEB9Hixkk",
			"hvTaLUAXmpilX93P_hNi6c0NMBUgg7p56UMD1O-1H7o",
			"S256",
		},
		{
			"Two",
			rand.NewSource(2),
			"L4KCy-L5aW8xRMCqTO1W29ln3CiXgGrzvtimOsoW4Ys",
			"611TRvjV45QvNMBnhueb1bkJV8Na6cQ8n5GOZ41C3UA",
			"S256",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cv, err := NewCodeVerifier(tt.s)
			if err != nil {
				t.Fatalf("failed to create code verifier: %v", err)
			}
			if got := cv.String(); got != tt.wantString {
				t.Errorf("got %v, want %v", got, tt.wantString)
			}
			if got := cv.Challenge(); got != tt.wantChallenge {
				t.Errorf("got %v, want %v", got, tt.wantChallenge)
			}
			if got := cv.ChallengeMethod(); got != tt.wantChallengeMethod {
				t.Errorf("got %v, want %v", got, tt.wantChallengeMethod)
			}
		})
	}
}

func TestCodeVerifierAndChallenge(t *testing.T) {
	tests := []struct {
		name                string
		cv                  CodeVerifier
		wantString          string
		wantChallenge       string
		wantChallengeMethod string
	}{
		{
			"RFC7636AppendixB",
			CodeVerifier{"dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk"},
			"dBjftJeZ4CVP-mB92K27uhbUJU1p1r_wW1gFWFOEjXk",
			"E9Melhoa2OwvFrEMTJguCHaoeK1t8URWbuGJSstw-cM",
			"S256",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.cv.String(); got != tt.wantString {
				t.Errorf("got %v, want %v", got, tt.wantString)
			}
			if got := tt.cv.Challenge(); got != tt.wantChallenge {
				t.Errorf("got %v, want %v", got, tt.wantChallenge)
			}
			if got := tt.cv.ChallengeMethod(); got != tt.wantChallengeMethod {
				t.Errorf("got %v, want %v", got, tt.wantChallengeMethod)
			}
		})
	}
}
