# go-sysconf

[![Build Status][1]][2] [![Go Report Card][3]][4]

`sysconf` for Go, without using cgo.

Supported operating systems: Linux, Darwin, FreeBSD, NetBSD.
Support for OpenBSD, DragonflyBSD and Solaris is planned but not yet implemented.

All POSIX.1 and POSIX.2 variables are supported, see [References](#references) for a complete list.

Additionally, the following non-standard variables are supported on some operating systems:

| Variable | Supported on |
|---|---|
| `SC_PHYS_PAGES`       | Linux, Darwin, FreeBSD, NetBSD |
| `SC_AVPHYS_PAGES`     | Linux |
| `SC_NPROCESSORS_CONF` | Linux, Darwin, FreeBSD, NetBSD |
| `SC_NPROCESSORS_ONLN` | Linux, Darwin, FreeBSD, NetBSD |
| `SC_UIO_MAXIOV`       | Linux |

## Usage

TODO

## Documentation

https://godoc.org/github.com/tklauser/go-sysconf

## References

* [POSIX documenation for`sysconf`](http://pubs.opengroup.org/onlinepubs/9699919799/functions/sysconf.html)
* [Linux manpage for `sysconf(3)`](http://man7.org/linux/man-pages/man3/sysconf.3.html)

[1]: https://travis-ci.org/tklauser/go-sysconf.svg?branch=master
[2]: https://travis-ci.org/tklauser/go-sysconf
[3]: https://goreportcard.com/badge/github.com/tklauser/go-sysconf
[4]: https://goreportcard.com/report/github.com/tklauser/go-sysconf
