// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"bufio"
	"encoding/binary"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"sync"

	"golang.org/x/sys/unix"
)

const (
	_AT_NULL   = 0  // End of auxiliary vector
	_AT_CLKTCK = 17 // Frequency of times()

	_SYSTEM_CLK_TCK = 100

	uintSize uint = 32 << (^uint(0) >> 63)
)

var (
	clktck     int64
	clktckOnce sync.Once
)

func getclktck() int64 {
	// I currently don't know a way to get the loaded-provided auxv, thus
	// get it from /proc/self/auxv on Linux.
	// Code based on cpu_linux.go in golang.org/x/sys/cpu
	buf, err := ioutil.ReadFile("/proc/self/auxv")
	if err == nil {
		pb := int(uintSize / 8)
		for i := 0; i < len(buf)-pb*2; i += pb * 2 {
			var tag, val uint
			switch uintSize {
			case 32:
				tag = uint(binary.LittleEndian.Uint32(buf[i:]))
				val = uint(binary.LittleEndian.Uint32(buf[i+pb:]))
			case 64:
				tag = uint(binary.LittleEndian.Uint64(buf[i:]))
				val = uint(binary.LittleEndian.Uint64(buf[i+pb:]))
			}

			switch tag {
			case _AT_CLKTCK:
				return int64(val)
			}
		}
	}
	return _SYSTEM_CLK_TCK
}

func readProcFsInt64(path string, fallback int64) int64 {
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

// based on readCPURange in github.com/iovisor/gobpf/pkg/cpuonline/cpu_range.go
func readCPURange(file string) (int64, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return 0, err
	}
	count := int64(0)
	cpuRangeStr := strings.Trim(string(buf), "\n ")
	for _, cpuRange := range strings.Split(cpuRangeStr, ",") {
		rangeOp := strings.SplitN(cpuRange, "-", 2)
		first, err := strconv.ParseUint(rangeOp[0], 10, 32)
		if err != nil {
			return 0, err
		}
		if len(rangeOp) == 1 {
			count++
			continue
		}
		last, err := strconv.ParseUint(rangeOp[1], 10, 32)
		if err != nil {
			return 0, err
		}
		count += int64(last - first + 1)
	}
	return count, nil
}

const sysfsCpuOnline = "/sys/devices/system/cpu/online"

func getNprocsSysfs() (int64, error) {
	return readCPURange(sysfsCpuOnline)
}

func getNprocsProcStat() (int64, error) {
	f, err := os.Open("/proc/stat")
	if err != nil {
		return -1, err
	}
	defer f.Close()

	count := int64(0)
	s := bufio.NewScanner(f)
	for s.Scan() {
		if line := strings.TrimSpace(s.Text()); strings.HasPrefix(line, "cpu") {
			l := strings.SplitN(line, " ", 2)
			_, err := strconv.ParseInt(l[0][3:], 10, 64)
			if err == nil {
				count++
			}
		} else {
			// The current format of /proc/stat has all the
			// cpu* lines at the beginning. Assume this
			// stays this way.
			break
		}
	}
	return count, nil
}

func getNprocs() int64 {
	count, err := getNprocsSysfs()
	if err == nil {
		return count
	}

	count, err = getNprocsProcStat()
	if err == nil {
		return count
	}

	// default to the value determined at runtime startup if all else fails
	return int64(runtime.NumCPU())
}

