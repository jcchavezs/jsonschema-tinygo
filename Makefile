.PHONY: build
build:
	tinygo build -o build/jsonschema -scheduler=none .

run: build
	./build/jsonschema