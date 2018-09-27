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
		t.Fatalf("Sysconf(CLK_TCK): %v", err)
	}
	t.Logf("CLK_TCK = %v", val)
}
