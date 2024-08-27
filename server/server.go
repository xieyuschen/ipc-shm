package main

import (
	"log"
	"syscall"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
	"golang.org/x/sys/unix"
)

const key = 1234                                // Replace with your key
const size = int(unsafe.Sizeof(pkg.MyStruct{})) // Size of your structure

func createSharedMemory() (uintptr, error) {
	id, _, err := syscall.Syscall(syscall.SYS_SHMGET, uintptr(key), uintptr(size), unix.IPC_CREAT|0666)
	if err != 0 {
		return 0, err
	}
	addr, _, err := syscall.Syscall(syscall.SYS_SHMAT, id, 0, 0)
	if err != 0 {
		return 0, err
	}
	return addr, nil
}

func writeToSharedMemory(addr uintptr, s pkg.MyStruct) {
	ptr := (*pkg.MyStruct)(unsafe.Pointer(addr))
	*ptr = s
}
func main() {
	myStruct := pkg.MyStruct{Field1: 42}
	copy(myStruct.Field2[:], "Hello")

	addr, err := createSharedMemory()
	if err != nil {
		log.Fatal(err)
	}

	writeToSharedMemory(addr, myStruct)
	for {
		time.Sleep(time.Second)
	}
}
