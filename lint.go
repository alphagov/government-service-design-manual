// *Very* basic markdown linting, since we seem to keep publishing broken things.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func goWalk(location string) chan string {
	chann := make(chan string) // unbuffered channel of synchronous error messages

	go func() {
		filepath.Walk(location, func(path string, f os.FileInfo, err error) error {

			// Only parse things that look like markdown
			if !strings.HasSuffix(path, ".md") {
				return nil
			}

			file, err := os.Open(path)

			if err != nil {
				panic(err)
			}

			defer file.Close()

			reader := bufio.NewReader(file)
			brackets, parens := 0, 0

			// Parse the file
			for {
				b, err := reader.ReadByte()

				if err != nil {
					break
				}

				switch b {
				case '[':
					brackets++
				case ']':
					brackets--
				case '(':
					parens++
				case ')':
					parens--
				}
			}

			// Check the results and accumulate any problems
			switch {
			case brackets < 0:
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra closing bracket?", path)
			case brackets > 0:
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra opening bracket?", path)
			case parens > 0:
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra closing parenthesis?", path)
			case parens > 0:
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra opening parenthesis?", path)
			}

			return nil
		})

		// Once we've walked all the files, close this channel to avoid deadlocks
		defer close(chann)
	}()

	return chann
}

func main() {
	flag.Parse()
	location := flag.Arg(0)

	if len(location) == 0 {
		fmt.Printf("Usage: go run lint.go path/to/markdown/dir\n")
		os.Exit(2)
	}

	chann := goWalk(location)

	exitCode := 0

	for msg := range chann {
		fmt.Println(msg)
		exitCode = 1
	}

	os.Exit(exitCode)
}
