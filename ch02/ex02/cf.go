package main

import (
	"bufio"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			converter(arg)
		}
	} else {
		stdin := bufio.NewScanner(os.Stdin)
		for stdin.Scan() {
			converter(stdin.Text())
		}
	}
}

func converter(s string) {

}
