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
	assertSuffix(t, <-collector, "extra closing bracket at line 3, column 15")
	assertChannelIsEmpty(t, collector)
}

func TestMissingEndingBracketMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link(http://example.com/)
this is another line
`)
	assertSuffix(t, <-collector, "extra opening bracket at line 3, column 1")
	assertChannelIsEmpty(t, collector)
}

func TestMissingStartingParenthesIsMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link]http://example.com/)
this is another line
`)
	assertSuffix(t, <-collector, "extra closing parenthesis at line 3, column 36")
	assertChannelIsEmpty(t, collector)
}

func TestMissingEndingParenthesIsMarkdown(t *testing.T) {
	collector := goLint(`
this is the first line
[this is a link](http://example.com/
this is another line
`)
	assertSuffix(t, <-collector, "extra opening parenthesis at line 3, column 17")
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

func TestLiteralEmDashesAreBad(t *testing.T) {
	collector := goLint(`
this is the first line — OK?
[this is a link](http://example.com/)
this is another line
`)
	assertPrefix(t, <-collector, "literal em dash at")
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

	assertSuffix(t, <-collector, "extra opening parenthesis at line 3, column 25")
	assertSuffix(t, <-collector, "extra opening parenthesis at line 7, column 1")
	assertChannelIsEmpty(t, collector)
}

func assertPrefix(t *testing.T, problem error, prefix string) {
	if !strings.HasPrefix(problem.Error(), prefix) {
		t.Log(problem)
		t.Fail()
	}
}

func assertSuffix(t *testing.T, problem error, suffix string) {
	if !strings.HasSuffix(problem.Error(), suffix) {
		t.Log(problem)
		t.Fail()
	}
}

func assertChannelIsEmpty(t *testing.T, chann chan error) {
	for e := range chann {
		t.Log(e)
		t.Fail()
	}
}

func goLint(markdown string) chan error {
	collector := make(chan error)
	go func() {
		defer close(collector)
		reader := bufio.NewReader(strings.NewReader(markdown))
		mdlint.Lint(reader, "test.md", collector)
	}()
	return collector
}
