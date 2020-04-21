// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_test

import (
	"testing"

	"github.com/tklauser/go-sysconf"
)

func TestSysconf(t *testing.T) {
	// Just test the basic functionality here. The comparison tests against
	// C.sysconf are in the test directory.
	val, err := sysconf.Sysconf(sysconf.SC_CLK_TCK)
	if err != nil {
		t.Errorf("Sysconf(SC_CLK_TCK): %v", err)
	}
	t.Logf("clock ticks = %v", val)

	_, err = sysconf.Sysconf(-1)
	if err == nil {
		t.Errorf("Sysconf(-1) returned %v, want non-nil", err)
	}
}

func TestOpenMax(t *testing.T) {

	openMax, err := sysconf.Sysconf(sysconf.SC_OPEN_MAX)
	if err != nil {
		t.Fatalf("Sysconf(SC_OPEN_MAX): %v", err)
	}

	// from https://pubs.opengroup.org/onlinepubs/009695399/basedefs/limits.h.html
	_POSIX_OPEN_MAX := int64(20)

	// according to sysconf(3), OPEN_MAX must be ≥ _POSIX_OPEN_MAX
	if openMax < _POSIX_OPEN_MAX {
		t.Errorf("Sysconf(SC_OPEN_MAX) (%d) expected to be greater or equal _POSIX_OPEN_MAX (%d)",
			openMax, _POSIX_OPEN_MAX)
	}
}
