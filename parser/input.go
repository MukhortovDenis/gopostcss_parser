package main

import (
	"bufio"
	"bytes"
	"log"
	"os"
)

func input(urlToCss string) bytes.Buffer {
	if urlToCss == "" {
		log.Fatal("Не получил строку css")
	}
	file, err := os.Open(urlToCss)
	if err != nil {
		log.Fatal(err)
	}
	buffer := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		buffer.WriteString(sc.Text())
	}
	defer file.Close()
	return buffer
	// if &buffer[0] == "\uFEFF" || buffer[0] == "\uFFFE" {

	// }
}
