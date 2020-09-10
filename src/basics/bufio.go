package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		// Create if file does not exist.
		// Read / Write.
		// Open the file and delete the file.
		// File permissions : 644(Owner, Group, Others)
		// Read(4), Write(2), Execute(1)
		"hello.txt",
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(0644),
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	// Create a write instance w that follows the io.writer interface.
	w := bufio.NewWriter(file)
	// Write "bufio hello world" to buffer with write instance.
	w.WriteString("bufio hello world")
	// Save contents of buffer to file.
	w.Flush()

	r := bufio.NewReader(file)
	// Get file information.
	fi, _ := file.Stat()
	// Create byte slices by file size.
	s := make([]byte, fi.Size())

	//relocate a position of file using os.SEEK_SET.
	file.Seek(0, os.SEEK_SET)
	// Read the contents of a file into a read instance and save it to b.
	r.Read(s)

	fmt.Println(string(s))

}
