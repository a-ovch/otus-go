package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	const msg = "Hello, OTUS!"

	fmt.Print(stringutil.Reverse(msg))
}
