// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build !darwin && !dragonfly && !freebsd && !linux && !netbsd && !openbsd && !solaris

package sysconf

import (
	"fmt"
	"runtime"
)

func pathconf(path string, name int) (int64, error) {
	return -1, fmt.Errorf("unsupported on %s", runtime.GOOS)
}

func fpathconf(fd int, name int) (int64, error) {
	return -1, fmt.Errorf("unsupported on %s", runtime.GOOS)
}
