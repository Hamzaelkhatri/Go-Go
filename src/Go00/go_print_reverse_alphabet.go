package main

func go_print_reverse_alphabet() {

	var i byte
	for i = 90 + 32; i >= 65+32; i-- {
		go_putchar(i)
	}
}
