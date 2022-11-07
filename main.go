package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if 1 >= len(os.Args) {
		fmt.Println(0)
	} else {
		fmt.Println(len(strings.Split(os.Args[1], " ")))
	}
}
