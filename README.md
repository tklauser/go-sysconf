# go-sysconf

`sysconf(3)` for Go, without using cgo.

Supported operating systems: Linux, FreeBSD.

Support for NetBSD, OpenBSD, DragonflyBSD, Darwin and Solaris is planned but not yet implemented.

All POSIX.1 and POSIX.2 variables are supported, see
http://man7.org/linux/man-pages/man3/sysconf.3.html

Additionally, the following non-standard variables are supported on some operating systems:

| Variable | Supported on |
|---|---|
| `SC_PHYS_PAGES`       | Linux, FreeBSD |
| `SC_AVPHYS_PAGES`     | Linux          |
| `SC_NPROCESSORS_CONF` | Linux, FreeBSD |
| `SC_NPROCESSORS_ONLN` | Linux, FreeBSD |

## Usage

TODO

## Documentation

https://godoc.org/github.com/tklauser/go-sysconf
