PROJECTNAME=$(shell basename "$(PWD)")

## help: Get all available command
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'

## lint: Run code linter
lint:
	golangci-lint run

## test: Run uunit test and coverage test
test:
	go test -coverprofile=coverage.out ./... && go tool cover -func coverage.out && go tool cover -html=coverage.out

## build: Build binary
build:
	go build -o ./build/url-shortener

## run: Build binary and execute the file
run: build
	./build/url-shortener
