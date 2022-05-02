package main

import "fmt"

func main() {
	fmt.Printf("%d\n", 21) //decimal
	fmt.Printf("%b\n", 21) //binary
	fmt.Printf("%o\n", 21) //Octal
	fmt.Printf("%x\n", 29) // Hexadecimal number lowercase
	fmt.Printf("%X\n", 32) // Hexadecimal number uppercase
	fmt.Printf("%#X", 39)  // Hexadecimal number uppercase with prefix - 0X

	for i := 0; i < 100; i++ {
		fmt.Printf("%d\t %b\t %o\t %#X\n", i, i, i, i)

	}
}
