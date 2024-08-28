package main

import (
	"fmt"
	"time"
	"unsafe"

	"github.com/xieyuschen/ipc-shm/pkg"
)

// notes
// 1. 'copy(myStruct.Field2[:], "consumer")' cannot change the value of structure,
// 		which is retrieved from the shm

func main() {
	addr, err := pkg.AccessSharedMemory()
	if err != nil {
		panic(err)
	}

	for {
		v := readFromSharedMemory(addr)
		addr += unsafe.Sizeof(pkg.MessageFixedBatch{})
		fmt.Println("the whole array are:")
		for _, msg := range v {

			fmt.Printf("	Field1: %d, Field2: %s\n", msg.Field1, msg.Field2)
		}

		time.Sleep(time.Second)
	}
}

func readFromSharedMemory(addr uintptr) pkg.MessageFixedBatch {
	ptr := (*pkg.MessageFixedBatch)(unsafe.Pointer(addr))
	return *ptr
}
