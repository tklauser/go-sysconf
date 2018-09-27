// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"sync"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	_AT_NULL   = 0  // End of auxiliary vector
	_AT_CLKTCK = 17 // Frequency of times()

	_SYSTEM_CLK_TCK = 100
)

var (
	clktck     int64
	clktckOnce sync.Once
)

func getclktck() int64 {
	// I currently don't know a way to get the loaded-provided auxv, thus
	// get it from /proc/self/auxv on Linux.
	// Code based on sysargs in runtime/os_linux.go
	if fd, err := unix.Open("/proc/self/auxv", unix.O_RDONLY, 0); err == nil {
		buf := make([]byte, 8192)
		n, err := unix.Read(fd, buf)
		unix.Close(fd)
		if err == nil {
			auxv := (*[10000]uintptr)(unsafe.Pointer(&buf[0]))[0:n]
			// Make sure buf is terminated, even if we didn't read
			// the whole file.
			auxv[len(auxv)-2] = _AT_NULL
			for i := 0; auxv[i] != _AT_NULL; i += 2 {
				tag, val := auxv[i], auxv[i+1]
				switch tag {
				case _AT_CLKTCK:
					return int64(val)
				}
			}
		}
	}
	return _SYSTEM_CLK_TCK
}

func readProcFs(path string, fallback int64) int64 {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return fallback
	}
	i, err := strconv.ParseInt(string(data[:len(data)-1]), 0, 64)
	if err != nil {
		return fallback
	}
	return i
}

// getMemPages computes mem*unit/os.Getpagesize(), but avoids overflowing int64.
func getMemPages(mem uint64, unit uint32) int64 {
	pageSize := os.Getpagesize()
	for unit > 1 && pageSize > 1 {
		unit >>= 1
		pageSize >>= 1
	}
	mem *= uint64(unit)
	for pageSize > 1 {
		pageSize >>= 1
		mem >>= 1
	}
	return int64(mem)
}

func getPhysPages() int64 {
	var si unix.Sysinfo_t
	err := unix.Sysinfo(&si)
	if err != nil {
		return int64(0)
	}
	return getMemPages(si.Totalram, si.Unit)
}

func getAvPhysPages() int64 {
	var si unix.Sysinfo_t
	err := unix.Sysinfo(&si)
	if err != nil {
		return int64(0)
	}
	return getMemPages(si.Freeram, si.Unit)
}

func getNprocs() int64 {
	// TODO(tk): parse /sys/devices/system/cpu/online like glibc

	runtime.LockOSThread()
	defer runtime.UnlockOSThread()

	var cpus unix.CPUSet
	err := unix.SchedGetaffinity(0, &cpus)
	if err == nil {
		return int64(cpus.Count())
	}

	return int64(runtime.NumCPU()) // default to the value determined at runtime startup if all else fails
}

func getNprocsConf() int64 {
	// TODO(tk): walk /sys/devices/system/cpu/cpu* entries like glibc

	return getNprocs()
}

