// Package main provides a starting point to the program
package main

import (
	"fmt"

	"github.com/iamveekthorr/utils"
)

func main() {
	var url string

	fmt.Print("Enter the url: ")
	urls := make(map[string]string)

	fmt.Scanln(&url)

	// transform to short code;

	// make new object
	urls[url] = utils.MakeShortCode(7)
	fmt.Printf("%v", urls)
}
