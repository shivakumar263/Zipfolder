package main

import (
	"archive/zip"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filereader, err := zip.OpenReader("Geography.zip")
	if err != nil {
		log.Fatal(err)
	}
	//opened the file
	defer filereader.Close()
	for _, file := range filereader.Reader.File {
		zippedfile, err := file.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer zippedfile.Close()
		targetdir := "./"
		extractedfilepath := filepath.Join(targetdir, file.Name)

		if file.FileInfo().IsDir() {
			fmt.Println("Creating directory:", extractedfilepath)
			os.Mkdir(extractedfilepath, file.Mode())
		} else {
			fmt.Println("Extracting filepath:", file.Name)
			outputFile, err := os.OpenFile(
				extractedfilepath,
				os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())

			if err != nil {
				log.Fatal(err)
			}
			defer outputFile.Close()
		}
	}
}
