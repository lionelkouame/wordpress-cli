package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

type Projects struct {
	XMLName  xml.Name  `xml:"projects"`
	Projects []Project `xml:"project"`
}

func readProjectsXml() Projects {
	var p Projects
	xmlFile, err := os.Open("projects.xml")

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Open  ok")

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &p)

	return p
}

func displayProjectList() {
	projects := readProjectsXml()

	fmt.Println("Choissisez votre projet wordpress")
	for i := 0; i < len(projects.Projects); i++ {
		fmt.Printf("=== Project === ")
		fmt.Println("    Project Name :   " + projects.Projects[i].Name)
		fmt.Println("    path: ")
		fmt.Print(projects.Projects[i].Path)
	}
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	Name    string   `xml:"name"`
	Path    string   `xml:"path"`
}
