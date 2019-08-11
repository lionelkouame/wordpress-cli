package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

const FILE_CONFIG = "config.yml"

func useExistingConfig() {

}

func createDefaultConfigFile() bool {
	var strInput string
	fmt.Println("Le fichier de configuration n'existe pas voulez-vous le créer ? (y/n)")

	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		strInput = scanner.Text()
	}
	if strInput == "y" {
		//TODO Create  the default file config.yml
	}
	fmt.Println(strInput)

	return true
}

func initConfig() bool {
	files, err := ioutil.ReadDir("./")
	isFileExist := false
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range files {
		fmt.Println(f.Name())
		if FILE_CONFIG == f.Name() {
			//TODO  implement useExistingConfig()
		}
	}
	if !isFileExist {
		useExistingConfig()
	}

	return true
}
