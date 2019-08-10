package main

import "fmt"

func sayHello() string {
	return "hello!"
}

func displayTextLine(content string, numberContent int, text string, textLen int) {
	space := 10
	var textL string
	var textR string
	var textD string
	cursorFirst := 0

	//first part
	for cursorFirst < ((numberContent / 2) - space) {
		textL += content
		textR += content

		cursorFirst++
	}
	textD += textL + "   " + text + "   " + textR

	fmt.Println(textD)
}

// content: the content of banner.
func displayBanner(content string, numberContent int, numberLine int, text string) {

	cursor := 0
	cursorLine := 0
	textLen := len(text)
	var contentBanner string
	for cursorLine < numberLine {
		var contentLine string

		for cursor < numberContent {
			contentLine += content
			cursor++
		}
		contentBanner += contentLine
		fmt.Println(contentBanner)
		if cursorLine == (numberLine / 2) {
			displayTextLine(content, numberContent, text, textLen)
		}
		cursorLine++
	}
}
