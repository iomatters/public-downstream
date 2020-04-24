package main

import (
	"fmt"
)

var mypassword="notfordisclosure"

func main() {
	if len(mypassword) > 0 {
		fmt.Println("The password is set")
	}
}
