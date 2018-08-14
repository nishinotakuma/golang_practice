// ./ex10 1000
package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%s\n", comma(os.Args[1]))
}

func comma(s string) string {
	b := &bytes.Buffer{}
	pre := len(s) % 3
	if pre == 0 {
		pre = 3
	}
	b.WriteString(s[:pre])
	for i := pre; i < len(s); i += 3 {
		b.WriteByte(',')
		b.WriteString(s[i : i+3])
	}
	return b.String()
}
