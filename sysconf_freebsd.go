// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"golang.org/x/sys/unix"
)

const (
	_HOST_NAME_MAX  = _MAXHOSTNAMELEN
	_LOGIN_NAME_MAX = MAXLOGNAME
	_SYMLOOP_MAX    = _MAXSYMLINKS
)

func sysconf(name int) (int64, error) {
	switch name {
	case SC_CLK_TCK:
		return _CLK_TCK, nil
	}

	return -1, unix.EINVAL
}
