// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"runtime"
)

func gensysconf() error {
	defs := "sysconf_defs_" + runtime.GOOS + ".go"
	cmd := exec.Command("go", "tool", "cgo", "-godefs", defs)
	defer os.RemoveAll("_obj")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprint(os.Stderr, string(b))
		return err
	}
	b, err = format.Source(b)
	if err != nil {
		return err
	}
	zsysconf := "z" + defs
	// TODO(tk): differentiate per GOARCH?
	if err := ioutil.WriteFile(zsysconf, b, 0644); err != nil {
		return err
	}
	return nil
}

func main() {
	if err := gensysconf(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
