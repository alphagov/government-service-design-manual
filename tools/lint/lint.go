// *Very* basic markdown linting, since we seem to keep publishing broken things.
package main

import (
	"./mdlint"
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func goWalk(location string, exclude []string) chan error {
	chann := make(chan error) // unbuffered channel of synchronous error messages

	go func() {
		// Once we've walked all the files, close this channel to avoid deadlocks
		defer close(chann)

		filepath.Walk(location, func(path string, f os.FileInfo, err error) error {

			// Only parse things that look like markdown
			if !strings.HasSuffix(path, ".md") {
				return nil
			}

			for _, prefixDir := range exclude {
				if strings.HasPrefix(path, prefixDir) {
					return nil
				}
			}

			file, err := os.Open(path)

			if err != nil {
				panic(err)
			}

			defer file.Close()

			reader := bufio.NewReader(file)

			mdlint.Lint(reader, path, chann)

			return nil
		})
	}()

	return chann
}

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s\n", os.Args[0])
	helpstring := `
go run lint.go [--exclude csv-exclude-paths] path/to/markdown/dir

path/to/markdown/dir - mandatory directory which will be recursively linted
csv-exclude-paths    - CSV list of paths to exclude
`
	fmt.Fprintf(os.Stderr, helpstring)
	os.Exit(2)
}

func main() {
	if os.Getenv("GOMAXPROCS") == "" {
		// Use all available cores if not otherwise specified
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var (
		excludePaths = flag.String("exclude", "", "Optional CSV list of paths to exclude")
	)

	flag.Usage = usage
	flag.Parse()

	location := flag.Arg(0)
	if len(location) == 0 {
		flag.Usage()
	}

	var excludeList []string

	if len(*excludePaths) > 0 {
		excludeList = strings.Split(*excludePaths, ",")
	} else {
		excludeList = []string{}
	}

	chann := goWalk(location, excludeList)

	exitCode := 0

	for err := range chann {
		fmt.Println(err)
		exitCode = 1
	}

	os.Exit(exitCode)
}
