# go-sysconf

`sysconf` for Go, without using cgo.

Supported operating systems: Linux, FreeBSD, NetBSD.
Support for OpenBSD, DragonflyBSD, Darwin and Solaris is planned but not yet implemented.

All POSIX.1 and POSIX.2 variables are supported, see [References](#references) for a complete list.

Additionally, the following non-standard variables are supported on some operating systems:

| Variable | Supported on |
|---|---|
| `SC_PHYS_PAGES`       | Linux, FreeBSD, NetBSD |
| `SC_AVPHYS_PAGES`     | Linux |
| `SC_NPROCESSORS_CONF` | Linux, FreeBSD, NetBSD |
| `SC_NPROCESSORS_ONLN` | Linux, FreeBSD, NetBSD |

## Usage

TODO

## Documentation

https://godoc.org/github.com/tklauser/go-sysconf

## References

* [POSIX documenation for`sysconf`](http://pubs.opengroup.org/onlinepubs/9699919799/functions/sysconf.html)
* [Linux manpage for `sysconf(3)`](http://man7.org/linux/man-pages/man3/sysconf.3.html)
