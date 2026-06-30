package main

import (
	"fmt"
	"github.com/Mi3i4/tiny_url/internal/shortener"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "usage: tiny-url-cli <url> [code-length]")
		os.Exit(1)
	}

	url := os.Args[1]

	codeLen := 8
	if len(os.Args) >= 3 {
		var err error
		codeLen, err = strconv.Atoi(os.Args[2])
		if err != nil || codeLen <= 0 {
			fmt.Fprintln(os.Stderr, "code-length must be a positive integer")
			os.Exit(1)
		}
	}

	code, err := shortener.Shorten(url, codeLen)
	if err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(code)
}
