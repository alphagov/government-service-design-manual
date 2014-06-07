package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
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

func TestEmptyString(t *testing.T) {
	collector := goLint("")
	assertChannelIsEmpty(t, collector)
}

func TestValidLink(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link](http://example.com/)
this is another line
`)
	assertChannelIsEmpty(t, collector)
}

func TestMissingStartingBracketMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
this is a link](http://example.com/)
this is another line
`)
	assertSuffix(t, <-collector, "extra closing bracket on line 3")
	assertChannelIsEmpty(t, collector)
}

func TestMissingEndingBracketMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link(http://example.com/)
this is another line
`)
	assertSuffix(t, <-collector, "extra opening bracket on line 3")
	assertChannelIsEmpty(t, collector)
}

func TestMissingStartingParenthesIsMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link]http://example.com/)
this is another line
`)
	assertSuffix(t, <-collector, "extra closing parenthesis on line 3")
	assertChannelIsEmpty(t, collector)
}

func TestMissingEndingParenthesIsMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link](http://example.com/
this is another line
`)
	assertSuffix(t, <-collector, "extra opening parenthesis on line 3")
	assertChannelIsEmpty(t, collector)
}

func TestLiteralEnDashesAreBad(t *testing.T) {
	collector := goLint(`
this is the first line – OK?
[this is a link](http://example.com/)
this is another line
`)
	assertPrefix(t, <-collector, "literal en dash at")
	assertChannelIsEmpty(t, collector)
}

func TestLineNumberOfProblemWithBracesIsCaptured(t *testing.T) {
	collector := goLint(`
this is the first line
this is the second line (and has an unclosed parethesis
[this is a link](http://example.com/) on the third line
this is the fourth line
this is the fith line
(for [longer content](link to long content)
it would be helpful to know
where in the file you should be looking to fix problems
		`)

	assertSuffix(t, <-collector, "extra opening parenthesis on line 3")
	assertSuffix(t, <-collector, "extra opening parenthesis on line 7")
	assertChannelIsEmpty(t, collector)
}

func assertPrefix(t *testing.T, problem string, prefix string) {
	if !strings.HasPrefix(problem, prefix) {
		t.Log(problem)
		t.Fail()
	}
}

func assertSuffix(t *testing.T, problem string, suffix string) {
	if !strings.HasSuffix(problem, suffix) {
		t.Log(problem)
		t.Fail()
	}
}

func assertChannelIsEmpty(t *testing.T, chann chan string) {
	for e := range chann {
		t.Log(e)
		t.Fail()
	}
}

func goLint(markdown string) chan string {
	collector := make(chan string)
	go func() {
		defer close(collector)
		reader := bufio.NewReader(strings.NewReader(markdown))
		Lint(reader, "test.md", collector)
	}()
	return collector
}

func Lint(reader *bufio.Reader, path string, chann chan<- string) {
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
		results := make([]interface{}, stack.Len())
		i := 0
		for top := stack.Pop(); top != nil; top = stack.Pop() {
			results[i] = top
			i++
		}

		// Reverse the results, since we used a LIFO data structure to capture them
		for i, j := 0, len(results)-1; i < j; i, j = i+1, j-1 {
			results[i], results[j] = results[j], results[i]
		}

		for _, line := range results {
			chann <- fmt.Sprintf("Bad Markdown URL in %s - extra opening %s on line %d", path, character, line)
		}
	}

	// Check the results and accumulate any problems
	checkHanging(brackets, "bracket")
	checkHanging(parens, "parenthesis")

	for line, _ := range enDashes {
		chann <- fmt.Sprintf("literal en dash at %s:%d - please use -- instead", path, line)
	}
}