func hasClock(clockid int32) bool {
	var res unix.Timespec
	if err := unix.ClockGetres(clockid, &res); err != nil {
		return false
	}
	return true
}

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func sysconf(name int) (int64, error) {
	switch name {
	case SC_AIO_LISTIO_MAX:
		return -1, nil
	case SC_AIO_MAX:
		return -1, nil
	case SC_AIO_PRIO_DELTA_MAX:
		return _AIO_PRIO_DELTA_MAX, nil
	case SC_ARG_MAX:
		argMax := int64(_POSIX_ARG_MAX)
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_STACK, &rlim); err == nil {
			argMax = max(argMax, int64(rlim.Cur/4))
		}
		return argMax, nil
	case SC_ATEXIT_MAX:
		return _INT_MAX, nil
	case SC_CHILD_MAX:
		childMax := int64(_POSIX_CHILD_MAX)
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NPROC, &rlim); err == nil && rlim.Cur != unix.RLIM_INFINITY {
			childMax = int64(rlim.Cur)
		}
		return childMax, nil
	case SC_CLK_TCK:
		clktckOnce.Do(func() { clktck = getclktck() })
		return clktck, nil
	case SC_DELAYTIMER_MAX:
		return _DELAYTIMER_MAX, nil
	case SC_MQ_OPEN_MAX:
		return -1, nil
	case SC_MQ_PRIO_MAX:
		return _MQ_PRIO_MAX, nil
	case SC_NGROUPS_MAX:
		return readProcFs("/proc/sys/kernel/ngroups_max", _NGROUPS_MAX), nil
	case SC_OPEN_MAX:
		openMax := int64(_OPEN_MAX)
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_NOFILE, &rlim); err == nil {
			openMax = int64(rlim.Cur)
		}
		return openMax, nil
	case SC_RTSIG_MAX:
		return _RTSIG_MAX, nil
	case SC_SEM_NSEMS_MAX:
		return -1, nil
	case SC_SEM_VALUE_MAX:
		return _SEM_VALUE_MAX, nil
	case SC_SIGQUEUE_MAX:
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_SIGPENDING, &rlim); err == nil {
			return int64(rlim.Cur), nil
		}
		return readProcFs("/proc/sys/kernel/rtsig-max", _POSIX_SIGQUEUE_MAX), nil
	case SC_STREAM_MAX:
		return _STREAM_MAX, nil
	case SC_THREAD_DESTRUCTOR_ITERATIONS:
		return _POSIX_THREAD_DESTRUCTOR_ITERATIONS, nil
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
		return _TTY_NAME_MAX, nil
	case SC_TZNAME_MAX:
		return _TZNAME_MAX, nil

	case SC_ADVISORY_INFO:
		return _POSIX_ADVISORY_INFO, nil
	case SC_ASYNCHRONOUS_IO:
		return _POSIX_ASYNCHRONOUS_IO, nil
	case SC_BARRIERS:
		return _POSIX_BARRIERS, nil
	case SC_CLOCK_SELECTION:
		return _POSIX_CLOCK_SELECTION, nil
	case SC_CPUTIME:
		if hasClock(unix.CLOCK_PROCESS_CPUTIME_ID) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
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
		return _POSIX_SAVED_IDS, nil
	case SC_SEMAPHORES:
		return _POSIX_SEMAPHORES, nil
	case SC_SHELL:
		return _POSIX_SHELL, nil
	case SC_THREAD_CPUTIME:
		if hasClock(unix.CLOCK_THREAD_CPUTIME_ID) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
	case SC_THREADS:
		return _POSIX_THREADS, nil
	case SC_TIMEOUTS:
		return _POSIX_TIMEOUTS, nil
	case SC_TIMERS:
		return _POSIX_TIMERS, nil

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

	case SC_2_C_VERSION:
		return _POSIX2_C_VERSION, nil
	case SC_2_UPE:
		return -1, nil
	case SC_2_VERSION:
		return _POSIX2_VERSION, nil
	case SC_2_C_DEV:
		return _POSIX2_C_DEV, nil
	case SC_2_LOCALEDEF:
		return _POSIX2_LOCALEDEF, nil
	case SC_2_SW_DEV:
		return _POSIX2_SW_DEV, nil

	case SC_PHYS_PAGES:
		return getPhysPages(), nil
	case SC_AVPHYS_PAGES:
		return getAvPhysPages(), nil
	case SC_MONOTONIC_CLOCK:
		if hasClock(unix.CLOCK_MONOTONIC) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
	case SC_NPROCESSORS_CONF:
		return getNprocsConf(), nil
	case SC_NPROCESSORS_ONLN:
		return getNprocs(), nil
	case SC_UIO_MAXIOV: // same as _SC_IOV_MAX
		return _UIO_MAXIOV, nil
	}

	return -1, unix.EINVAL
}
