// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd || solaris

package sysconf_cgotest

/*
#include <stdlib.h>
#include <unistd.h>
*/
import "C"

import (
	"os"
	"testing"
	"unsafe"

	"github.com/tklauser/go-sysconf"
)

func testPathconfCgoMatch(t *testing.T) {
	for _, tc := range append([]testCase{
		{sysconf.PC_LINK_MAX, C._PC_LINK_MAX, "LINK_MAX"},
		{sysconf.PC_NAME_MAX, C._PC_NAME_MAX, "NAME_MAX"},
		{sysconf.PC_PATH_MAX, C._PC_PATH_MAX, "PATH_MAX"},
		{sysconf.PC_PIPE_BUF, C._PC_PIPE_BUF, "PIPE_BUF"},
	}, pathconfTestCases...) {
		t.Run(tc.name, func(t *testing.T) {
			testPathconfGoCgo(t, tc)
			testFpathconfGoCgo(t, tc)
		})
	}
}

func testPathconfGoCgo(t *testing.T, tc testCase) {
	t.Helper()

	if tc.goVar != int(tc.cVar) {
		t.Errorf("PC_* parameter value for %v is %v, want %v", tc.name, tc.goVar, tc.cVar)
	}

	path := "/"

	goVal, goErr := sysconf.Pathconf(path, tc.goVar)
	if goErr != nil {
		t.Errorf("Pathconf(%s, %s/%d): %v", path, tc.name, tc.goVar, goErr)
		return
	}
	t.Logf("%s = %v", tc.name, goVal)

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cVal, cErr := C.pathconf(cPath, tc.cVar)
	if cErr != nil {
		t.Errorf("C.pathconf(%s, %s/%d): %v", path, tc.name, tc.cVar, cErr)
		return
	}

	if goVal != int64(cVal) {
		t.Errorf("Pathconf(%s, %s/%d) returned %v, want %v", path, tc.name, tc.goVar, goVal, cVal)
	}
}

func testFpathconfGoCgo(t *testing.T, tc testCase) {
	t.Helper()

	f, err := os.Open("/")
	if err != nil {
		t.Fatalf("Open(/): %v", err)
	}
	defer f.Close()

	goVal, goErr := sysconf.Fpathconf(int(f.Fd()), tc.goVar)
	if goErr != nil {
		t.Errorf("Fpathconf(%s/%d): %v", tc.name, tc.goVar, goErr)
		return
	}

	cVal, cErr := C.fpathconf(C.int(f.Fd()), tc.cVar)
	if cErr != nil {
		t.Errorf("C.fpathconf(%s/%d): %v", tc.name, tc.cVar, cErr)
		return
	}

	if goVal != int64(cVal) {
		t.Errorf("Fpathconf(%s/%d) returned %v, want %v", tc.name, tc.goVar, goVal, cVal)
	}
}
