// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"sync"

	"golang.org/x/sys/unix"
)

const (
	_HOST_NAME_MAX  = _MAXHOSTNAMELEN - 1
	_LOGIN_NAME_MAX = _MAXLOGNAME
	_SYMLOOP_MAX    = _MAXSYMLINKS
)

var (
	clktck     int64
	clktckOnce sync.Once
)

// sysconf implements sysconf(3) as in the Darwin libc, version 1244.30.3
// (derived from the FreeBSD libc).
func sysconf(name int) (int64, error) {
	switch name {
	case SC_AIO_LISTIO_MAX:
		fallthrough
	case SC_AIO_MAX:
		return sysctl32("kern.aiomax"), nil
	case SC_AIO_PRIO_DELTA_MAX:
		return -1, nil
	case SC_ARG_MAX:
		return sysctl32("kern.argmax"), nil
	case SC_ATEXIT_MAX:
		return _INT_MAX, nil
	case SC_CHILD_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NPROC, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return int64(rlim.Cur), nil
			}
		}
		return -1, nil
	case SC_CLK_TCK:
		return _CLK_TCK, nil
	case SC_DELAYTIMER_MAX:
		return -1, nil
	case SC_IOV_MAX:
		return _IOV_MAX, nil
	case SC_MQ_OPEN_MAX:
		return -1, nil
	case SC_MQ_PRIO_MAX:
		return -1, nil
	case SC_NGROUPS_MAX:
		return sysctl32("kern.ngroups"), nil
	case SC_OPEN_MAX, SC_STREAM_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err == nil {
			if rlim.Cur != unix.RLIM_INFINITY {
				return int64(rlim.Cur), nil
			}
		}
		return -1, nil
	case SC_RTSIG_MAX:
		return -1, nil
	case SC_SEM_NSEMS_MAX:
		return sysctl32("kern.sysv.semmns"), nil
	case SC_SEM_VALUE_MAX:
		return _POSIX_SEM_VALUE_MAX, nil
	case SC_SIGQUEUE_MAX:
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
		return -1, nil
	case SC_TTY_NAME_MAX:
		// should be _PATH_DEV instead of "/"
		return pathconf("/", _PC_NAME_MAX), nil
	case SC_TZNAME_MAX:
		return pathconf(_PATH_ZONEINFO, _PC_NAME_MAX), nil

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
	case SC_SAVED_IDS:
		return sysctl32("kern.saved_ids"), nil
	case SC_SEMAPHORES:
		return _POSIX_SEMAPHORES, nil
	case SC_THREADS:
		return _POSIX_THREADS, nil
	case SC_TIMEOUTS:
		return _POSIX_TIMEOUTS, nil
	case SC_TIMERS:
		return _POSIX_TIMERS, nil
	case SC_VERSION:
		// TODO(tk): darwin libc uses sysctl(CTL_KERN, KERN_POSIX1)
		return _POSIX_VERSION, nil

	case SC_XOPEN_CRYPT:
		return _XOPEN_CRYPT, nil
	case SC_XOPEN_REALTIME:
		return _XOPEN_REALTIME, nil
	case SC_XOPEN_STREAMS:
		return -1, nil
	case SC_XOPEN_UNIX:
		return _XOPEN_UNIX, nil
	case SC_XOPEN_VERSION:
		return _XOPEN_VERSION, nil
	case SC_XOPEN_XCU_VERSION:
		return _XOPEN_XCU_VERSION, nil

	case SC_2_VERSION:
		return _POSIX2_VERSION, nil
	case SC_2_C_DEV:
		return _POSIX2_C_DEV, nil
	case SC_2_LOCALEDEF:
		return _POSIX2_LOCALEDEF, nil
	case SC_2_SW_DEV:
		return _POSIX2_SW_DEV, nil
	case SC_2_UPE:
		return _POSIX2_UPE, nil

	case SC_PHYS_PAGES:
		return sysctl64("hw.memsize") / int64(unix.Getpagesize()), nil
	case SC_MONOTONIC_CLOCK:
		return _POSIX_MONOTONIC_CLOCK, nil
	case SC_NPROCESSORS_CONF:
		fallthrough
	case SC_NPROCESSORS_ONLN:
		return sysctl32("hw.ncpu"), nil
	}

	return -1, unix.EINVAL
}
