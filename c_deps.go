// +build dummy

// This file is part of a workaround for `go mod vendor` which won't vendor
// C files if there's no Go file in the same directory.
// This would prevent the bundled c-libs from beeing removed when vendored.
//
// This Go file imports the c directory and subdirectories where there is
// another go (named after the direcory) file which is the second part
// of this workaround.
//
// These files combined make it so `go mod vendor` behaves correctly.
// See this issue for reference: https://github.com/golang/go/issues/26366

package hid

import (
	_ "github.com/reyemxela/hid/hidapi"
	_ "github.com/reyemxela/hid/hidapi/hidapi"
	_ "github.com/reyemxela/hid/hidapi/libusb"
	_ "github.com/reyemxela/hid/hidapi/linux"
	_ "github.com/reyemxela/hid/hidapi/mac"
	_ "github.com/reyemxela/hid/hidapi/windows"

	_ "github.com/reyemxela/hid/libusb"
	_ "github.com/reyemxela/hid/libusb/libusb"
	_ "github.com/reyemxela/hid/libusb/libusb/os"
)
