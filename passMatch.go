package main

import (
	"fmt"
	"regexp"
)

func main() {
	fmt.Println("Enter your password")
	var str string
	fmt.Scanln(&str)
	re := regexp.MustCompile(str)
	var pass string = re.FindString("manish")
	if pass == "manish" {
		fmt.Println("You have successfully logged in...")
	} else {
		fmt.Println("Password is wrong")
	}

}
