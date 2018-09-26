// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run mksysconf.go

package sysconf

import (
	"golang.org/x/sys/unix"
)

func Sysconf(name int) (int64, error) {
	// OS-specific sysconf
	if sc, err := sysconf(name); err == nil {
		return sc, nil
	}

	// POSIX.1 sysconf
	switch name {
	case SC_HOST_NAME_MAX:
		return _HOST_NAME_MAX, nil
	case SC_LOGIN_NAME_MAX:
		return _LOGIN_NAME_MAX, nil
	case SC_NGROUPS_MAX:
		return _NGROUPS_MAX, nil
	case SC_PAGESIZE:
		return int64(unix.Getpagesize()), nil
	case SC_RE_DUP_MAX:
		return _RE_DUP_MAX, nil
	case SC_STREAM_MAX:
		return _STREAM_MAX, nil
	case SC_SYMLOOP_MAX:
		return _SYMLOOP_MAX, nil
	case SC_TTY_NAME_MAX:
		return _TTY_NAME_MAX, nil
	case SC_TZNAME_MAX:
		return _TZNAME_MAX, nil
	case SC_VERSION:
		return _POSIX_VERSION, nil
	}

	// POSIX.2 sysconf
	switch name {
	case SC_BC_BASE_MAX:
		return _BC_BASE_MAX, nil
	case SC_BC_DIM_MAX:
		return _BC_DIM_MAX, nil
	case SC_BC_SCALE_MAX:
		return _BC_SCALE_MAX, nil
	case SC_BC_STRING_MAX:
		return _BC_STRING_MAX, nil
	case SC_COLL_WEIGHTS_MAX:
		return _COLL_WEIGHTS_MAX, nil
	case SC_EXPR_NEST_MAX:
		return _EXPR_NEST_MAX, nil
	case SC_LINE_MAX:
		return _LINE_MAX, nil
	case SC_2_VERSION:
		return _POSIX2_VERSION, nil
	case SC_2_C_DEV:
		return _POSIX2_C_DEV, nil
	case SC_2_C_VERSION:
		return _POSIX2_C_VERSION, nil
	case SC_2_FORT_DEV:
		return -1, nil
	case SC_2_FORT_RUN:
		return -1, nil
	}

	return -1, unix.EINVAL
}
