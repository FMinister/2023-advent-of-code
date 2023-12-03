package helperfunctions

import (
	"bufio"
	"log"
	"os"
)

func OpenAndReadFile(filename string) []string {
	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var fileText []string

	for fileScanner.Scan() {
		fileText = append(fileText, fileScanner.Text())
	}

	return fileText
}
