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

## build: Compile source code and create binary file
build:
	@echo ">> Building App..."
	@go build -o ./bin/url-shortener main.go
	@echo ">> Finished"

## run: Build binary and execute the file
run: build
	./bin/url-shortener

## start: Start the program
start: build run