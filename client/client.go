package main

import (
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
)

// Replace with your key
const size = int(unsafe.Sizeof(pkg.MyStruct{})) // Size of your structure

func accessSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), 0666)
	if err != 0 {
		return 0, err
	}
	addr, _, err := syscall.Syscall(syscall.SYS_SHMAT, id, 0, 0)
	if err != 0 {
		return 0, err
	}
	return addr, nil
}

func readFromSharedMemory(addr uintptr) pkg.MyStruct {
	ptr := (*pkg.MyStruct)(unsafe.Pointer(addr))
	return *ptr
}

const key = 1234

func main() {

	for {
		addr, err := accessSharedMemory()
		if err != nil {
			log.Fatal(err)
		}

		myStruct := readFromSharedMemory(addr)
		log.Printf("Field1: %d, Field2: %s", myStruct.Field1, myStruct.Field2)

		time.Sleep(time.Second)
	}

}
