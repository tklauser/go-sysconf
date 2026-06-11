// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_cgotest

import "testing"

// The actual test functions are in the respective *_cgotest.go files so that they can use cgo.
func TestSysconfCgoMatch(t *testing.T)  { testSysconfCgoMatch(t) }
func TestPathconfCgoMatch(t *testing.T) { testPathconfCgoMatch(t) }
