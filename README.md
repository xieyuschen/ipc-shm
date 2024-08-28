# ipc-shm

This is a POC to see whether it's possible to communicate between two processes
via IPC of sharing memory way. The `shm` is used to refer the _sharing memory_ way in IPC.

## Producing and Consuming

### Simple Structure
When we keep writing `pkg.Message` into the shm block and read from another process,
it looks normal and won't cause some fatal error when the value doesn't actual exist in
the certain memory slot.

```go
2024/08/28 14:41:43 Field1: 16, Field2: producer
2024/08/28 14:41:44 Field1: 17, Field2: producer
2024/08/28 14:41:45 Field1: 0, Field2: 
2024/08/28 14:41:46 Field1: 0, Field2:
```

