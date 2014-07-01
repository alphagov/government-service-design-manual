// *Very* basic markdown linting, since we seem to keep publishing broken things.
package mdlint

import (
	"bufio"
	"bytes"
	"container/list"
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

func Lint(reader *bufio.Reader, path string, chann chan<- error) {
	brackets, parens := list.New(), list.New()
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
			brackets.PushFront(SyntaxError{line, column, pos})
		case ']':
			top := brackets.Front()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing bracket at line %d, column %d`, path, line, column)
				chann <- usefulError(path, pos, basicError)
			} else {
				brackets.Remove(top)
			}
		case '(':
			parens.PushFront(SyntaxError{line, column, pos})
		case ')':
			top := parens.Front()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing parenthesis at line %d, column %d`, path, line, column)
				chann <- usefulError(path, pos, basicError)
			} else {
				parens.Remove(top)
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

func checkHanging(list *list.List, character string, chann chan<- error, path string) {
	for e := list.Back(); e != nil; e = e.Prev() {
		syntaxError := e.Value.(SyntaxError)
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
