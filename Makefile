compile:
	./compile.sh

clean:
	rm -rf _site/*

aspell:
	find service-manual -type f -name '*.md' -exec aspell -d, --master=british --mode=sgml -c '{}' \;

test:
	go test tools/lint/mdlint/mdlint_test.go

help:
	@echo "Please use 'make <target>' where <target> is one of"
	@echo "  clean     to remove the build artefacts"
	@echo "  compile   to generate the site artefacts - you probably want this one"
	@echo "  aspell    to interactively spellcheck all of the content"
