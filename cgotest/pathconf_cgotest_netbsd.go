// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_cgotest

/*
#include <unistd.h>
*/
import "C"

import "github.com/tklauser/go-sysconf"

var pathconfTestCases = []testCase{
	{sysconf.PC_SYMLINK_MAX, C._PC_SYMLINK_MAX, "SYMLINK_MAX"},
	{sysconf.PC_SYNC_IO, C._PC_SYNC_IO, "SYNC_IO"},
}
