package main

import "fmt"
import "math/rand"
import "./utils"
import "github.com/rammurat/golang-learning/packages/live-utils"

func main() {
	// print random number
	fmt.Println(rand.Intn(100))

	// print variable from your package
	fmt.Println(localUtil.StoreId)

	// Print langId from server util
	fmt.Println(liveUtil.LangId)
}
