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
	{sysconf.PC_MAX_INPUT, C._PC_MAX_INPUT, "MAX_INPUT"},
	{sysconf.PC_2_SYMLINKS, C._PC_2_SYMLINKS, "2_SYMLINKS"},
	{sysconf.PC_REC_INCR_XFER_SIZE, C._PC_REC_INCR_XFER_SIZE, "REC_INCR_XFER_SIZE"},
	{sysconf.PC_REC_MAX_XFER_SIZE, C._PC_REC_MAX_XFER_SIZE, "REC_MAX_XFER_SIZE"},
	{sysconf.PC_REC_MIN_XFER_SIZE, C._PC_REC_MIN_XFER_SIZE, "REC_MIN_XFER_SIZE"},
	{sysconf.PC_REC_XFER_ALIGN, C._PC_REC_XFER_ALIGN, "REC_XFER_ALIGN"},
}
