package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func getEnvVar(varName string) string {
    err := godotenv.Load(".env")
    if err != nil {
        fmt.Printf("Error loading .env file")
    }

    value := os.Getenv(varName)
    if value == "" {
        fmt.Printf("Error: %s environment variable not set\n", varName)
    }
    return value
}

func main() {
    webSiteURL := getEnvVar("URL")
    folderName := getEnvVar("FOLDER_NAME")
    selector := getEnvVar("SELECTOR")
    attr := getEnvVar("ATTRIBUTE")
    targetExt := getEnvVar("TARGET_EXT")

    doc, err := fetchDocument(webSiteURL)
    if err != nil {
        panic(fmt.Sprintf("Error fetching document: %s", err))
    }
    createFolder(folderName)
    download(doc, folderName, selector, attr, targetExt)
}
