package main

import "syscall"

func go_is_negative(i int) {
	if i >= 0 {
		syscall.Write(1, []byte{'P'})
	} else {
		syscall.Write(1, []byte{'N'})
	}
}
