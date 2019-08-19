package main

import (
	"fmt"
)

func main() {
	//init application
	displayBanner("%", 100, 3, "Wordpress CLI")
	config := initConfig()

	displayProjectList()

	fmt.Println(config)
}
