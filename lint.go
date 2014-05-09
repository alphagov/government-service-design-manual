// *Very* basic markdown linting, since we seem to keep publishing broken things.
package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func visit(path string, f os.FileInfo, err error) error {
	if !strings.HasSuffix(path, ".md") {
		return nil
	}

	file, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)
	braces := 0

	for {
		b, err := reader.ReadByte()

		if err != nil {
			break
		}

		switch b {
		case '[':
			braces++
		case ']':
			braces--
		}
	}

	switch {
	case braces < 0:
		fmt.Printf("Bad Markdown URL in %s - extra closing brace?\n", path)
	case braces > 0:
		fmt.Printf("Bad Markdown URL in %s - extra opening brace?\n", path)
	}

	return nil
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	filepath.Walk(root, visit)
}
