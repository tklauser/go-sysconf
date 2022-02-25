// Copyright 2018 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore
// +build ignore

package sysconf

/*
#include <limits.h>
#include <paths.h>
#include <pwd.h>
#include <stdio.h>
#include <sys/param.h>
#include <time.h>
#include <unistd.h>

#ifndef _PATH_ZONEINFO
# define _PATH_ZONEINFO	"/usr/share/zoneinfo" // TZDATA from tzcode/stdtime/tzfile.h
#endif
*/
import "C"

// sysconf variables
const (
	SC_ARG_MAX                      = C._SC_ARG_MAX
	SC_CHILD_MAX                    = C._SC_CHILD_MAX
	SC_NGROUPS_MAX                  = C._SC_NGROUPS_MAX
	SC_OPEN_MAX                     = C._SC_OPEN_MAX
	SC_JOB_CONTROL                  = C._SC_JOB_CONTROL
	SC_SAVED_IDS                    = C._SC_SAVED_IDS
	SC_VERSION                      = C._SC_VERSION
	SC_BC_BASE_MAX                  = C._SC_BC_BASE_MAX
	SC_BC_DIM_MAX                   = C._SC_BC_DIM_MAX
	SC_BC_SCALE_MAX                 = C._SC_BC_SCALE_MAX
	SC_BC_STRING_MAX                = C._SC_BC_STRING_MAX
	SC_COLL_WEIGHTS_MAX             = C._SC_COLL_WEIGHTS_MAX
	SC_EXPR_NEST_MAX                = C._SC_EXPR_NEST_MAX
	SC_LINE_MAX                     = C._SC_LINE_MAX
	SC_RE_DUP_MAX                   = C._SC_RE_DUP_MAX
	SC_2_VERSION                    = C._SC_2_VERSION
	SC_2_C_BIND                     = C._SC_2_C_BIND
	SC_2_C_DEV                      = C._SC_2_C_DEV
	SC_2_CHAR_TERM                  = C._SC_2_CHAR_TERM
	SC_2_FORT_DEV                   = C._SC_2_FORT_DEV
	SC_2_FORT_RUN                   = C._SC_2_FORT_RUN
	SC_2_LOCALEDEF                  = C._SC_2_LOCALEDEF
	SC_2_SW_DEV                     = C._SC_2_SW_DEV
	SC_2_UPE                        = C._SC_2_UPE
	SC_STREAM_MAX                   = C._SC_STREAM_MAX
	SC_TZNAME_MAX                   = C._SC_TZNAME_MAX
	SC_PAGESIZE                     = C._SC_PAGESIZE
	SC_PAGE_SIZE                    = C._SC_PAGE_SIZE
	SC_FSYNC                        = C._SC_FSYNC
	SC_XOPEN_SHM                    = C._SC_XOPEN_SHM
	SC_SYNCHRONIZED_IO              = C._SC_SYNCHRONIZED_IO
	SC_IOV_MAX                      = C._SC_IOV_MAX
	SC_MAPPED_FILES                 = C._SC_MAPPED_FILES
	SC_MEMLOCK                      = C._SC_MEMLOCK
	SC_MEMLOCK_RANGE                = C._SC_MEMLOCK_RANGE
	SC_MEMORY_PROTECTION            = C._SC_MEMORY_PROTECTION
	SC_LOGIN_NAME_MAX               = C._SC_LOGIN_NAME_MAX
	SC_MONOTONIC_CLOCK              = C._SC_MONOTONIC_CLOCK
	SC_CLK_TCK                      = C._SC_CLK_TCK
	SC_ATEXIT_MAX                   = C._SC_ATEXIT_MAX
	SC_THREADS                      = C._SC_THREADS
	SC_SEMAPHORES                   = C._SC_SEMAPHORES
	SC_BARRIERS                     = C._SC_BARRIERS
	SC_TIMERS                       = C._SC_TIMERS
	SC_SPIN_LOCKS                   = C._SC_SPIN_LOCKS
	SC_READER_WRITER_LOCKS          = C._SC_READER_WRITER_LOCKS
	SC_GETGR_R_SIZE_MAX             = C._SC_GETGR_R_SIZE_MAX
	SC_GETPW_R_SIZE_MAX             = C._SC_GETPW_R_SIZE_MAX
	SC_CLOCK_SELECTION              = C._SC_CLOCK_SELECTION
	SC_ASYNCHRONOUS_IO              = C._SC_ASYNCHRONOUS_IO
	SC_AIO_LISTIO_MAX               = C._SC_AIO_LISTIO_MAX
	SC_AIO_MAX                      = C._SC_AIO_MAX
	SC_MESSAGE_PASSING              = C._SC_MESSAGE_PASSING
	SC_MQ_OPEN_MAX                  = C._SC_MQ_OPEN_MAX
	SC_MQ_PRIO_MAX                  = C._SC_MQ_PRIO_MAX
	SC_PRIORITY_SCHEDULING          = C._SC_PRIORITY_SCHEDULING
	SC_THREAD_DESTRUCTOR_ITERATIONS = C._SC_THREAD_DESTRUCTOR_ITERATIONS
	SC_THREAD_KEYS_MAX              = C._SC_THREAD_KEYS_MAX
	SC_THREAD_STACK_MIN             = C._SC_THREAD_STACK_MIN
	SC_THREAD_THREADS_MAX           = C._SC_THREAD_THREADS_MAX
	SC_THREAD_ATTR_STACKADDR        = C._SC_THREAD_ATTR_STACKADDR
	SC_THREAD_ATTR_STACKSIZE        = C._SC_THREAD_ATTR_STACKSIZE
	SC_THREAD_PRIORITY_SCHEDULING   = C._SC_THREAD_PRIORITY_SCHEDULING
	SC_THREAD_PRIO_INHERIT          = C._SC_THREAD_PRIO_INHERIT
	SC_THREAD_PRIO_PROTECT          = C._SC_THREAD_PRIO_PROTECT
	SC_THREAD_PROCESS_SHARED        = C._SC_THREAD_PROCESS_SHARED
	SC_THREAD_SAFE_FUNCTIONS        = C._SC_THREAD_SAFE_FUNCTIONS
	SC_TTY_NAME_MAX                 = C._SC_TTY_NAME_MAX
	SC_HOST_NAME_MAX                = C._SC_HOST_NAME_MAX
	SC_PASS_MAX                     = C._SC_PASS_MAX
	SC_REGEXP                       = C._SC_REGEXP
	SC_SHELL                        = C._SC_SHELL
	SC_SYMLOOP_MAX                  = C._SC_SYMLOOP_MAX

	// not supported or implemented yet
	SC_V6_ILP32_OFF32   = C._SC_V6_ILP32_OFF32
	SC_V6_ILP32_OFFBIG  = C._SC_V6_ILP32_OFFBIG
	SC_V6_LP64_OFF64    = C._SC_V6_LP64_OFF64
	SC_V6_LPBIG_OFFBIG  = C._SC_V6_LPBIG_OFFBIG
	SC_2_PBS            = C._SC_2_PBS
	SC_2_PBS_ACCOUNTING = C._SC_2_PBS_ACCOUNTING
	SC_2_PBS_CHECKPOINT = C._SC_2_PBS_CHECKPOINT
	SC_2_PBS_LOCATE     = C._SC_2_PBS_LOCATE
	SC_2_PBS_MESSAGE    = C._SC_2_PBS_MESSAGE
	SC_2_PBS_TRACK      = C._SC_2_PBS_TRACK

	SC_SPAWN                 = C._SC_SPAWN
	SC_SHARED_MEMORY_OBJECTS = C._SC_SHARED_MEMORY_OBJECTS

	SC_TIMER_MAX        = C._SC_TIMER_MAX
	SC_SEM_NSEMS_MAX    = C._SC_SEM_NSEMS_MAX
	SC_CPUTIME          = C._SC_CPUTIME
	SC_THREAD_CPUTIME   = C._SC_THREAD_CPUTIME
	SC_DELAYTIMER_MAX   = C._SC_DELAYTIMER_MAX
	SC_SIGQUEUE_MAX     = C._SC_SIGQUEUE_MAX
	SC_REALTIME_SIGNALS = C._SC_REALTIME_SIGNALS

	// extensions found in Solaris and Linux
	SC_PHYS_PAGES = C._SC_PHYS_PAGES

	// commonly provided extensiosn
	SC_NPROCESSORS_CONF = C._SC_NPROCESSORS_CONF
	SC_NPROCESSORS_ONLN = C._SC_NPROCESSORS_ONLN

	// native variables
	SC_SCHED_RT_TS   = C._SC_SCHED_RT_TS
	SC_SCHED_PRI_MIN = C._SC_SCHED_PRI_MIN
	SC_SCHED_PRI_MAX = C._SC_SCHED_PRI_MAX
)

