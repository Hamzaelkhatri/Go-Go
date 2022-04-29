package main

func go_print_numbers() {
	var i byte = 48

	for ; i <= 57; i++ {
		go_putchar(i)
	}
}
