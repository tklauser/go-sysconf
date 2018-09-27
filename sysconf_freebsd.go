// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"golang.org/x/sys/unix"
)

const (
	_HOST_NAME_MAX  = _MAXHOSTNAMELEN
	_LOGIN_NAME_MAX = _MAXLOGNAME
	_SYMLOOP_MAX    = _MAXSYMLINKS
)

func sysconf(name int) (int64, error) {
	switch name {
	case SC_ARG_MAX:
		return sysctl32("kern.argmax")
	case SC_CHILD_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NPROC, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return rlim.Cur, nil
			}
		}
		return -1, nil
	case SC_STREAM_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return rlim.Cur, nil
			}
		}
		return -1, nil
	case SC_TTY_NAME_MAX:
		return pathconf(_PATH_DEV, _PC_NAME_MAX), nil
	case SC_TZNAME_MAX:
		return pathconf(_PATH_ZONEINFO, _PC_NAME_MAX), nil
	case SC_CLK_TCK:
		return _CLK_TCK, nil
	case SC_PHYS_PAGES:
		if val, err := unix.SysctlUint64("hw.availpages"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	case SC_NPROCESSORS_CONF:
		fallthrough
	case SC_NPROCESSORS_ONLN:
		if val, err := unix.SysctlUint32("hw.ncpu"); err == nil {
			return int64(val), nil
		}
		return -1, nil
	}

	return -1, unix.EINVAL
}
