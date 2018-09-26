// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"os"
	"runtime"
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

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func sysconf(name int) (int64, error) {
	switch name {
	case SC_ARG_MAX:
		argMax := int64(_POSIX_ARG_MAX)
		var rlim unix.Rlimit
		if err := unix.Getrlimit(unix.RLIMIT_STACK, &rlim); err == nil {
			argMax = max(argMax, int64(rlim.Cur/4))
		}
		return argMax, nil
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
	case SC_PHYS_PAGES:
		return getPhysPages(), nil
	case SC_AVPHYS_PAGES:
		return getAvPhysPages(), nil
	case SC_NPROCESSORS_CONF:
		return getNprocsConf(), nil
	case SC_NPROCESSORS_ONLN:
		return getNprocs(), nil
	}

	return -1, unix.EINVAL
}
