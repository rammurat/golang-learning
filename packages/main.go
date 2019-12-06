package main

import (
	"fmt"
	"math/rand"
	"utils"
)

func main() {
	// print random number 
	fmt.Println(rand.Intn(100))

	// print variable from your package
	fmt.Println(utils.StoreId)
}