func getNprocsConf() int64 {
	// TODO(tk): read /sys/devices/system/cpu/present instead?
	d, err := os.Open("/sys/devices/system/cpu")
	if err == nil {
		defer d.Close()
		fis, err := d.Readdir(-1)
		if err == nil {
			count := int64(0)
			for _, fi := range fis {
				if name := fi.Name(); fi.IsDir() && strings.HasPrefix(name, "cpu") {
					_, err := strconv.ParseInt(name[3:], 10, 64)
					if err == nil {
						count++
					}
				}
			}
			return count
		}
	}

	// TODO(tk): fall back to reading /proc/cpuinfo on legacy systems
	// without sysfs?

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
		childMax := int64(-1)
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
	case SC_GETGR_R_SIZE_MAX:
		return _NSS_BUFLEN_GROUP, nil
	case SC_GETPW_R_SIZE_MAX:
		return _NSS_BUFLEN_PASSWD, nil
	case SC_MQ_OPEN_MAX:
		return -1, nil
	case SC_MQ_PRIO_MAX:
		return _MQ_PRIO_MAX, nil
	case SC_NGROUPS_MAX:
		return readProcFsInt64("/proc/sys/kernel/ngroups_max", _NGROUPS_MAX), nil
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
		return readProcFsInt64("/proc/sys/kernel/rtsig-max", _POSIX_SIGQUEUE_MAX), nil
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
		return -1, nil

	case SC_CPUTIME:
		if hasClock(unix.CLOCK_PROCESS_CPUTIME_ID) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
	case SC_MONOTONIC_CLOCK:
		if hasClock(unix.CLOCK_MONOTONIC) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
	case SC_SAVED_IDS:
		return _POSIX_SAVED_IDS, nil
	case SC_SPAWN:
		return _POSIX_SPAWN, nil
	case SC_SPIN_LOCKS:
		return _POSIX_SPIN_LOCKS, nil
	case SC_SPORADIC_SERVER:
		return _POSIX_SPORADIC_SERVER, nil
	case SC_SYNCHRONIZED_IO:
		return _POSIX_SYNCHRONIZED_IO, nil
	case SC_THREAD_ATTR_STACKADDR:
		return _POSIX_THREAD_ATTR_STACKADDR, nil
	case SC_THREAD_ATTR_STACKSIZE:
		return _POSIX_THREAD_ATTR_STACKSIZE, nil
	case SC_THREAD_CPUTIME:
		if hasClock(unix.CLOCK_THREAD_CPUTIME_ID) {
			return _POSIX_VERSION, nil
		}
		return -1, nil
	case SC_THREAD_PRIORITY_SCHEDULING:
		return _POSIX_THREAD_PRIORITY_SCHEDULING, nil
	case SC_THREAD_PROCESS_SHARED:
		return _POSIX_THREAD_PROCESS_SHARED, nil
	case SC_THREAD_SAFE_FUNCTIONS:
		return _POSIX_THREAD_SAFE_FUNCTIONS, nil
	case SC_THREAD_SPORADIC_SERVER:
		return _POSIX_THREAD_SPORADIC_SERVER, nil
	case SC_TRACE:
		return _POSIX_TRACE, nil
	case SC_TRACE_EVENT_FILTER:
		return _POSIX_TRACE_EVENT_FILTER, nil
	case SC_TRACE_EVENT_NAME_MAX:
		return -1, nil
	case SC_TRACE_INHERIT:
		return _POSIX_TRACE_INHERIT, nil
	case SC_TRACE_LOG:
		return _POSIX_TRACE_LOG, nil
	case SC_TRACE_NAME_MAX:
		return -1, nil
	case SC_TRACE_SYS_MAX:
		return -1, nil
	case SC_TRACE_USER_EVENT_MAX:
		return -1, nil
	case SC_TYPED_MEMORY_OBJECTS:
		return _POSIX_TYPED_MEMORY_OBJECTS, nil

	case SC_V7_ILP32_OFF32:
		return -1, nil
	case SC_V7_ILP32_OFFBIG:
		return -1, nil
	case SC_V7_LP64_OFF64:
		return _POSIX_V7_LP64_OFF64, nil
	case SC_V7_LPBIG_OFFBIG:
		return _POSIX_V7_LPBIG_OFFBIG, nil

	case SC_V6_ILP32_OFF32:
		return -1, nil
	case SC_V6_ILP32_OFFBIG:
		return -1, nil
	case SC_V6_LP64_OFF64:
		return _POSIX_V6_LP64_OFF64, nil
	case SC_V6_LPBIG_OFFBIG:
		return _POSIX_V6_LPBIG_OFFBIG, nil

	case SC_2_C_VERSION:
		return _POSIX2_C_VERSION, nil
	case SC_2_CHAR_TERM:
		return _POSIX2_CHAR_TERM, nil
	case SC_2_PBS,
		SC_2_PBS_ACCOUNTING,
		SC_2_PBS_CHECKPOINT,
		SC_2_PBS_LOCATE,
		SC_2_PBS_MESSAGE,
		SC_2_PBS_TRACK:
		return -1, nil
	case SC_2_UPE:
		return -1, nil

	case SC_XOPEN_CRYPT:
		return _XOPEN_CRYPT, nil
	case SC_XOPEN_ENH_I18N:
		return _XOPEN_ENH_I18N, nil
	case SC_XOPEN_REALTIME:
		return _XOPEN_REALTIME, nil
	case SC_XOPEN_REALTIME_THREADS:
		return _XOPEN_REALTIME_THREADS, nil
	case SC_XOPEN_SHM:
		return _XOPEN_SHM, nil
	case SC_XOPEN_STREAMS:
		return -1, nil
	case SC_XOPEN_UNIX:
		return _XOPEN_UNIX, nil
	case SC_XOPEN_VERSION:
		return _XOPEN_VERSION, nil
	case SC_XOPEN_XCU_VERSION:
		return _XOPEN_XCU_VERSION, nil

	case SC_PHYS_PAGES:
		return getPhysPages(), nil
	case SC_AVPHYS_PAGES:
		return getAvPhysPages(), nil
	case SC_NPROCESSORS_CONF:
		return getNprocsConf(), nil
	case SC_NPROCESSORS_ONLN:
		return getNprocs(), nil
	case SC_UIO_MAXIOV: // same as _SC_IOV_MAX
		return _UIO_MAXIOV, nil
	}

	return -1, errInvalid
}
