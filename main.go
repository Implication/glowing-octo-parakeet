package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	lines, bytes := flag.Bool("l", false, "Count lines"), flag.Bool("b", false, "Count lines")

	flag.Parse()
	fmt.Println(count(os.Stdin, *lines, *bytes))
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)

	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0

	for scanner.Scan() {
		if countBytes {
			wc += len(scanner.Bytes())
		} else {
			wc++
		}
	}

	return wc
}
