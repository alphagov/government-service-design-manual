// *Very* basic markdown linting, since we seem to keep publishing broken things.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type Stack struct {
	top  *Element
	size int
}

type Element struct {
	value interface{}
	next  *Element
}

// NewStack returns a new stack.
func NewStack() *Stack {
	return &Stack{}
}

// Return the stack's length
func (s *Stack) Len() int {
	return s.size
}

// Push a new element onto the stack
func (s *Stack) Push(value interface{}) {
	s.top = &Element{value, s.top}
	s.size++
}

// Remove the top element from the stack and return its value
// If the stack is empty, return nil
func (s *Stack) Pop() (value interface{}) {
	if s.size > 0 {
		value, s.top = s.top.value, s.top.next
		s.size--
		return
	}
	return nil
}

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

			lint(reader, path, chann)

			return nil
		})
	}()

	return chann
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
	helpstring := `
go run lint.go path/to/markdown/dir
`
	fmt.Fprintf(os.Stderr, helpstring)
	os.Exit(2)
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		// Use all available cores if not otherwise specified
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	flag.Usage = usage
	flag.Parse()

	location := flag.Arg(0)
	if len(location) == 0 {
		flag.Usage()
	}

	chann := goWalk(location)

	exitCode := 0

	for msg := range chann {
		fmt.Println(msg)
		exitCode = 1
	}

	os.Exit(exitCode)
}

func lint(reader *bufio.Reader, path string, chann chan<- string) {
	// Stack for parsing each, to get the line number where it was encountered
	brackets, parens := NewStack(), NewStack()
	line := 1
	enDashes := make(map[int]int)

	// Parse the file
	for {
		r, _, err := reader.ReadRune()

		if err != nil {
			if err != io.EOF {
				chann <- fmt.Sprintf("Error reading from %s - %s", path, err)
			}
			break
		}

		switch r {
		case '[':
			brackets.Push(line)
		case ']':
			top := brackets.Pop()
			if top == nil {
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra closing bracket on line %d", path, line)
			}
		case '(':
			parens.Push(line)
		case ')':
			top := parens.Pop()
			if top == nil {
				chann <- fmt.Sprintf("Bad Markdown URL in %s - extra closing parenthesis on line %d", path, line)
			}
		case '–':
			enDashes[line]++
		case '\n':
			line++
		}
	}

	checkHanging := func(stack *Stack, character string) {
		for top := stack.Pop(); top != nil; top = stack.Pop() {
			chann <- fmt.Sprintf("Bad Markdown URL in %s - extra opening %s on line %d", path, character, top)
		}
	}

	// Check the results and accumulate any problems
	checkHanging(brackets, "bracket")
	checkHanging(parens, "parenthesis")

	for line, _ := range enDashes {
		chann <- fmt.Sprintf("literal en dash at %s:%d - please use -- instead", path, line)
	}
}
