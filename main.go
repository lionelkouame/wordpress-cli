package main

import (
	"fmt"
	"math/rand"
)

func main() {
	displayBanner("%", 100, 3, "Wordpress CLI")
	fmt.Println("My favorite number is", rand.Intn(10))
}
