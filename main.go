package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
)

func _main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatal("error", "error", err)
	}

	str := ""
	for {
		data := make([]byte, 8)
		n, err := f.Read(data)
		if err != nil {
			break
		}
		//fmt.Printf("data: %s\n", string(data))
		//fmt.Printf("no of bytes read: %d\n", n)

		// set data equal to either the entire length of data (8 bytes) or
		// cut short when there are less than 8 bytes
		data = data[:n]
		// IndexByte tells you where a byte is in a []bytes
		// if byte does not exist (returns -1), skip this step
		if i := bytes.IndexByte(data, '\n'); i != -1 {
			// add stringified []byte (to length i - because that is where the index is)
			// to the empty string
			str += string(data[:i])
			// set data to latter part of data from
			data = data[i+1:]
			fmt.Printf("read: %s\n", str)
			str = ""
		}

		str += string(data)
	}

	if len(str) != 0 {
		fmt.Printf("read: %s\n", str)
	}

}
