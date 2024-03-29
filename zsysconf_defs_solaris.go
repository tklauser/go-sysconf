// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs sysconf_defs_solaris.go

//go:build solaris

package sysconf

const (
	SC_AIO_LISTIO_MAX               = 0x12
	SC_AIO_MAX                      = 0x13
	SC_AIO_PRIO_DELTA_MAX           = 0x14
	SC_ARG_MAX                      = 0x1
	SC_ATEXIT_MAX                   = 0x4c
	SC_BC_BASE_MAX                  = 0x36
	SC_BC_DIM_MAX                   = 0x37
	SC_BC_SCALE_MAX                 = 0x38
	SC_BC_STRING_MAX                = 0x39
	SC_CHILD_MAX                    = 0x2
	SC_CLK_TCK                      = 0x3
	SC_COLL_WEIGHTS_MAX             = 0x3a
	SC_DELAYTIMER_MAX               = 0x16
	SC_EXPR_NEST_MAX                = 0x3b
	SC_GETGR_R_SIZE_MAX             = 0x239
	SC_GETPW_R_SIZE_MAX             = 0x23a
	SC_HOST_NAME_MAX                = 0x2df
	SC_IOV_MAX                      = 0x4d
	SC_LINE_MAX                     = 0x3c
	SC_LOGIN_NAME_MAX               = 0x23b
	SC_MQ_OPEN_MAX                  = 0x1d
	SC_MQ_PRIO_MAX                  = 0x1e
	SC_NGROUPS_MAX                  = 0x4
	SC_OPEN_MAX                     = 0x5
	SC_PAGE_SIZE                    = 0xb
	SC_PAGESIZE                     = 0xb
	SC_THREAD_DESTRUCTOR_ITERATIONS = 0x238
	SC_THREAD_KEYS_MAX              = 0x23c
	SC_THREAD_STACK_MIN             = 0x23d
	SC_THREAD_THREADS_MAX           = 0x23e
	SC_RE_DUP_MAX                   = 0x3d
	SC_RTSIG_MAX                    = 0x22
	SC_SEM_NSEMS_MAX                = 0x24
	SC_SEM_VALUE_MAX                = 0x25
	SC_SIGQUEUE_MAX                 = 0x27
	SC_STREAM_MAX                   = 0x10
	SC_SYMLOOP_MAX                  = 0x2e8
	SC_TIMER_MAX                    = 0x2c
	SC_TTY_NAME_MAX                 = 0x23f
	SC_TZNAME_MAX                   = 0x11

	SC_ADVISORY_INFO              = 0x2db
	SC_ASYNCHRONOUS_IO            = 0x15
	SC_BARRIERS                   = 0x2dc
	SC_CLOCK_SELECTION            = 0x2dd
	SC_CPUTIME                    = 0x2de
	SC_FSYNC                      = 0x17
	SC_IPV6                       = 0x2fa
	SC_JOB_CONTROL                = 0x6
	SC_MAPPED_FILES               = 0x18
	SC_MEMLOCK                    = 0x19
	SC_MEMLOCK_RANGE              = 0x1a
	SC_MEMORY_PROTECTION          = 0x1b
	SC_MESSAGE_PASSING            = 0x1c
	SC_MONOTONIC_CLOCK            = 0x2e0
	SC_PRIORITIZED_IO             = 0x1f
	SC_PRIORITY_SCHEDULING        = 0x20
	SC_RAW_SOCKETS                = 0x2fb
	SC_READER_WRITER_LOCKS        = 0x2e1
	SC_REALTIME_SIGNALS           = 0x21
	SC_REGEXP                     = 0x2e2
	SC_SAVED_IDS                  = 0x7
	SC_SEMAPHORES                 = 0x23
	SC_SHARED_MEMORY_OBJECTS      = 0x26
	SC_SHELL                      = 0x2e3
	SC_SPAWN                      = 0x2e4
	SC_SPIN_LOCKS                 = 0x2e5
	SC_SPORADIC_SERVER            = 0x2e6
	SC_SS_REPL_MAX                = 0x2e7
	SC_SYNCHRONIZED_IO            = 0x2a
	SC_THREAD_ATTR_STACKADDR      = 0x241
	SC_THREAD_ATTR_STACKSIZE      = 0x242
	SC_THREAD_CPUTIME             = 0x2e9
	SC_THREAD_PRIO_INHERIT        = 0x244
	SC_THREAD_PRIO_PROTECT        = 0x245
	SC_THREAD_PRIORITY_SCHEDULING = 0x243
	SC_THREAD_PROCESS_SHARED      = 0x246
	SC_THREAD_SAFE_FUNCTIONS      = 0x247
	SC_THREAD_SPORADIC_SERVER     = 0x2ea
	SC_THREADS                    = 0x240
	SC_TIMEOUTS                   = 0x2eb
	SC_TIMERS                     = 0x2b
	SC_TRACE                      = 0x2ec
	SC_TRACE_EVENT_FILTER         = 0x2ed
	SC_TRACE_EVENT_NAME_MAX       = 0x2ee
	SC_TRACE_INHERIT              = 0x2ef
	SC_TRACE_LOG                  = 0x2f0
	SC_TRACE_NAME_MAX             = 0x2f1
	SC_TRACE_SYS_MAX              = 0x2f2
	SC_TRACE_USER_EVENT_MAX       = 0x2f3
	SC_TYPED_MEMORY_OBJECTS       = 0x2f4
	SC_VERSION                    = 0x8

	SC_V6_ILP32_OFF32  = 0x2f5
	SC_V6_ILP32_OFFBIG = 0x2f6
	SC_V6_LP64_OFF64   = 0x2f7
	SC_V6_LPBIG_OFFBIG = 0x2f8

	SC_2_C_BIND         = 0x2d
	SC_2_C_DEV          = 0x2e
	SC_2_C_VERSION      = 0x2f
	SC_2_CHAR_TERM      = 0x42
	SC_2_FORT_DEV       = 0x30
	SC_2_FORT_RUN       = 0x31
	SC_2_LOCALEDEF      = 0x32
	SC_2_PBS            = 0x2d4
	SC_2_PBS_ACCOUNTING = 0x2d5
	SC_2_PBS_CHECKPOINT = 0x2d6
	SC_2_PBS_LOCATE     = 0x2d8
	SC_2_PBS_MESSAGE    = 0x2d9
	SC_2_PBS_TRACK      = 0x2da
	SC_2_SW_DEV         = 0x33
	SC_2_UPE            = 0x34
	SC_2_VERSION        = 0x35

	SC_XOPEN_CRYPT            = 0x3e
	SC_XOPEN_ENH_I18N         = 0x3f
	SC_XOPEN_REALTIME         = 0x2ce
	SC_XOPEN_REALTIME_THREADS = 0x2cf
	SC_XOPEN_SHM              = 0x40
	SC_XOPEN_STREAMS          = 0x2f9
	SC_XOPEN_UNIX             = 0x4e
	SC_XOPEN_VERSION          = 0xc
	SC_XOPEN_XCU_VERSION      = 0x43

	SC_PHYS_PAGES       = 0x1f4
	SC_AVPHYS_PAGES     = 0x1f5
	SC_NPROCESSORS_CONF = 0xe
	SC_NPROCESSORS_ONLN = 0xf
)
