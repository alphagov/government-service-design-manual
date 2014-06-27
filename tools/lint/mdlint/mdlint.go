// *Very* basic markdown linting, since we seem to keep publishing broken things.
package mdlint

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

type SyntaxError struct {
	line     int
	column   int
	position int64
}

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

func Lint(reader *bufio.Reader, path string, chann chan<- error) {
	// Stack for parsing each, to get the line number where it was encountered
	brackets, parens := NewStack(), NewStack()
	line := 1
	column := 1
	enDashes := make(map[int]int)
	var pos int64 = 0

	// Parse the file
	for {
		r, size, err := reader.ReadRune()
		pos += int64(size)

		if err != nil {
			if err != io.EOF {
				chann <- fmt.Errorf("Error reading from %s - %s", path, err)
			}
			break
		}

		switch r {
		case '[':
			brackets.Push(SyntaxError{line, column, pos})
		case ']':
			top := brackets.Pop()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing bracket at line %d, column %d`, path, line, column)
				chann <- usefulError(path, pos, basicError)
			}
		case '(':
			parens.Push(SyntaxError{line, column, pos})
		case ')':
			top := parens.Pop()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing parenthesis at line %d, column %d`, path, line, column)
				chann <- usefulError(path, pos, basicError)
			}
		case '–':
			enDashes[line]++
		case '\n':
			line++
			column = 0
		}

		column++
	}

	// Check the results and accumulate any problems
	checkHanging(brackets, "bracket", chann, path)
	checkHanging(parens, "parenthesis", chann, path)

	for line, _ := range enDashes {
		chann <- fmt.Errorf("literal en dash at %s:%d - please use -- instead", path, line)
	}
}

func checkHanging(stack *Stack, character string, chann chan<- error, path string) {
	results := make([]SyntaxError, stack.Len())
	i := 0
	for top := stack.Pop(); top != nil; top = stack.Pop() {
		results[i] = top.(SyntaxError)
		i++
	}

	// Reverse the results, since we used a LIFO data structure to capture them
	for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
		results[i], results[j] = results[j], results[i]
	}

	for _, syntaxError := range results {
		basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra opening %s at line %d, column %d`, path, character, syntaxError.line, syntaxError.column)
		chann <- usefulError(path, syntaxError.position, basicError)
	}
}

func highlightError(f io.Reader, pos int64) (line, col int, highlight string) {
	line = 1
	br := bufio.NewReader(f)
	lastLine := ""
	thisLine := new(bytes.Buffer)
	for n := int64(0); n < pos; n++ {
		b, err := br.ReadByte()
		if err != nil {
			break
		}
		if b == '\n' {
			lastLine = thisLine.String()
			thisLine.Reset()
			line++
			col = 1
		} else {
			if utf8.RuneStart(b) {
				col++
			}
			thisLine.WriteByte(b)
		}
	}
	if line > 1 {
		highlight += fmt.Sprintf("%5d: %s\n", line-1, lastLine)
	}
	highlight += fmt.Sprintf("%5d: %s\n", line, thisLine.String())
	highlight += fmt.Sprintf("%s^\n", strings.Repeat(" ", col+5))
	return
}

func usefulError(path string, pos int64, basicError error) error {
	f, err := os.Open(path)

	if err != nil {
		// We don't have a file that we can read? Maybe just operating on a buffer.
		return basicError
	}

	defer f.Close()

	if _, serr := f.Seek(0, os.SEEK_SET); serr != nil {
		log.Fatalf("seek error: %v", serr)
	}

	line, col, highlight := highlightError(f, pos)
	extra := fmt.Sprintf(":\nError at line %d, column %d (file offset %d):\n%s",
		line, col, pos, highlight)

	return fmt.Errorf("Error parsing Markdown in %s%s\n",
		f.Name(), extra)
}
