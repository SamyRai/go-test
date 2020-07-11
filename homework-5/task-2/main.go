package main

import (
	"fmt"
	"log"
	"os"
)

func openFile(fileName string) (file *os.File) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return
}

func readFile(file *os.File) []byte {
	stat, err := file.Stat()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	return bs
}

func fileContentAsString(fileName string) string {
	file := openFile(fileName)
	defer file.Close()

	return string(readFile(file))
}

func main() {
	fmt.Println(fileContentAsString("main.go"))
}
