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
	{sysconf.PC_FILESIZEBITS, C._PC_FILESIZEBITS, "FILESIZEBITS"},
	{sysconf.PC_2_SYMLINKS, C._PC_2_SYMLINKS, "2_SYMLINKS"},
	{sysconf.PC_SYMLINK_MAX, C._PC_SYMLINK_MAX, "SYMLINK_MAX"},
	{sysconf.PC_SYNC_IO, C._PC_SYNC_IO, "SYNC_IO"},
	{sysconf.PC_TIMESTAMP_RESOLUTION, C._PC_TIMESTAMP_RESOLUTION, "TIMESTAMP_RESOLUTION"},
}
