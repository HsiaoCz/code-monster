build:
	@go build -o bin/tmp main.go

run: build
	@./bin/tmp

test:
	@go test -v ./...

gen:
	@templ generate
