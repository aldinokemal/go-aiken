package main

import (
	"fmt"
	aiken "github.com/aldinokemal/go-aiken"
)

func main() {
	result, _ := aiken.ReadAiken("aiken-example.txt")
	fmt.Println(result)
}
