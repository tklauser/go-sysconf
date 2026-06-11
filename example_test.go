// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package sysconf_test

import (
	"fmt"
	"os"

	"github.com/tklauser/go-sysconf"
)

func ExampleSysconf_clktck() {
	clktck, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Sysconf: %v\n", err)
	}
	fmt.Printf("sysconf(SC_CLK_TCK) = %v\n", clktck)
}

func ExampleSysconf_invalidParameter() {
	_, err := sysconf.Sysconf(-1)
	fmt.Print(err)

	// Output: invalid argument
}

func ExamplePathconf_namemax() {
	// get the maximum filename length for the root directory
	namemax, err := sysconf.Pathconf("/", sysconf.PC_NAME_MAX)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Pathconf: %v\n", err)
	}
	fmt.Printf(`pathconf("/", PC_NAME_MAX) = %v\n`, namemax)
}

func ExamplePathconf_invalidParameter() {
	_, err := sysconf.Pathconf("/", -1)
	fmt.Print(err)

	// Output: invalid argument
}
