package main

import (
	"fmt"
	"syscall"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
)

func main() {
	msg := pkg.Message{
		Field1: 1,
	}
	copy(msg.Field2[:], "producer")

	addr, err := pkg.CreateSharedMemory()
	beginningAddr := addr
	defer func() {
		_, _, err = syscall.Syscall(syscall.SYS_SHMDT, beginningAddr, 0, 0)
		fmt.Println("++", err)
	}()
	if err != nil {
		panic(err)
	}

	for i := 0; i < 2; i++ {
		msg.Field1 = int32(i)
		input := pkg.MessageFixedBatch{msg, msg}
		writeToSharedMemory(addr, input)
		addr += unsafe.Sizeof(pkg.MessageFixedBatch{})
		time.Sleep(time.Second)
		fmt.Println("write a new structure:", i)
	}
	for {

	}
}

func writeToSharedMemory(addr uintptr, s pkg.MessageFixedBatch) {
	ptr := (*pkg.MessageFixedBatch)(unsafe.Pointer(addr))
	*ptr = s
}
