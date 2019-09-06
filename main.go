package main

import "fmt"

func main() {
	//init application
	displayBanner("%", 100, 3, "Wordpress CLI")
	displayProjectList()

	//get user choice of project
	s := ScannerUser{"Faites votre choix :"}
	uChoice := sProject(s)
	fmt.Println(uChoice)
	n, p := getProject(uChoice)

	fmt.Println(n)
	fmt.Println(p)

}
