package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func fetchDocument(websiteURL string) (*goquery.Document, error) {
	res, err := http.Get(websiteURL)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		fmt.Println("Error making request:", err)
		return nil, err
	}

	return doc, nil
}

func createImgFile(folderName, targetExt string, i int) *os.File {
	itemName := fmt.Sprintf("%d.%s", i, targetExt)
	filepath := path.Join(folderName, itemName)

	imgFile, err := os.Create(filepath)
	if err != nil {
		panic(fmt.Sprintf("Error creating image file: %s", err))
	}

	return imgFile
}

func fetchImg(s *goquery.Selection, attr string) io.ReadCloser {
	imgURL, exist := s.Attr(attr)
	if exist != true {
		panic(fmt.Sprintf("Error finding image url: %s", imgURL))
	}
	// Remove all whitespaces
	imgURL = strings.Join(strings.Fields(imgURL), "")

	imgRes, err := http.Get(imgURL)
	if err != nil {
		panic(fmt.Sprintf("Error creating image file: %s", err))
	}

	return imgRes.Body
}

func download(doc *goquery.Document, folderName, selector, attr, targetExt string) {
	fmt.Println("Downloading...")

	doc.Find(selector).Each(func(i int, s *goquery.Selection) {
		imgFile := createImgFile(folderName, targetExt, i)
		defer imgFile.Close()

		img := fetchImg(s, attr)
		defer img.Close()

		_, err := io.Copy(imgFile, img)
		if err != nil {
			panic(fmt.Sprintf("Error creating image file: %s", err))
		}

		fmt.Println("Downloaded:", i)
	})

	fmt.Println("Done!")
}
