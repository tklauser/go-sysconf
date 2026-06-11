// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build darwin || dragonfly || freebsd || netbsd || openbsd || solaris

package sysconf

import "golang.org/x/sys/unix"

func pathconf(path string, name int) (int64, error) {
	val, err := unix.Pathconf(path, name)
	if err != nil {
		return -1, err
	}
	return int64(val), nil
}

func fpathconf(fd int, name int) (int64, error) {
	val, err := unix.Fpathconf(fd, name)
	if err != nil {
		return -1, err
	}
	return int64(val), nil
}
