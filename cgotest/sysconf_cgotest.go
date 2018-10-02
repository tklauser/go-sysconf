// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build darwin freebsd linux netbsd

package sysconf_cgotest

import "C"

type testCase struct {
	goVar int
	cVar  C.int
	name  string
}
