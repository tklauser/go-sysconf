// Copyright 2026 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package sysconf

import (
	"errors"

	"golang.org/x/sys/unix"
)

const (
	linuxLinkMax      = 127
	linuxLongLinkMax  = 65000
	linuxSymlinks     = 1
	linuxFileSizeBits = 64
	linuxRecXferAlign = 4096
	linuxRecMinXfer   = 4096
	linuxAllocSizeMin = 4096
)

func pathconfConst(name int) (int64, error) {
	switch name {
	case PC_PATH_MAX:
		return _PATH_MAX, nil
	case PC_MAX_CANON:
		return _MAX_CANON, nil
	case PC_MAX_INPUT:
		return _MAX_INPUT, nil
	case PC_PIPE_BUF:
		return _PIPE_BUF, nil
	case PC_CHOWN_RESTRICTED:
		// XXX: why is _POSIX_CHOWN_RESTRICTED not 1?
		return 1, nil
	case PC_NO_TRUNC:
		return _POSIX_NO_TRUNC, nil
	case PC_VDISABLE:
		return _POSIX_VDISABLE, nil
	case PC_SYNC_IO,
		PC_PRIO_IO,
		PC_REC_INCR_XFER_SIZE,
		PC_REC_MAX_XFER_SIZE,
		PC_SOCK_MAXBUF,
		PC_SYMLINK_MAX:
		return -1, nil
	case PC_FILESIZEBITS:
		return linuxFileSizeBits, nil
	case PC_2_SYMLINKS:
		return linuxSymlinks, nil
	}

	return -1, unix.EINVAL
}

func isReg(mode uint32) bool { return mode&unix.S_IFMT == unix.S_IFREG }
func isBlk(mode uint32) bool { return mode&unix.S_IFMT == unix.S_IFBLK }

func pathconf(path string, name int) (int64, error) {
	if val, err := pathconfConst(name); err == nil {
		return val, nil
	}

	switch name {
	case PC_ASYNC_IO:
		var st unix.Stat_t
		if err := unix.Stat(path, &st); err != nil {
			return -1, err
		}
		if isReg(st.Mode) || isBlk(st.Mode) {
			return 1, nil
		}
		return -1, nil
	}

	statfs := func() (*unix.Statfs_t, error) {
		var st unix.Statfs_t
		if err := unix.Statfs(path, &st); err != nil {
			return nil, err
		}
		return &st, nil
	}

	return pathconfStatfs(name, statfs)
}

func fpathconf(fd int, name int) (int64, error) {
	if val, err := pathconfConst(name); err == nil {
		return val, nil
	}

	switch name {
	case PC_ASYNC_IO:
		var st unix.Stat_t
		if err := unix.Fstat(fd, &st); err != nil {
			return -1, err
		}
		if isReg(st.Mode) || isBlk(st.Mode) {
			return 1, nil
		}
		return -1, nil
	}

	statfs := func() (*unix.Statfs_t, error) {
		var st unix.Statfs_t
		if err := unix.Fstatfs(fd, &st); err != nil {
			return nil, err
		}
		return &st, nil
	}

	return pathconfStatfs(name, statfs)
}

type statfsFn func() (*unix.Statfs_t, error)

func (statfs statfsFn) ValueWithFallback(value valueFn, fallback int64) (int64, error) {
	st, err := statfs()
	if err != nil {
		if errors.Is(err, unix.ENOSYS) {
			return fallback, nil
		}
		return -1, err
	}
	return value(st)
}

func (statfs statfsFn) Value(value valueFn) (int64, error) {
	st, err := statfs()
	if err != nil {
		return -1, err
	}
	return value(st)
}

type valueFn func(*unix.Statfs_t) (int64, error)

func pathconfStatfs(name int, statfs statfsFn) (int64, error) {
	switch name {
	case PC_LINK_MAX:
		return statfs.ValueWithFallback(func(st *unix.Statfs_t) (int64, error) {
			switch uint32(st.Type) {
			case unix.BTRFS_SUPER_MAGIC,
				unix.EXT4_SUPER_MAGIC,
				unix.OVERLAYFS_SUPER_MAGIC,
				unix.XFS_SUPER_MAGIC:
				return linuxLongLinkMax, nil
			}
			return linuxLinkMax, nil
		}, linuxLinkMax)
	case PC_NAME_MAX:
		return statfs.ValueWithFallback(func(st *unix.Statfs_t) (int64, error) {
			return int64(st.Namelen), nil
		}, _NAME_MAX)
	case PC_REC_MIN_XFER_SIZE:
		return statfs.Value(func(st *unix.Statfs_t) (int64, error) {
			if st.Bsize > 0 {
				return int64(st.Bsize), nil
			}
			return linuxRecMinXfer, nil
		})
	case PC_REC_XFER_ALIGN:
		return statfs.Value(func(st *unix.Statfs_t) (int64, error) {
			if st.Frsize > 0 {
				return int64(st.Frsize), nil
			}
			return linuxRecXferAlign, nil
		})
	case PC_ALLOC_SIZE_MIN:
		return statfs.Value(func(st *unix.Statfs_t) (int64, error) {
			if st.Frsize > 0 {
				return int64(st.Frsize), nil
			}
			return linuxAllocSizeMin, nil
		})
	}
	return -1, unix.EINVAL
}
