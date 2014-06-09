package mdlint

import (
	"../mdlint"
	"bufio"
	"strings"
	"testing"
)

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
		mdlint.Lint(reader, "test.md", collector)
	}()
	return collector
}
