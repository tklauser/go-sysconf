// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"golang.org/x/sys/unix"
)

const (
	_HOST_NAME_MAX  = _MAXHOSTNAMELEN - 1
	_LOGIN_NAME_MAX = _MAXLOGNAME
	_SYMLOOP_MAX    = _MAXSYMLINKS
)

func sysconf(name int) (int64, error) {
	switch name {
	case SC_AIO_LISTIO_MAX:
		return sysctl32("p1003_1b.aio_listio_max"), nil
	case SC_AIO_MAX:
		return sysctl32("p1003_1b.aio_max"), nil
	case SC_AIO_PRIO_DELTA_MAX:
		return sysctl32("p1003_1b.aio_prio_delta_max"), nil
	case SC_ARG_MAX:
		return sysctl32("kern.argmax"), nil
	case SC_ATEXIT_MAX:
		return _ATEXIT_SIZE, nil
	case SC_CHILD_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NPROC, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return rlim.Cur, nil
			}
		}
		return -1, nil
	case SC_CLK_TCK:
		return _CLK_TCK, nil
	case SC_DELAYTIMER_MAX:
		return sysctl32("p1003_1b.delaytimer_max"), nil
	case SC_IOV_MAX:
		return sysctl32("kern.iov_max"), nil
	case SC_MQ_OPEN_MAX:
		return yesno(sysctl32("p1003_1b.mq_open_max")), nil
	case SC_MQ_PRIO_MAX:
		return _MQ_PRIO_MAX, nil
	case SC_NGROUPS_MAX:
		return sysctl32("kern.ngroups"), nil
	case SC_OPEN_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return rlim.Cur, nil
			}
		}
		return -1, nil
	case SC_RTSIG_MAX:
		return sysctl32("p1003_1b.rtsig_max"), nil
	case SC_SEM_NSEMS_MAX:
		return -1, nil
	case SC_SEM_VALUE_MAX:
		return _SEM_VALUE_MAX, nil
	case SC_SIGQUEUE_MAX:
		return sysctl32("p1003_1b.sigqueue_max"), nil
	case SC_STREAM_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return rlim.Cur, nil
			}
		}
		return -1, nil
	case SC_THREAD_DESTRUCTOR_ITERATIONS:
		return _PTHREAD_DESTRUCTOR_ITERATIONS, nil
	case SC_THREAD_KEYS_MAX:
		return _PTHREAD_KEYS_MAX, nil
	case SC_THREAD_PRIO_INHERIT:
		return _POSIX_THREAD_PRIO_INHERIT, nil
	case SC_THREAD_PRIO_PROTECT:
		return _POSIX_THREAD_PRIO_PROTECT, nil
	case SC_THREAD_STACK_MIN:
		return _PTHREAD_STACK_MIN, nil
	case SC_THREAD_THREADS_MAX:
		return -1, nil
	case SC_TIMER_MAX:
		return yesno(sysctl32("p1003_1b.timer_max")), nil
	case SC_TTY_NAME_MAX:
		return pathconf(_PATH_DEV, _PC_NAME_MAX), nil
	case SC_TZNAME_MAX:
		return pathconf(_PATH_ZONEINFO, _PC_NAME_MAX), nil

	case SC_IPV6:
		if _POSIX_IPV6 == 0 {
			fd, err := unix.Socket(unix.AF_INET6, unix.SOCK_DGRAM, 0)
			if err == nil && fd >= 0 {
				unix.Close(fd)
				return int64(200112), nil
			}
			return 0, nil
		}
		return _POSIX_IPV6, nil
	case SC_PRIORITY_SCHEDULING:
		if _POSIX_PRIORITY_SCHEDULING == 0 {
			return yesno(sysctl32("p1003_1b.priority_scheduling")), nil
		}
	case SC_SAVED_IDS:
		return yesno(sysctl32("kern.saved_ids")), nil
	case SC_SEMAPHORES:
		if _POSIX_SEMAPHORES == 0 {
			return yesno(sysctl32("p1003_1b.semaphores")), nil
		}
		return _POSIX_SEMAPHORES, nil
	case SC_THREAD_CPUTIME:
		return _POSIX_THREAD_CPUTIME, nil
	case SC_TIMERS:
		if _POSIX_TIMERS == 0 {
			return yesno(sysctl32("p1003_1b.timers")), nil
		}
		return _POSIX_TIMERS, nil

	case SC_2_UPE:
		return _POSIX2_UPE, nil

	case SC_XOPEN_CRYPT:
		return _XOPEN_CRYPT, nil
	case SC_XOPEN_REALTIME:
		return _XOPEN_REALTIME, nil
	case SC_XOPEN_STREAMS:
		return -1, nil
	case SC_XOPEN_UNIX:
		return _XOPEN_UNIX, nil

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
