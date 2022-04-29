package main

func go_print_alphabet() {

	var i byte = 65
	for i = 65 + 32; i <= 90+32; i++ {
		go_putchar(i)
	}
}

func main() {

	go_print_alphabet()
}
