package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var width = flag.Int("w", 256, "hash width (256 or 384 or 512)")

func main() {
	flag.Parse()
	var function func(b []byte) []byte
	switch *width {
	case 256:
		function = func(b []byte) []byte {
			h := sha256.Sum256(b)
			return h[:]
		}
	case 384:
		function = func(b []byte) []byte {
			h := sha512.Sum384(b)
			return h[:]
		}
	case 512:
		function = func(b []byte) []byte {
			h := sha512.Sum512(b)
			return h[:]
		}
	default:
		log.Fatal("Unexpected width")
	}
	fmt.Println("ffw")
	fmt.Println(os.Stdin)
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(reader)
	b, err := ioutil.ReadAll(os.Stdin)
	fmt.Println(err)
	fmt.Println(b)
	if err != nil {
		log.Fatal(err)
	}
	a := []byte{0}
	fmt.Printf("%x\n", function(a))
}
