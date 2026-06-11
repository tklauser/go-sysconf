// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

// Pathconf returns the value of a pathconf(3) configurable pathname variable.
// The name parameter should be a PC_* constant defined in this package. The
// implementation is GOOS-specific and certain PC_* constants might not be
// defined for all GOOSes.
func Pathconf(path string, name int) (int64, error) {
	return pathconf(path, name)
}

// Fpathconf returns the value of an fpathconf(3) configurable pathname
// variable for an open file descriptor. The name parameter should be a PC_*
// constant defined in this package.
func Fpathconf(fd int, name int) (int64, error) {
	return fpathconf(fd, name)
}
