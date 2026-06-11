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
	{sysconf.PC_ASYNC_IO, C._PC_ASYNC_IO, "ASYNC_IO"},
}
