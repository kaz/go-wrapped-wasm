.PHONY: run
run: dist
	python -m http.server

dist: node_modules wrapped.wasm
	npx webpack

node_modules:
	npm i

wrapped.wasm:
	GOOS=js GOARCH=wasm go build -o $@
