// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"os"
	"sync"

	"golang.org/x/sys/unix"
)

const (
	_HOST_NAME_MAX  = _MAXHOSTNAMELEN
	_LOGIN_NAME_MAX = _MAXLOGNAME
	_SYMLOOP_MAX    = _MAXSYMLINKS

	_POSIX2_C_DEV = -1
)

var (
	clktck     int64
	clktckOnce sync.Once
)

func pathconf(path string, name int) (int64, error) {
	val, err := unix.Pathconf(path, name)
	return int64(val), err
}

func sysconf(name int) (int64, error) {
	switch name {
	case SC_ARG_MAX:
		if val, err := unix.SysctlUint32("kern.argmax"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	case SC_CHILD_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NPROC, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return int64(rlim.Cur), nil
			}
		}
		return -1, nil
	case SC_STREAM_MAX:
		if val, err := unix.SysctlUint32("user.stream_max"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	case SC_TTY_NAME_MAX:
		return pathconf(_PATH_DEV, _PC_NAME_MAX)
	case SC_TZNAME_MAX:
		return pathconf(_PATH_ZONEINFO, _PC_NAME_MAX)
	case SC_CLK_TCK:
		clktckOnce.Do(func() {
			clktck = -1
			if val, err := unix.SysctlUint32("kern.clockrate"); err == nil {
				clktck = int64(val)
			}
		})
		return clktck, nil
	case SC_PHYS_PAGES:
		if val, err := unix.SysctlUint64("hw.physmem64"); err == nil {
			return int64(val / uint64(os.Getpagesize())), nil
		}
		return -1, nil
	case SC_NPROCESSORS_CONF:
		if val, err := unix.SysctlUint32("hw.ncpu"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	case SC_NPROCESSORS_ONLN:
		if val, err := unix.SysctlUint32("hw.ncpuonline"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	}

	return -1, unix.EINVAL
}
