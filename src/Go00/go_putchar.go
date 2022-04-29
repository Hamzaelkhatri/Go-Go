package main

import (
	"syscall"
)

func go_putchar(c byte) {
	syscall.Write(1, []byte{c})
}
