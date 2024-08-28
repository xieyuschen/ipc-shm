package pkg

const ShmId = uintptr(12345)
const Size = 16

type MessageFixedBatch [2]Message

// Message is writen by producer and read by consumer by the IPC via sharing memeory way
type Message struct {
	// note that the flexible 'int' type cannot be used here
	Field1 int32
	Field2 [10]byte
}
