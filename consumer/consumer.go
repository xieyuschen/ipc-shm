package main

import (
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
)

// notes
// 1. 'copy(myStruct.Field2[:], "consumer")' cannot change the value of structure,
// 		which is retrieved from the shm

func main() {
	addr, err := accessSharedMemory()
	if err != nil {
		panic(err)
	}

	for {
		v := readFromSharedMemory(addr)
		addr += pkg.Size
		log.Printf("Field1: %d, Field2: %s\n", v.Field1, v.Field2)
		time.Sleep(time.Second)
	}
}

func accessSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET,
		pkg.ShmId, pkg.Size, 0666)
	if err != 0 {
		return 0, err
	}
	addr, _, err := syscall.Syscall(syscall.SYS_SHMAT, id, 0, 0)
	if err != 0 {
		return 0, err
	}
	return addr, nil
}

func readFromSharedMemory(addr uintptr) pkg.Message {
	ptr := (*pkg.Message)(unsafe.Pointer(addr))
	return *ptr
}
