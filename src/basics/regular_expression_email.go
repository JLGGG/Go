package main

import (
	"fmt"
	"regexp"
)

func main() {
	var checkEmailAddress, _ = regexp.Compile(
		//The string must be in the front because ^[_a-z0-9+-.]+@ section is appended ^.
		//Also, there should be a character because it has a +.
		//The part (\\.[a-z0-9-]+)* is the sub-domain part. There is * at the end, so that part may or may not exist.
		//(\\.[a-z]{2,4})$ is a TLD domain. Only lowercase letters can be used, and should be at the end of the string.
		//You can specify the minimum number, maximum number, such as {2,4}.
		//For example, you can express .com, .kr, .co, .info, etc. Adjust the length in a flexible way as the TLD domain may be longer than four digits.
		"^[_a-z0-9+-.]+@[a-z0-9-]+(\\.[a-z0-9-]+)*(\\.[a-z]{2,4})$",
		)
	//True
	fmt.Println(checkEmailAddress.MatchString("dgist@dgist.ac.kr"))
	fmt.Println(checkEmailAddress.MatchString("dgist+korea@example.com"))
	fmt.Println(checkEmailAddress.MatchString("hello-world@example.com"))
	//False
	fmt.Println(checkEmailAddress.MatchString("@korea.com"))
	fmt.Println(checkEmailAddress.MatchString("korea@korea"))
	fmt.Println(checkEmailAddress.MatchString("korea@korea.cooom"))
}