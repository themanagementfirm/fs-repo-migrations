// +build !plan9,!windows,!wasm

package reuseport

import (
	"syscall"

	"github.com/ipfs/fs-repo-migrations/ipfs-10-to-11/_vendor/golang.org/x/sys/unix"
)

func Control(network, address string, c syscall.RawConn) error {
	var err error
	c.Control(func(fd uintptr) {
		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEADDR, 1)
		if err != nil {
			return
		}

		err = unix.SetsockoptInt(int(fd), unix.SOL_SOCKET, unix.SO_REUSEPORT, 1)
		if err != nil {
			return
		}
	})
	return err
}