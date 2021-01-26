// Copyright 2021 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"syscall"
	"unsafe"
)

// TODO: remove sysconf wrapper once https://golang.org/cl/286593 is merged and x/sys updated.

//go:cgo_import_dynamic libc_sysconf sysconf "libc.so"
//go:linkname procSysconf libc_sysconf

var procSysconf uintptr

//go:linkname sysvicall6 syscall.sysvicall6
func sysvicall6(trap, nargs, a1, a2, a3, a4, a5, a6 uintptr) (r1, r2 uintptr, err syscall.Errno)

func sysconf(name int) (int64, error) {
	n, _, errno := sysvicall6(uintptr(unsafe.Pointer(&procSysconf)), 1, uintptr(name), 0, 0, 0, 0, 0)
	if errno != 0 {
		return -1, errInvalid
	}
	return int64(n), nil
}
