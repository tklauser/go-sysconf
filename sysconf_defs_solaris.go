// Copyright 2021 Tobias Klauser. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:build ignore

package sysconf

/*
#include <limits.h>
#include <netdb.h>
#include <stdio.h>
#include <unistd.h>
*/
import "C"

// sysconf variables
const (
	SC_AIO_LISTIO_MAX               = C._SC_AIO_LISTIO_MAX
	SC_AIO_MAX                      = C._SC_AIO_MAX
	SC_AIO_PRIO_DELTA_MAX           = C._SC_AIO_PRIO_DELTA_MAX
	SC_ARG_MAX                      = C._SC_ARG_MAX
	SC_ATEXIT_MAX                   = C._SC_ATEXIT_MAX
	SC_BC_BASE_MAX                  = C._SC_BC_BASE_MAX
	SC_BC_DIM_MAX                   = C._SC_BC_DIM_MAX
	SC_BC_SCALE_MAX                 = C._SC_BC_SCALE_MAX
	SC_BC_STRING_MAX                = C._SC_BC_STRING_MAX
	SC_CHILD_MAX                    = C._SC_CHILD_MAX
	SC_CLK_TCK                      = C._SC_CLK_TCK
	SC_COLL_WEIGHTS_MAX             = C._SC_COLL_WEIGHTS_MAX
	SC_DELAYTIMER_MAX               = C._SC_DELAYTIMER_MAX
	SC_EXPR_NEST_MAX                = C._SC_EXPR_NEST_MAX
	SC_GETGR_R_SIZE_MAX             = C._SC_GETGR_R_SIZE_MAX
	SC_GETPW_R_SIZE_MAX             = C._SC_GETPW_R_SIZE_MAX
	SC_HOST_NAME_MAX                = C._SC_HOST_NAME_MAX
	SC_IOV_MAX                      = C._SC_IOV_MAX
	SC_LINE_MAX                     = C._SC_LINE_MAX
	SC_LOGIN_NAME_MAX               = C._SC_LOGIN_NAME_MAX
	SC_MQ_OPEN_MAX                  = C._SC_MQ_OPEN_MAX
	SC_MQ_PRIO_MAX                  = C._SC_MQ_PRIO_MAX
	SC_NGROUPS_MAX                  = C._SC_NGROUPS_MAX
	SC_OPEN_MAX                     = C._SC_OPEN_MAX
	SC_PAGE_SIZE                    = C._SC_PAGE_SIZE
	SC_PAGESIZE                     = C._SC_PAGESIZE
	SC_THREAD_DESTRUCTOR_ITERATIONS = C._SC_THREAD_DESTRUCTOR_ITERATIONS
	SC_THREAD_KEYS_MAX              = C._SC_THREAD_KEYS_MAX
	SC_THREAD_STACK_MIN             = C._SC_THREAD_STACK_MIN
	SC_THREAD_THREADS_MAX           = C._SC_THREAD_THREADS_MAX
	SC_RE_DUP_MAX                   = C._SC_RE_DUP_MAX
	SC_RTSIG_MAX                    = C._SC_RTSIG_MAX
	SC_SEM_NSEMS_MAX                = C._SC_SEM_NSEMS_MAX
	SC_SEM_VALUE_MAX                = C._SC_SEM_VALUE_MAX
	SC_SIGQUEUE_MAX                 = C._SC_SIGQUEUE_MAX
	SC_STREAM_MAX                   = C._SC_STREAM_MAX
	SC_SYMLOOP_MAX                  = C._SC_SYMLOOP_MAX
	SC_TIMER_MAX                    = C._SC_TIMER_MAX
	SC_TTY_NAME_MAX                 = C._SC_TTY_NAME_MAX
	SC_TZNAME_MAX                   = C._SC_TZNAME_MAX

	SC_ADVISORY_INFO              = C._SC_ADVISORY_INFO
	SC_ASYNCHRONOUS_IO            = C._SC_ASYNCHRONOUS_IO
	SC_BARRIERS                   = C._SC_BARRIERS
	SC_CLOCK_SELECTION            = C._SC_CLOCK_SELECTION
	SC_CPUTIME                    = C._SC_CPUTIME
	SC_FSYNC                      = C._SC_FSYNC
	SC_IPV6                       = C._SC_IPV6
	SC_JOB_CONTROL                = C._SC_JOB_CONTROL
	SC_MAPPED_FILES               = C._SC_MAPPED_FILES
	SC_MEMLOCK                    = C._SC_MEMLOCK
	SC_MEMLOCK_RANGE              = C._SC_MEMLOCK_RANGE
	SC_MEMORY_PROTECTION          = C._SC_MEMORY_PROTECTION
	SC_MESSAGE_PASSING            = C._SC_MESSAGE_PASSING
	SC_MONOTONIC_CLOCK            = C._SC_MONOTONIC_CLOCK
	SC_PRIORITIZED_IO             = C._SC_PRIORITIZED_IO
	SC_PRIORITY_SCHEDULING        = C._SC_PRIORITY_SCHEDULING
	SC_RAW_SOCKETS                = C._SC_RAW_SOCKETS
	SC_READER_WRITER_LOCKS        = C._SC_READER_WRITER_LOCKS
	SC_REALTIME_SIGNALS           = C._SC_REALTIME_SIGNALS
	SC_REGEXP                     = C._SC_REGEXP
	SC_SAVED_IDS                  = C._SC_SAVED_IDS
	SC_SEMAPHORES                 = C._SC_SEMAPHORES
	SC_SHARED_MEMORY_OBJECTS      = C._SC_SHARED_MEMORY_OBJECTS
	SC_SHELL                      = C._SC_SHELL
	SC_SPAWN                      = C._SC_SPAWN
	SC_SPIN_LOCKS                 = C._SC_SPIN_LOCKS
	SC_SPORADIC_SERVER            = C._SC_SPORADIC_SERVER
	SC_SS_REPL_MAX                = C._SC_SS_REPL_MAX
	SC_SYNCHRONIZED_IO            = C._SC_SYNCHRONIZED_IO
	SC_THREAD_ATTR_STACKADDR      = C._SC_THREAD_ATTR_STACKADDR
	SC_THREAD_ATTR_STACKSIZE      = C._SC_THREAD_ATTR_STACKSIZE
	SC_THREAD_CPUTIME             = C._SC_THREAD_CPUTIME
	SC_THREAD_PRIO_INHERIT        = C._SC_THREAD_PRIO_INHERIT
	SC_THREAD_PRIO_PROTECT        = C._SC_THREAD_PRIO_PROTECT
	SC_THREAD_PRIORITY_SCHEDULING = C._SC_THREAD_PRIORITY_SCHEDULING
	SC_THREAD_PROCESS_SHARED      = C._SC_THREAD_PROCESS_SHARED
	SC_THREAD_SAFE_FUNCTIONS      = C._SC_THREAD_SAFE_FUNCTIONS
	SC_THREAD_SPORADIC_SERVER     = C._SC_THREAD_SPORADIC_SERVER
	SC_THREADS                    = C._SC_THREADS
	SC_TIMEOUTS                   = C._SC_TIMEOUTS
	SC_TIMERS                     = C._SC_TIMERS
	SC_TRACE                      = C._SC_TRACE
	SC_TRACE_EVENT_FILTER         = C._SC_TRACE_EVENT_FILTER
	SC_TRACE_EVENT_NAME_MAX       = C._SC_TRACE_EVENT_NAME_MAX
	SC_TRACE_INHERIT              = C._SC_TRACE_INHERIT
	SC_TRACE_LOG                  = C._SC_TRACE_LOG
	SC_TRACE_NAME_MAX             = C._SC_TRACE_NAME_MAX
	SC_TRACE_SYS_MAX              = C._SC_TRACE_SYS_MAX
	SC_TRACE_USER_EVENT_MAX       = C._SC_TRACE_USER_EVENT_MAX
	SC_TYPED_MEMORY_OBJECTS       = C._SC_TYPED_MEMORY_OBJECTS
	SC_VERSION                    = C._SC_VERSION

	SC_V6_ILP32_OFF32  = C._SC_V6_ILP32_OFF32
	SC_V6_ILP32_OFFBIG = C._SC_V6_ILP32_OFFBIG
	SC_V6_LP64_OFF64   = C._SC_V6_LP64_OFF64
	SC_V6_LPBIG_OFFBIG = C._SC_V6_LPBIG_OFFBIG

	SC_2_C_BIND         = C._SC_2_C_BIND
	SC_2_C_DEV          = C._SC_2_C_DEV
	SC_2_C_VERSION      = C._SC_2_C_VERSION
	SC_2_CHAR_TERM      = C._SC_2_CHAR_TERM
	SC_2_FORT_DEV       = C._SC_2_FORT_DEV
	SC_2_FORT_RUN       = C._SC_2_FORT_RUN
	SC_2_LOCALEDEF      = C._SC_2_LOCALEDEF
	SC_2_PBS            = C._SC_2_PBS
	SC_2_PBS_ACCOUNTING = C._SC_2_PBS_ACCOUNTING
	SC_2_PBS_CHECKPOINT = C._SC_2_PBS_CHECKPOINT
	SC_2_PBS_LOCATE     = C._SC_2_PBS_LOCATE
	SC_2_PBS_MESSAGE    = C._SC_2_PBS_MESSAGE
	SC_2_PBS_TRACK      = C._SC_2_PBS_TRACK
	SC_2_SW_DEV         = C._SC_2_SW_DEV
	SC_2_UPE            = C._SC_2_UPE
	SC_2_VERSION        = C._SC_2_VERSION

	SC_XOPEN_CRYPT            = C._SC_XOPEN_CRYPT
	SC_XOPEN_ENH_I18N         = C._SC_XOPEN_ENH_I18N
	SC_XOPEN_REALTIME         = C._SC_XOPEN_REALTIME
	SC_XOPEN_REALTIME_THREADS = C._SC_XOPEN_REALTIME_THREADS
	SC_XOPEN_SHM              = C._SC_XOPEN_SHM
	SC_XOPEN_STREAMS          = C._SC_XOPEN_STREAMS
	SC_XOPEN_UNIX             = C._SC_XOPEN_UNIX
	SC_XOPEN_VERSION          = C._SC_XOPEN_VERSION
	SC_XOPEN_XCU_VERSION      = C._SC_XOPEN_XCU_VERSION

	// non-standard variables
	SC_PHYS_PAGES       = C._SC_PHYS_PAGES
	SC_AVPHYS_PAGES     = C._SC_AVPHYS_PAGES
	SC_NPROCESSORS_CONF = C._SC_NPROCESSORS_CONF
	SC_NPROCESSORS_ONLN = C._SC_NPROCESSORS_ONLN
)
