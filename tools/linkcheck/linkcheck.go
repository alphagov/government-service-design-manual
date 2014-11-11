package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/jabley/markdown"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type link struct {
	URL string
}

func goWalk(location string) chan error {
	chann := make(chan error) // unbuffered channel of synchronous error messages

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

			// fmt.Fprintf(os.Stderr, "Checking <%s>\n", path)

			linkProducer := NewLinkProducer(chann, path)

			parser := markdown.NewParser(&markdown.Extensions{Smart: true})
			parser.Markdown(reader, linkProducer)

			// fmt.Fprintf(os.Stderr, "Checked <%s>\n", path)

			return nil
		})
	}()

	return chann
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
	helpstring := `
go run linkcheck.go path/to/markdown/dir

path/to/markdown/dir - mandatory directory which will be recursively checked for internal links
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

	for err := range chann {
		fmt.Println(err)
		exitCode = 1
	}

	os.Exit(exitCode)
}

type linkProducer struct {
	errors     chan error
	filepath   string
	links      []link
	namedLinks []string
}

func NewLinkProducer(errors chan error, path string) markdown.Formatter {
	return &linkProducer{errors, path, []link{}, []string{}}
}

func (l *linkProducer) Finish() {
	// we've parsed the file, so we should have the full object
	// model to play with.
	for _, link := range l.links {
		if isFragment(link) {
			continue
		}

		l.checkWorks(link)
	}
}

func (l *linkProducer) FormatBlock(tree *markdown.Element) {
	for tree != nil {
		l.elem(tree)
		tree = tree.Next
	}
	return
}

func (l *linkProducer) elem(tree *markdown.Element) {
	switch tree.Key {
	case markdown.LINK:
		{
			if isInternalLink(tree.Contents.Link.URL) {
				l.links = append(l.links, link{tree.Contents.Link.URL})
				// fmt.Fprintf(os.Stderr, "Found link %v\n", tree.Contents.Link.URL)
			}
		}
	}
	l.FormatBlock(tree.Children)
}

func isInternalLink(URL string) (r bool) {
	if strings.HasPrefix(URL, "http://") || strings.HasPrefix(URL, "https://") || strings.HasPrefix(URL, "//") {
		return
	}
	return true
}

func (p *linkProducer) checkWorks(l link) {

	if strings.HasPrefix(l.URL, "service-manual") {
		p.errors <- fmt.Errorf("Bad link to <%v> in <%v>.\nShould it start with a forward slash?", l.URL, p.filepath)
		return
	}

	if !strings.HasPrefix(l.URL, "/service-manual") {
		return
	}

	// fmt.Fprintf(os.Stderr, "Checking link %v\n", l.URL)

	options := getURLOptions(removeFragment(l.URL))

	for _, option := range options {
		if fileExists(option) {
			return
		}
	}

	p.errors <- fmt.Errorf("Bad link to <%v> in <%v>.\nTried options %v", l.URL, p.filepath, options)
}

// /service-manual/the-team/service-manager.html
// =>
// /service-manual/the-team/service-manager.md
//
//
// /service-manual/the-team/service-manager/
// =>
// /service-manual/the-team/service-manager/index.md
// /service-manual/the-team/service-manager/index.html
//
//
// /service-manual/the-team/service-manager
// =>
// /service-manual/the-team/service-manager/index.md
// /service-manual/the-team/service-manager/index.html
// /service-manual/the-team/service-manager.md
//
//
// /service-manual/the-team/service-manager.pdf
// =>
// /service-manual/the-team/service-manager.pdf
func getURLOptions(URL string) []string {
	res := make([]string, 3)

	ext := filepath.Ext(URL)

	switch ext {
	case ".html":
		{
			res[0] = "." + URL[:len(URL)-5] + ".md"
			return res[0:1]
		}
	case "":
		{
			if strings.HasSuffix(URL, "/") {
				res[0] = "." + URL + "index.md"
				res[1] = "." + URL + "index.html"
				return res[0:2]
			} else {
				res[0] = "." + URL + "/index.md"
				res[1] = "." + URL + "/index.html"
				res[2] = "." + URL + ".md"
				return res
			}
		}
	default:
		{
			res[0] = "." + URL
			return res[0:1]
		}
	}
}

func isFragment(l link) (r bool) {
	return strings.HasPrefix(l.URL, "#")
}

func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)

	return err == nil
}

func removeFragment(URL string) string {
	if i := strings.Index(URL, "#"); i != -1 {
		return URL[0:i]
	}
	return URL
}
