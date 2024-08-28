package pkg

import (
	"golang.org/x/sys/unix"
	"syscall"
)

func AccessSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET,
		ShmId, Size, 0666)
	if err != 0 {
		return 0, err
	}
	addr, _, err := syscall.Syscall(syscall.SYS_SHMAT, id, 0, 0)
	if err != 0 {
		return 0, err
	}
	return addr, nil
}

func CreateSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET,
		ShmId,
		Size,
		unix.IPC_CREAT|0666)
	if err != 0 {
		return 0, err
	}
	addr, _, err := syscall.Syscall(syscall.SYS_SHMAT, id, 0, 0)
	if err != 0 {
		return 0, err
	}
	return addr, nil
}
