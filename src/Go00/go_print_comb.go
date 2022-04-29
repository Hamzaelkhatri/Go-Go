package main

import "syscall"

func go_print_comb() {
	// var j int = 0

	for i := 0; i < 10; i++ {
		for b := 1; b < 10; b++ {
			for c := 2; c < 10; c++ {
				if c != i && i != b && c != b {
					syscall.Write(1, []byte{byte(i + 48)})
					syscall.Write(1, []byte{byte(b + 48)})
					syscall.Write(1, []byte{byte(c + 48)})

					if i == 7 && b == 8 && c == 9 {
						i = 11
						c = 11
						b = 11
						break
					}
					syscall.Write(1, []byte{','})
				}

			}
		}
	}
}

func main() {
	go_print_comb()
}
