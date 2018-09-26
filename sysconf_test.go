// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_test

import (
	"os"
	"testing"

	"github.com/tklauser/go-sysconf"
)

func TestSysconf(t *testing.T) {
	testCases := []struct {
		val      int
		name     string
		expect   bool
		expected int64
	}{
		// POSIX.1
		{sysconf.SC_ARG_MAX, "ARG_MAX", false, 0},
		{sysconf.SC_CHILD_MAX, "CHILD_MAX", false, 0},
		{sysconf.SC_CLK_TCK, "clock ticks", false, 0},
		{sysconf.SC_PAGESIZE, "PAGESIZE", true, int64(os.Getpagesize())},
		{sysconf.SC_SYMLOOP_MAX, "SYMLOOP_MAX", false, 0},
		{sysconf.SC_VERSION, "_POSIX_VERSION", false, 0},
		// POSIX.2
		{sysconf.SC_BC_BASE_MAX, "BC_BASE_MAX", false, 0},
		{sysconf.SC_COLL_WEIGHTS_MAX, "COLL_WEIGHTS_MAX", false, 0},
		{sysconf.SC_LINE_MAX, "LINE_MAX", false, 0},
		{sysconf.SC_2_VERSION, "_POSIX2_VERSION", false, 0},
		{sysconf.SC_2_C_DEV, "_POSIX2_C_DEV", false, 0},
		{sysconf.SC_2_C_VERSION, "_POSIX2_C_VERSION", false, 0},
		// non-standard
		{sysconf.SC_PHYS_PAGES, "number of pages of physical memory", false, 0},
		{sysconf.SC_AVPHYS_PAGES, "number of pages of physical memory", false, 0},
		{sysconf.SC_NPROCESSORS_CONF, "number of processors configured", false, 0},
		{sysconf.SC_NPROCESSORS_ONLN, "number of processors online", false, 0},
	}

	for _, tc := range testCases {
		val, err := sysconf.Sysconf(tc.val)
		if err != nil {
			t.Fatalf("Sysconf(%s/%d): %v", tc.name, tc.val, err)
		}
		t.Logf("%s = %v", tc.name, val)
		if tc.expect && val != tc.expected {
			t.Errorf("Sysconf(%s/%d):expected %v, got %v", tc.name, tc.val, tc.expected, val)
		}
	}
}
