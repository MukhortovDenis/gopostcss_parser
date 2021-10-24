package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func input(urlToCss string) {
	if urlToCss == "" {
		log.Fatal("Не получил строку css")
	}
	file, err := os.Open(urlToCss)
	if err != nil {
		log.Fatal(err)
	}
	wr := bytes.Buffer{}
	sc := bufio.NewScanner(file)
	for sc.Scan() {
		wr.WriteString(sc.Text())
	}
	fmt.Println(wr.String())
	defer file.Close()
}
