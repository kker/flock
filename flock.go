package flock

import (
	"os"
	"syscall"
)

// TryLock 获取排他锁
func TryLock(filePath string) (succ bool, err error) {
	fh, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR, 0600)
	if err != nil {
		return false, err
	}
	flock := syscall.Flock_t{
		Type: syscall.F_WRLCK,
		Pid:  int32(os.Getpid()),
	}

	if err := syscall.FcntlFlock(fh.Fd(), syscall.F_SETLK, &flock); err != nil {
		return false, nil
	}
	return true, nil
}
