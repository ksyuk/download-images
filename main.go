package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func receiveInput(reader *bufio.Reader, message string) string {
	fmt.Printf("%s:", message)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error receiving input:", err)
	}
	input = strings.TrimSpace(input)

	return input
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	webSiteURL := receiveInput(reader, "URL")
	folderName := receiveInput(reader, "folder name")
	selector := receiveInput(reader, "selector")
	attr := receiveInput(reader, "attribute")
	targetExt := receiveInput(reader, "target ext")

	doc, err := fetchDocument(webSiteURL)
	if err != nil {
		panic(fmt.Sprintf("Error fetching document: %s", err))
	}
	createFolder(folderName)
	download(doc, folderName, selector, attr, targetExt)

}
