// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf_cgotest

import "testing"

// The actual test functions are defined in the respective {sys,path}conf_cgotest_*.go files so that
// they can use cgo.

func TestSysconfCgoMatch(t *testing.T)        { testSysconfCgoMatch(t) }
func TestSysconfCgoMatchInvalid(t *testing.T) { testSysconfCgoMatchInvalid(t) }
func TestPathconfCgoMatch(t *testing.T)       { testPathconfCgoMatch(t) }
