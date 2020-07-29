package main

import (
	"fmt"
	"regexp"
)

func main() {
	match, _ := regexp.MatchString("He", "Hello world")
	fmt.Println("pattern : \"He\", input string : \"Hello world\", result :", match)

	match, _ = regexp.MatchString("Goo.", "Good")
	fmt.Println("pattern : \"Goo.\", input string : \"Good\", result :", match)

	match, _ = regexp.MatchString("[a-zA-Z0-9]+", "Hello world 2020")
	fmt.Println("pattern : \"[a-zA-Z0-9]+\", input string : \"Hello world 2020\", result :", match)

	match, _ = regexp.MatchString("[a-zA-Z0-9]*", "")
	fmt.Println("pattern : \"[a-zA-Z0-9]*\", input string : \"\", result :", match)

	match, _ = regexp.MatchString("[0-9]+\\.[0-9]+", "3.5")
	fmt.Println("pattern : \"[0-9]+\\\\.[0-9]+\", input string : \"3.5\", result :", match)

	match, _ = regexp.MatchString("[0-9]+-[0-9]+", "3-")
	fmt.Println("pattern : \"[0-9]+-[0-9]+\", input string : \"3-\", result :", match)

	match, _ = regexp.MatchString("[0-9]+-[0-9]*", "3-")
	fmt.Println("pattern : \"[0-9]+-[0-9]*\", input string : \"3-\", result :", match)

	match, _ = regexp.MatchString("[^A-Z]+", "hello")
	fmt.Println("pattern : \"[^A-Z]+\", input string : hello, result :", match)

	match, _ = regexp.MatchString("[^A-Za-z]+", "Hello")
	fmt.Println("pattern : \"[^A-Za-z]+\", input string : Hello, result :", match)
}
