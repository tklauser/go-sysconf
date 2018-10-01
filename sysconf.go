// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sysconf implements the sysconf(3) function and provides the
// associated SC_* constants to query system configuration values.
package sysconf

import (
	"golang.org/x/sys/unix"
)

//go:generate go run mksysconf.go

// Sysconf returns the value of a sysconf(3) configurable system variable.
func Sysconf(name int) (int64, error) {
	// OS-specific sysconf
	if sc, err := sysconf(name); err == nil {
		return sc, nil
	}

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
	case SC_HOST_NAME_MAX:
		return _HOST_NAME_MAX, nil
	case SC_LINE_MAX:
		return _LINE_MAX, nil
	case SC_LOGIN_NAME_MAX:
		return _LOGIN_NAME_MAX, nil
	case SC_PAGESIZE: // same as SC_PAGE_SIZE
		return int64(unix.Getpagesize()), nil
	case SC_RE_DUP_MAX:
		return _RE_DUP_MAX, nil
	case SC_SYMLOOP_MAX:
		return _SYMLOOP_MAX, nil

	case SC_ADVISORY_INFO:
		return _POSIX_ADVISORY_INFO, nil
	case SC_ASYNCHRONOUS_IO:
		return _POSIX_ASYNCHRONOUS_IO, nil
	case SC_BARRIERS:
		return _POSIX_BARRIERS, nil
	case SC_CLOCK_SELECTION:
		return _POSIX_CLOCK_SELECTION, nil
	case SC_CPUTIME:
		return _POSIX_CPUTIME, nil
	case SC_FSYNC:
		return _POSIX_FSYNC, nil
	case SC_IPV6:
		return _POSIX_IPV6, nil
	case SC_JOB_CONTROL:
		return _POSIX_JOB_CONTROL, nil
	case SC_MAPPED_FILES:
		return _POSIX_MAPPED_FILES, nil
	case SC_MEMLOCK:
		return _POSIX_MEMLOCK, nil
	case SC_MEMLOCK_RANGE:
		return _POSIX_MEMLOCK_RANGE, nil
	case SC_MONOTONIC_CLOCK:
		return _POSIX_MONOTONIC_CLOCK, nil
	case SC_SEMAPHORES:
		return _POSIX_SEMAPHORES, nil
	case SC_SHELL:
		return _POSIX_SHELL, nil
	case SC_THREADS:
		return _POSIX_THREADS, nil
	case SC_TIMEOUTS:
		return _POSIX_TIMEOUTS, nil
	case SC_VERSION:
		return _POSIX_VERSION, nil

	case SC_2_C_DEV:
		return _POSIX2_C_DEV, nil
	case SC_2_FORT_DEV:
		return -1, nil
	case SC_2_FORT_RUN:
		return -1, nil
	case SC_2_VERSION:
		return _POSIX2_VERSION, nil
	}

	return -1, unix.EINVAL
}
