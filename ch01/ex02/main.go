package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		var line string = strconv.Itoa(i) + " " + os.Args[i]
		fmt.Println(line)
	}
}
