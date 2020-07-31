package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	str := "hello world"
	n, err = file.Write([]byte(str))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "Complete byte save")

	fInformation, err := file.Stat()
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fInformation.Size())
	file.Seek(0, os.SEEK_SET) //relocate a position of file using os.SEEK_SET.

	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "Complete byte read")
	fmt.Println(string(data))
}
