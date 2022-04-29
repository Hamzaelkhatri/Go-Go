package main

import (
	"syscall"
)

func main() {
	var buff [1024]byte

	syscall.Read(0, buff[:])
	syscall.Write(1, buff[:])
}
