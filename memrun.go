package main

// modified - from https://github.com/guitmz/memrun

import (
	"fmt"
	"log"
	"os"
	"syscall"
	"unsafe"
)

const (
	mfdCloexec  = 0x0001
	memfdCreate = 319
)

//this will generate the binary data file
//go:generate go-bindata $BINARY

//BinaryName will be set by ldflags
var BinaryName = "unset"

func runFromMemory() {
	fdName := ""
	fd, _, _ := syscall.Syscall(memfdCreate, uintptr(unsafe.Pointer(&fdName)), uintptr(mfdCloexec), 0)

	data, err := Asset(BinaryName)
	if err != nil {
		log.Fatal(err)
	}

	_, _ = syscall.Write(int(fd), data)

	fdPath := fmt.Sprintf("/proc/self/fd/%d", fd)
	_ = syscall.Exec(fdPath, []string{BinaryName}, os.Environ())
}

func main() {
	runFromMemory()
}
