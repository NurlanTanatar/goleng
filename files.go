package main

import (
	"bufio"
	"fmt"
	"os"
)

func WorkWithFiles() {
	file1, _ := os.Create("text.txt")
	file1.WriteString("any info on this?")
	if fileInfo, err := file1.Stat(); err != nil {
		panic(err)
	} else {
		fmt.Printf("fileInfo: %v\n", fileInfo)
	}
	if data, err := os.ReadFile(file1.Name()); err != nil {
		panic(err)
	} else {
		fmt.Println(string(data))
	}
	defer file1.Close()
	defer os.Remove(file1.Name())

}

func ReadingBigFile(filename string) {
	var count int = 1
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	for {
		s, _ := reader.ReadString(';')
		if s == "0;" {
			break
		}
		count++
	}
	fmt.Println(count)
	defer file.Close()
}
