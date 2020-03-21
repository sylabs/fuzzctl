// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package client

import "net/http"

func setUserAgent(rt http.RoundTripper, agent string) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return &uaRT{
		rt:    rt,
		agent: agent,
	}
}

type uaRT struct {
	rt    http.RoundTripper
	agent string
}

// RoundTrip sets the user agent, and then calls the next RoundTripper.
func (ug *uaRT) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Set("User-Agent", ug.agent)
	return ug.rt.RoundTrip(r)
}
