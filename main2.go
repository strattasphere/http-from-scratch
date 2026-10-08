package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("messages2.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	str := ""
	for {
		b := make([]byte, 8)
		n, err := f.Read(b)
		if err != nil {
			break
		}

		if i := bytes.IndexByte(b, '\n'); i != -1 {
			str += string(b[:i])
			// set data to latter part of data from
			b = b[i+1:]
			fmt.Printf("read: %s\n", str)
			str = ""
		}

		b = b[:n]

		str += string(b)
	}

	if len(str) != 0 {
		fmt.Printf("str: %s", str)
	}
}
