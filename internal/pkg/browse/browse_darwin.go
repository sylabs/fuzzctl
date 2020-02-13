// Copyright (c) 2020, Sylabs, Inc. All rights reserved.

package browse

import (
	"os/exec"
)

// Open opens url in a browser.
func (*Browser) Open(url string) error {
	return exec.Command("open", url).Start()
}
