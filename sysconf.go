// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package sysconf implements the sysconf(3), pathconf(3), and fpathconf(3)
// functions and provides the associated constants to query system configuration
// values.
package sysconf

//go:generate go run mksysconf.go

// Sysconf returns the value of a sysconf(3) runtime system parameter.
// The name parameter should be a SC_* constant defined in this package. The
// implementation is GOOS-specific and certain SC_* constants might not be
// defined for all GOOSes.
func Sysconf(name int) (int64, error) {
	return sysconf(name)
}
