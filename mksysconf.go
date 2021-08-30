// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"os/exec"
	"regexp"
	"runtime"
)

func gensysconf(in, out, goos, goarch string) error {
	if _, err := os.Stat(in); err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	cmd := exec.Command("go", "tool", "cgo", "-godefs", in)
	defer os.RemoveAll("_obj")
	b, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Fprintln(os.Stderr, string(b))
		return err
	}

	goBuild, build := goos, goos
	if goarch != "" {
		goBuild = fmt.Sprintf("%s && %s", goos, goarch)
		build = fmt.Sprintf("%s,%s", goos, goarch)
	}

	r := fmt.Sprintf(`$1

//go:build %s
// +build %s`, goBuild, build)
	cgoCommandRegex := regexp.MustCompile(`(cgo -godefs .*)`)
	b = cgoCommandRegex.ReplaceAll(b, []byte(r))

	b, err = format.Source(b)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(out, b, 0644)
}

func main() {
	goos, goarch := runtime.GOOS, runtime.GOARCH
	if goos == "illumos" {
		goos = "solaris"
	}
	defs := fmt.Sprintf("sysconf_defs_%s.go", goos)
	if err := gensysconf(defs, "z"+defs, goos, ""); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	vals := fmt.Sprintf("sysconf_values_%s.go", runtime.GOOS)
	// sysconf variable values are GOARCH-specific, thus write per GOARCH
	zvals := fmt.Sprintf("zsysconf_values_%s_%s.go", runtime.GOOS, runtime.GOARCH)
	if err := gensysconf(vals, zvals, goos, goarch); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
