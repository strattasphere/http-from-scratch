package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func _main() {
	f, err := os.Open("messages.txt")
	if err != nil {
		log.Fatalf("could not open: %s: %s\n", "messages.txt", err)
	}

	lines := getLinesFromReader(f)
	for line := range lines {
		fmt.Printf("read: %s\n", line)
	}

}

func getLinesFromReader(f io.ReadCloser) <-chan string {
	lines := make(chan string, 8)

	go func() {
		defer f.Close()
		defer close(lines)

		currentLineContents := ""
		for {
			b := make([]byte, 8)
			n, err := f.Read(b)
			if err != nil {
				if currentLineContents != "" {
					lines <- currentLineContents
				}
				if errors.Is(err, io.EOF) {
					break
				}
				fmt.Printf("error: %s\n", err.Error())
				return
			}

			str := string(b[:n])
			parts := strings.Split(str, "\n")
			for i := 0; i < len(parts)-1; i++ {
				lines <- fmt.Sprintf("%s%s", currentLineContents, parts[i])
				currentLineContents = ""
			}

			currentLineContents += parts[len(parts)-1]

		}
	}()

	return lines

}
