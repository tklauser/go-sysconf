// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package sysconf

/*
#include <limits.h>
#include <sys/param.h>
#include <time.h>
#include <unistd.h>
*/
import "C"

// sysconf variables
const (
	// POSIX.1 variables
	SC_ARG_MAX        = C._SC_ARG_MAX
	SC_CHILD_MAX      = C._SC_CHILD_MAX
	SC_HOST_NAME_MAX  = C._SC_HOST_NAME_MAX
	SC_LOGIN_NAME_MAX = C._SC_LOGIN_NAME_MAX
	SC_NGROUPS_MAX    = C._SC_NGROUPS_MAX
	SC_CLK_TCK        = C._SC_CLK_TCK
	SC_OPEN_MAX       = C._SC_OPEN_MAX
	SC_PAGESIZE       = C._SC_PAGESIZE
	SC_RE_DUP_MAX     = C._SC_RE_DUP_MAX
	SC_STREAM_MAX     = C._SC_STREAM_MAX
	SC_SYMLOOP_MAX    = C._SC_SYMLOOP_MAX
	SC_TTY_NAME_MAX   = C._SC_TTY_NAME_MAX
	SC_TZNAME_MAX     = C._SC_TZNAME_MAX
	SC_VERSION        = C._SC_VERSION

	// POSIX.2 variables (limits for utilities)
	SC_BC_BASE_MAX      = C._SC_BC_BASE_MAX
	SC_BC_DIM_MAX       = C._SC_BC_DIM_MAX
	SC_BC_SCALE_MAX     = C._SC_BC_SCALE_MAX
	SC_BC_STRING_MAX    = C._SC_BC_STRING_MAX
	SC_COLL_WEIGHTS_MAX = C._SC_COLL_WEIGHTS_MAX
	SC_EXPR_NEST_MAX    = C._SC_EXPR_NEST_MAX
	SC_LINE_MAX         = C._SC_LINE_MAX
	SC_2_VERSION        = C._SC_2_VERSION
	SC_2_C_DEV          = C._SC_2_C_DEV
	SC_2_FORT_DEV       = C._SC_2_FORT_DEV
	SC_2_FORT_RUN       = C._SC_2_FORT_RUN
	SC_2_LOCALEDEF      = C._SC_2_LOCALEDEF
	SC_2_SW_DEV         = C._SC_2_SW_DEV

	// non-standard variables
	SC_PHYS_PAGES       = C._SC_PHYS_PAGES
	SC_NPROCESSORS_CONF = C._SC_NPROCESSORS_CONF
	SC_NPROCESSORS_ONLN = C._SC_NPROCESSORS_ONLN
)

// sysconf values
const (
	_CLK_TCK        = C.CLK_TCK
	_MAXHOSTNAMELEN = C.MAXHOSTNAMELEN
	_MAXLOGNAME     = C.MAXLOGNAME
	_MAXSYMLINKS    = C.MAXSYMLINKS

	_POSIX_VERSION  = C._POSIX_VERSION
	_POSIX2_VERSION = C._POSIX2_VERSION

	_POSIX2_C_DEV = C._POSIX2_C_DEV

	_POSIX_ARG_MAX   = C._POSIX_ARG_MAX
	_POSIX_CHILD_MAX = C._POSIX_CHILD_MAX

	_NGROUPS_MAX = C.NGROUPS_MAX
	_RE_DUP_MAX  = C.RE_DUP_MAX

	_BC_BASE_MAX      = C.BC_BASE_MAX
	_BC_DIM_MAX       = C.BC_DIM_MAX
	_BC_SCALE_MAX     = C.BC_SCALE_MAX
	_BC_STRING_MAX    = C.BC_STRING_MAX
	_COLL_WEIGHTS_MAX = C.COLL_WEIGHTS_MAX
	_EXPR_NEST_MAX    = C.EXPR_NEST_MAX
	_LINE_MAX         = C.LINE_MAX
)
