// ./ex11 -34411.534534
package main

import (
	"bytes"
	"fmt"
	"os"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}

func comma(s string) string {
	b := bytes.Buffer{}
	intStart := 0
	if s[0] == '+' || s[0] == '-' {
		b.WriteByte(s[0])
		intStart = 1
	}
	intEnd := strings.Index(s, ".")
	if intEnd == -1 {
		intEnd = len(s)
	}
	intPart := s[intStart:intEnd]
	pre := len(intPart) % 3
	if pre > 0 {
		b.Write([]byte(intPart[:pre]))
		if len(intPart) > pre {
			b.WriteString(",")
		}
	}
	for i, c := range intPart[pre:] {
		if i%3 == 0 && i != 0 {
			b.WriteString(",")
		}
		b.WriteRune(c)
	}
	b.WriteString(s[intEnd:])
	return b.String()
}
