// Copyright 2018 Tobias Klauser. All rights reserved.
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

type testCase struct {
	goVar int
	cVar  C.int
	name  string
}

func testSysconfCgoMatch(t *testing.T) {
	for _, tc := range sysconfTestCases {
		t.Run(tc.name, func(t *testing.T) {
			testSysconfGoCgo(t, tc)
		})
	}
}

func testSysconfCgoMatchInvalid(t *testing.T) {
	for _, tc := range sysconfTestCasesInvalid {
		t.Run(tc.name, func(t *testing.T) {
			testSysconfGoCgoInvalid(t, tc)
		})
	}
}

func testSysconfGoCgo(t *testing.T, tc testCase) {
	t.Helper()

	if tc.goVar != int(tc.cVar) {
		t.Errorf("SC_* parameter value for %v is %v, want %v", tc.name, tc.goVar, tc.cVar)
	}

	goVal, goErr := sysconf.Sysconf(tc.goVar)
	if goErr != nil {
		t.Fatalf("Sysconf(%s/%d): %v", tc.name, tc.goVar, goErr)
	}
	t.Logf("%s = %v", tc.name, goVal)

	cVal, cErr := C.sysconf(tc.cVar)
	if cErr != nil {
		t.Fatalf("C.sysconf(%s/%d): %v", tc.name, tc.cVar, cErr)
	}

	if goVal != int64(cVal) {
		t.Errorf("Sysconf(%v/%d) returned %v, want %v", tc.name, tc.goVar, goVal, cVal)
	}
}

func testSysconfGoCgoInvalid(t *testing.T, tc testCase) {
	t.Helper()

	if tc.goVar != int(tc.cVar) {
		t.Errorf("SC_* parameter value for %v is %v, want %v", tc.name, tc.goVar, tc.cVar)
	}

	_, goErr := sysconf.Sysconf(tc.goVar)
	if goErr == nil {
		t.Errorf("Sysconf(%s/%d) unexpectedly returned without error", tc.name, tc.goVar)
	}

	_, cErr := C.sysconf(tc.cVar)
	if cErr == nil {
		t.Errorf("C.sysconf(%s/%d) unexpectedly returned without error", tc.name, tc.goVar)
	}
}

func testPathconfCgoMatch(t *testing.T) {
	for _, tc := range append([]testCase{
		{sysconf.PC_LINK_MAX, C._PC_LINK_MAX, "LINK_MAX"},
		{sysconf.PC_NAME_MAX, C._PC_NAME_MAX, "NAME_MAX"},
		{sysconf.PC_PATH_MAX, C._PC_PATH_MAX, "PATH_MAX"},
		{sysconf.PC_PIPE_BUF, C._PC_PIPE_BUF, "PIPE_BUF"},
	}, pathconfTestCases...) {
		t.Run(tc.name, func(t *testing.T) {
			testPathconfGoCgo(t, tc)
		})
	}
}

func testPathconfCgoMatchInvalid(t *testing.T) {
	for _, tc := range pathconfTestCasesInvalid {
		t.Run(tc.name, func(t *testing.T) {
			testPathconfGoCgoInvalid(t, tc)
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
		t.Fatalf("Pathconf(%s, %s/%d): %v", path, tc.name, tc.goVar, goErr)
	}
	t.Logf("%s = %v", tc.name, goVal)

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	cVal, cErr := C.pathconf(cPath, tc.cVar)
	if cErr != nil {
		t.Fatalf("C.pathconf(%s, %s/%d): %v", path, tc.name, tc.cVar, cErr)
	}

	if goVal != int64(cVal) {
		t.Errorf("Pathconf(%s, %s/%d) returned %v, want %v", path, tc.name, tc.goVar, goVal, cVal)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Open(%s): %v", path, err)
	}
	defer f.Close()

	goVal, goErr = sysconf.Fpathconf(int(f.Fd()), tc.goVar)
	if goErr != nil {
		t.Fatalf("Fpathconf(%s/%d): %v", tc.name, tc.goVar, goErr)
	}

	cVal, cErr = C.fpathconf(C.int(f.Fd()), tc.cVar)
	if cErr != nil {
		t.Fatalf("C.fpathconf(%s/%d): %v", tc.name, tc.cVar, cErr)
	}

	if goVal != int64(cVal) {
		t.Errorf("Fpathconf(%s/%d) returned %v, want %v", tc.name, tc.goVar, goVal, cVal)
	}
}

func testPathconfGoCgoInvalid(t *testing.T, tc testCase) {
	t.Helper()

	if tc.goVar != int(tc.cVar) {
		t.Errorf("PC_* parameter value for %v is %v, want %v", tc.name, tc.goVar, tc.cVar)
	}

	path := "/"

	_, goErr := sysconf.Pathconf(path, tc.goVar)
	if goErr == nil {
		t.Errorf("Pathconf(%s, %s/%d) unexpectedly returned without error", path, tc.name, tc.goVar)
	}

	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))

	_, cErr := C.pathconf(cPath, tc.cVar)
	if cErr == nil {
		t.Errorf("C.pathconf(%s, %s/%d) unexpectedly returned without error", path, tc.name, tc.goVar)
	}

	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("Open(%s): %v", path, err)
	}
	defer f.Close()

	_, goErr = sysconf.Fpathconf(int(f.Fd()), tc.goVar)
	if goErr == nil {
		t.Errorf("Fpathconf(%s/%d) unexpectedly returned without error", tc.name, tc.goVar)
	}

	_, cErr = C.fpathconf(C.int(f.Fd()), tc.cVar)
	if cErr == nil {
		t.Errorf("C.fpathconf(%s/%d) unexpectedly returned without error", tc.name, tc.goVar)
	}

}
