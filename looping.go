package main

import "fmt"

func main() {
	ini := 1
	for ini <= 5 {
		fmt.Printf("Welcome %d times.\n", ini)
		ini = ini + 1
	}

	for i := 6; i <= 10; i++ {
		fmt.Printf("Welcome %d times.\n", i)
	}
}
