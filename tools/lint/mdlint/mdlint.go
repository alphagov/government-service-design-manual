// Package mdlint implements functions for *very* basic markdown
// linting, since we seem to keep publishing broken things.
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

type syntaxError struct {
	line     int
	column   int
	position int64
}

type lintResult struct {
	brackets *list.List
	parens   *list.List
	enDashes map[int]int
	emDashes map[int]int
}

// Lint reads the provided reader (with an optional associated path)
// and checks the markdown for basic errors. Any errors found are
// sent to the provided out channel
func Lint(reader *bufio.Reader, path string, out chan<- error) {

	lintResult := lint(reader, path, out)

	// Check the results and accumulate any problems
	checkHanging(lintResult.brackets, "bracket", out, path)
	checkHanging(lintResult.parens, "parenthesis", out, path)

	for line := range lintResult.enDashes {
		out <- fmt.Errorf("literal en dash at %s:%d - please use -- instead", path, line)
	}

	for line := range lintResult.emDashes {
		out <- fmt.Errorf("literal em dash at %s:%d - please use --- instead", path, line)
	}
}

func lint(reader *bufio.Reader, path string, out chan<- error) lintResult {
	res := lintResult{list.New(), list.New(), make(map[int]int), make(map[int]int)}
	line := 1
	column := 1
	pos := int64(0)

	// Parse the file
	for {
		r, size, err := reader.ReadRune()
		pos += int64(size)

		if err != nil {
			if err != io.EOF {
				out <- fmt.Errorf("Error reading from %s - %s", path, err)
			}
			break
		}

		switch r {
		case '[':
			res.brackets.PushFront(syntaxError{line, column, pos})
		case ']':
			top := res.brackets.Front()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing bracket at line %d, column %d`, path, line, column)
				out <- usefulError(path, pos, basicError)
			} else {
				res.brackets.Remove(top)
			}
		case '(':
			res.parens.PushFront(syntaxError{line, column, pos})
		case ')':
			top := res.parens.Front()
			if top == nil {
				basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra closing parenthesis at line %d, column %d`, path, line, column)
				out <- usefulError(path, pos, basicError)
			} else {
				res.parens.Remove(top)
			}
		case '–':
			res.enDashes[line]++
		case '—':
			res.emDashes[line]++
		case '\n':
			line++
			column = 0
		}

		column++
	}
	return res
}

func checkHanging(list *list.List, character string, out chan<- error, path string) {
	for e := list.Back(); e != nil; e = e.Prev() {
		syntaxError := e.Value.(syntaxError)
		basicError := fmt.Errorf(`Bad Markdown URL in %s:
	extra opening %s at line %d, column %d`, path, character, syntaxError.line, syntaxError.column)
		out <- usefulError(path, syntaxError.position, basicError)
	}
}

func highlightError(f io.Reader, pos int64) (line int, col int, highlight string) {
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

	return fmt.Errorf("Error parsing Markdown in %s%s",
		f.Name(), extra)
}