// sysconf values
const (
	_MAXHOSTNAMELEN = C.MAXHOSTNAMELEN
	_MAXLOGNAME     = C.MAXLOGNAME
	_MAXSYMLINKS    = C.MAXSYMLINKS

	_POSIX_ARG_MAX                      = C._POSIX_ARG_MAX
	_POSIX_CHILD_MAX                    = C._POSIX_CHILD_MAX
	_POSIX_CPUTIME                      = C._POSIX_CPUTIME
	_POSIX_DELAYTIMER_MAX               = C._POSIX_DELAYTIMER_MAX
	_POSIX_PRIORITY_SCHEDULING          = C._POSIX_PRIORITY_SCHEDULING
	_POSIX_REGEXP                       = C._POSIX_REGEXP
	_POSIX_SHARED_MEMORY_OBJECTS        = C._POSIX_SHARED_MEMORY_OBJECTS
	_POSIX_SHELL                        = C._POSIX_SHELL
	_POSIX_SIGQUEUE_MAX                 = C._POSIX_SIGQUEUE_MAX
	_POSIX_SPAWN                        = C._POSIX_SPAWN
	_POSIX_THREAD_ATTR_STACKADDR        = C._POSIX_THREAD_ATTR_STACKADDR
	_POSIX_THREAD_ATTR_STACKSIZE        = C._POSIX_THREAD_ATTR_STACKSIZE
	_POSIX_THREAD_CPUTIME               = C._POSIX_THREAD_CPUTIME
	_POSIX_THREAD_DESTRUCTOR_ITERATIONS = C._POSIX_THREAD_DESTRUCTOR_ITERATIONS
	_POSIX_THREAD_KEYS_MAX              = C._POSIX_THREAD_KEYS_MAX
	_POSIX_THREAD_PRIO_PROTECT          = C._POSIX_THREAD_PRIO_PROTECT
	_POSIX_THREAD_SAFE_FUNCTIONS        = C._POSIX_THREAD_SAFE_FUNCTIONS
	_POSIX_TIMER_MAX                    = C._POSIX_TIMER_MAX
	_POSIX_VERSION                      = C._POSIX_VERSION

	_POSIX2_VERSION = C._POSIX2_VERSION

	_FOPEN_MAX  = C.FOPEN_MAX
	_NAME_MAX   = C.NAME_MAX
	_RE_DUP_MAX = C.RE_DUP_MAX

	_BC_BASE_MAX      = C.BC_BASE_MAX
	_BC_DIM_MAX       = C.BC_DIM_MAX
	_BC_SCALE_MAX     = C.BC_SCALE_MAX
	_BC_STRING_MAX    = C.BC_STRING_MAX
	_COLL_WEIGHTS_MAX = C.COLL_WEIGHTS_MAX
	_EXPR_NEST_MAX    = C.EXPR_NEST_MAX
	_LINE_MAX         = C.LINE_MAX

	_GETGR_R_SIZE_MAX = C._GETGR_R_SIZE_MAX
	_GETPW_R_SIZE_MAX = C._GETPW_R_SIZE_MAX

	_PATH_DEV      = C._PATH_DEV
	_PATH_ZONEINFO = C._PATH_ZONEINFO

	_PASSWORD_LEN = C._PASSWORD_LEN
)

// pathconf variables

const _PC_NAME_MAX = C._PC_NAME_MAX
