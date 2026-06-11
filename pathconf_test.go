// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package sysconf_test

import (
	"bytes"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/tklauser/go-sysconf"
)

func TestPathconf(t *testing.T) {
	for _, tc := range []struct {
		name int
		desc string
	}{
		{sysconf.PC_NAME_MAX, "PC_NAME_MAX"},
		{sysconf.PC_PATH_MAX, "PC_PATH_MAX"},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			val, err := sysconf.Pathconf("/", tc.name)
			if err != nil {
				t.Fatalf("Pathconf(/, %s): %v", tc.desc, err)
			}
			if val < 0 {
				t.Fatalf("Pathconf(/, %s) = %d, want non-negative", tc.desc, val)
			}

			f, err := os.Open("/")
			if err != nil {
				t.Fatalf("Open(/): %v", err)
			}
			defer f.Close()

			fval, err := sysconf.Fpathconf(int(f.Fd()), tc.name)
			if err != nil {
				t.Fatalf("Fpathconf(/, %s): %v", tc.desc, err)
			}
			if fval != val {
				t.Fatalf("Fpathconf(/, %s) = %d, want %d", tc.desc, fval, val)
			}
		})
	}
}

func TestPathconfGetconf(t *testing.T) {
	testCases := []struct {
		goVar int
		name  string
	}{
		{sysconf.PC_NAME_MAX, "NAME_MAX"},
		{sysconf.PC_PATH_MAX, "PATH_MAX"},
		{sysconf.PC_MAX_CANON, "MAX_CANON"},
		{sysconf.PC_PIPE_BUF, "PIPE_BUF"},
	}

	getconf, err := exec.LookPath("getconf")
	if err != nil {
		t.Skipf("getconf not found in PATH: %v", err)
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			path := "/"
			cmd := exec.Command(getconf, tc.name, path)
			var out bytes.Buffer
			cmd.Stdout = &out
			if err := cmd.Run(); err != nil {
				t.Skipf("failed to invoke getconf: %v", err)
			}
			want, err := strconv.ParseInt(strings.TrimSpace(out.String()), 10, 64)
			if err != nil {
				t.Errorf("strconv.ParseInt: %v", err)
			}

			got, err := sysconf.Pathconf(path, tc.goVar)
			if err != nil {
				t.Errorf("Pathconf(%s, %s/%d): %v", path, tc.name, tc.goVar, err)
			}
			if got != want {
				t.Errorf("Pathconf(%s, %s/%d) returned %v, want %v", path, tc.name, tc.goVar, got, want)
			}
		})
	}
}
