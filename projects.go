package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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

	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	xml.Unmarshal(byteValue, &p)

	return p
}

func displayProjectList() {
	projects := readProjectsXml()

	fmt.Println("Choissisez votre projet wordpress")
	for i := 0; i < len(projects.Projects); i++ {
		fmt.Println("Choice => [" + strconv.FormatInt(int64(i), 10) + "]")
		fmt.Println("Project Name :" + projects.Projects[i].Name)
		fmt.Println("    path: " + projects.Projects[i].Path)
	}
}

func getProject(cursor string) (name, path string) {
	projects := readProjectsXml()
	for i := 0; i < len(projects.Projects); i++ {
		iStr := strconv.FormatInt(int64(i), 10)
		if cursor == iStr {
			name = projects.Projects[i].Name
			path = projects.Projects[i].Path
		}
	}
	return name, path
}

type Project struct {
	XMLName xml.Name `xml:"project"`
	Name    string   `xml:"name"`
	Path    string   `xml:"path"`
}
