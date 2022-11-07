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
		str := strings.Trim(os.Args[1], "")

		if len(str) > 0 {
			fmt.Println(len(strings.Split(str, " ")))
		} else {
			fmt.Println(len(str))
		}
	}
}
