// *Very* basic markdown linting, since we seem to keep publishing broken things.
package mdlint

import (
	"bufio"
	"fmt"
	"io"
)

type Position struct {
	line   int
	column int
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

func Lint(reader *bufio.Reader, path string, chann chan<- string) {
	// Stack for parsing each, to get the line number where it was encountered
	brackets, parens := NewStack(), NewStack()
	line := 1
	column := 1
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
			brackets.Push(Position{line, column})
		case ']':
			top := brackets.Pop()
			if top == nil {
				chann <- fmt.Sprintf(`Bad Markdown URL in %s:
extra closing bracket at line %d, column %d`, path, line, column)
			}
		case '(':
			parens.Push(Position{line, column})
		case ')':
			top := parens.Pop()
			if top == nil {
				chann <- fmt.Sprintf(`Bad Markdown URL in %s:
extra closing parenthesis at line %d, column %d`, path, line, column)
			}
		case '–':
			enDashes[line]++
		case '\n':
			line++
			column = 0
		}

		column++
	}

	checkHanging := func(stack *Stack, character string) {
		results := make([]Position, stack.Len())
		i := 0
		for top := stack.Pop(); top != nil; top = stack.Pop() {
			results[i] = top.(Position)
			i++
		}

		// Reverse the results, since we used a LIFO data structure to capture them
		for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
			results[i], results[j] = results[j], results[i]
		}

		for _, position := range results {
			chann <- fmt.Sprintf("Bad Markdown URL in %s - extra opening %s at line %d, column %d", path, character, position.line, position.column)
		}
	}

	// Check the results and accumulate any problems
	checkHanging(brackets, "bracket")
	checkHanging(parens, "parenthesis")

	for line, _ := range enDashes {
		chann <- fmt.Sprintf("literal en dash at %s:%d - please use -- instead", path, line)
	}
}
