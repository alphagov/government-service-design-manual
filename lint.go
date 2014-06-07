// *Very* basic markdown linting, since we seem to keep publishing broken things.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func goWalk(location string) chan string {
	chann := make(chan string) // unbuffered channel of synchronous error messages

	go func() {
		// Once we've walked all the files, close this channel to avoid deadlocks
		defer close(chann)

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
			brackets, parens, line := 0, 0, 1
			enDashes := make(map[int] int)

			// Parse the file
			for {
				r, _, err := reader.ReadRune()

				if err != nil {
					break
				}

				switch r {
				case '[':
					brackets++
				case ']':
					brackets--
				case '(':
					parens++
				case ')':
					parens--
				case '–':
					enDashes[line]++
				case '\n':
					line++
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
			case len(enDashes) > 0:
				for line, _ := range enDashes {
					chann <- fmt.Sprintf("literal en dash at %s:%d - please use -- instead", path, line)
				}
			}

			return nil
		})
	}()

	return chann
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		// Use all available cores if not otherwise specified
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

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
