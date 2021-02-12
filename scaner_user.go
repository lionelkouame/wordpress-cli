package main

import (
	"bufio"
	"fmt"
	"os"
)

type ScannerUser struct {
	label string
}

func sProject(s ScannerUser) string {
	var strI string
	fmt.Println(s.label)

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strI = scanner.Text()
	}

	return strI
}
