package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.OpenFile(
		"hello_world.txt",
		// Create if file does not exist.
		// Read / Write.
		// Open the file and delete the file.
		// File permissions : 644(Owner, Group, Others)
		// Read(4), Write(2), Execute(1)
		os.O_CREATE|os.O_RDWR|os.O_TRUNC,
		os.FileMode(644))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	n := 0
	str := "Hello"
	n, err = file.Write([]byte(str)) // Convert str to []byte slice.
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(n, "Data saved")

	fi, err := file.Stat() // Read the file information.
	if err != nil {
		fmt.Println(err)
		return
	}

	var data = make([]byte, fi.Size()) // Create slices for file size.
	// After writing function above, the pointer in the file moved,
	// so use the Seek() function to move the pointer first.
	file.Seek(0, os.SEEK_SET)
	n, err = file.Read(data)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(n, "Complete byte read")
	fmt.Println(string(data))
}