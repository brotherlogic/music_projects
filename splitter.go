package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")
	fmt.Printf("Found %v lines\n", len(lines))

	bar := 0
	week := 1
	for i, line := range lines {
		if i >= bar {
			fmt.Printf("Week %v\n", week)
			bar = (len(lines) * week) / 52
			week++
		}
		fmt.Printf("%v\n", line)
	}
}
