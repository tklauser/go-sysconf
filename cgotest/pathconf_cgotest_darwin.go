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
	{sysconf.PC_CHOWN_RESTRICTED, C._PC_CHOWN_RESTRICTED, "CHOWN_RESTRICTED"},
	{sysconf.PC_NO_TRUNC, C._PC_NO_TRUNC, "NO_TRUNC"},
	{sysconf.PC_2_SYMLINKS, C._PC_2_SYMLINKS, "2_SYMLINKS"},
	{sysconf.PC_ALLOC_SIZE_MIN, C._PC_ALLOC_SIZE_MIN, "ALLOC_SIZE_MIN"},
	{sysconf.PC_ASYNC_IO, C._PC_ASYNC_IO, "ASYNC_IO"},
	{sysconf.PC_FILESIZEBITS, C._PC_FILESIZEBITS, "FILESIZEBITS"},
	{sysconf.PC_PRIO_IO, C._PC_PRIO_IO, "PRIO_IO"},
	{sysconf.PC_REC_INCR_XFER_SIZE, C._PC_REC_INCR_XFER_SIZE, "REC_INCR_XFER_SIZE"},
	{sysconf.PC_REC_MAX_XFER_SIZE, C._PC_REC_MAX_XFER_SIZE, "REC_MAX_XFER_SIZE"},
	{sysconf.PC_REC_MIN_XFER_SIZE, C._PC_REC_MIN_XFER_SIZE, "REC_MIN_XFER_SIZE"},
	{sysconf.PC_REC_XFER_ALIGN, C._PC_REC_XFER_ALIGN, "REC_XFER_ALIGN"},
	{sysconf.PC_SYMLINK_MAX, C._PC_SYMLINK_MAX, "SYMLINK_MAX"},
	{sysconf.PC_SYNC_IO, C._PC_SYNC_IO, "SYNC_IO"},
}
