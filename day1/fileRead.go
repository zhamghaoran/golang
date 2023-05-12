package main

import (
	"fmt"
	"os"
	"strings"
)

type faq struct {
	num      int
	FileName string
}

func main() {
	counts := make(map[string]faq)
	for _, filename := range os.Args[1:] {
		data, err := os.ReadFile(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dump3: %v\n", err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line] = faq{
				num:      counts[line].num + 1,
				FileName: filename,
			}
		}
	}
	for line, n := range counts {
		if n.num > 1 {
			fmt.Printf("%s  %d  %s\n", line, n.num, n.FileName)
		}
	}
}
