compile:
	./compile.sh

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

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  clean     to remove the build artefacts"
	@echo "  compile   to generate the site artefacts - you probably want this one"
	@echo "  aspell    to interactively spellcheck all of the content"
