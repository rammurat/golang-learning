package main

import "fmt"

func main() {
	for i := 0; i <5; i++ {
		fmt.Printf("%d\n", i) // numbers
		fmt.Printf("%q\n", i) // UTF 8 format
	}
}