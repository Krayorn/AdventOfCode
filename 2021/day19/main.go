package main

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var content string

func main() {
	values := strings.Split(content, "\n")
}
