.PHONY: clean aspell test lint help

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  clean     to remove the build artefacts"
	@echo "  lint      to lint the markdown"
	@echo "  test      to test the lint tools"
	@echo "  aspell    to interactively spellcheck all of the content"

clean:
	rm -rf _site/*

aspell:
	find service-manual -type f -name '*.md' -exec aspell -d, --master=british --mode=sgml -c '{}' \;

test:
	# go install code.google.com/p/go.tools/cmd/vet
	go vet ./tools/lint/
	# go get github.com/golang/lint/golint
	golint tools/lint/
	go test tools/lint/mdlint/mdlint_test.go

lint:
	go run tools/lint/lint.go service-manual
