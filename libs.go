// usb - Self contained USB and HID library for Go
// Copyright 2019 The library Authors
//
// This library is free software: you can redistribute it and/or modify it under
// the terms of the GNU Lesser General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option) any
// later version.
//
// The library is distributed in the hope that it will be useful, but WITHOUT ANY
// WARRANTY; without even the implied warranty of MERCHANTABILITY or FITNESS FOR
// A PARTICULAR PURPOSE. See the GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License along
// with the library. If not, see <http://www.gnu.org/licenses/>.

//go:build (freebsd && cgo) || (linux && cgo) || (darwin && !ios && cgo) || (windows && cgo)
// +build freebsd,cgo linux,cgo darwin,!ios,cgo windows,cgo

package usb

/*
#cgo CFLAGS: -I./hidapi/hidapi
#cgo CFLAGS: -I /usr/include/libusb-1.0
#cgo CFLAGS: -DDEFAULT_VISIBILITY=""
#cgo CFLAGS: -DPOLL_NFDS_TYPE=int

#cgo linux CFLAGS: -DOS_LINUX -D_GNU_SOURCE -DHAVE_SYS_TIME_H -I./hidapi/hidapi
#cgo linux,!android LDFLAGS: -lrt -lusb-1.0
#cgo darwin CFLAGS: -DOS_DARWIN -DHAVE_SYS_TIME_H
#cgo darwin LDFLAGS: -framework CoreFoundation -framework IOKit -lobjc
#cgo windows CFLAGS: -DOS_WINDOWS
#cgo windows LDFLAGS: -lsetupapi
#cgo freebsd CFLAGS: -DOS_FREEBSD
#cgo freebsd LDFLAGS: -lusb
#cgo openbsd CFLAGS: -DOS_OPENBSD

#ifdef OS_LINUX
	#include "hidapi/libusb/hid.c"
#elif OS_DARWIN
	#include "hidapi/mac/hid.c"
#elif OS_WINDOWS
	#include "hidapi/windows/hid.c"
#elif OS_FREEBSD
	#include <libusb.h>
	#include "hidapi/libusb/hid.c"
#elif DOS_OPENBSD
	#include "hidapi/libusb/hid.c"
#endif
*/
import "C"
