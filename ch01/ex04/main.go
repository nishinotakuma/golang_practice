// go run main.go test1.txt test2.txt
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var counts = map[string]int{}
	var fileNames = map[string]map[string]string{}
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileNames)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileNames)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			keys := []string{}
			for k := range fileNames[line] {
				keys = append(keys, k)
			}
			fmt.Printf("%d\t%v\t%s\n", n, keys, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileNames map[string]map[string]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if _, v := fileNames[input.Text()]; !v {
			fileNames[input.Text()] = map[string]string{}
		}
		fileNames[input.Text()][f.Name()] = ""
	}
}
