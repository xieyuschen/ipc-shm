package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
	"golang.org/x/sys/unix"
)

func main() {
	msg := pkg.Message{
		Field1: 1,
	}
	copy(msg.Field2[:], "producer")

	addr, err := createSharedMemory()
	beginningAddr := addr
	defer func() {
		_, _, err = syscall.Syscall(syscall.SYS_SHMDT, beginningAddr, 0, 0)
		fmt.Println("++", err)
	}()
	if err != nil {
		panic(err)
	}

	for i := 0; ; i++ {
		fmt.Println("write a new structure:", i)
		msg.Field1 = int32(i)
		writeToSharedMemory(addr, msg)
		addr += pkg.Size
		time.Sleep(time.Second)
	}
}

func createSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET,
		pkg.ShmId,
		pkg.Size,
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

func writeToSharedMemory(addr uintptr, s pkg.Message) {
	ptr := (*pkg.Message)(unsafe.Pointer(addr))
	*ptr = s
}
