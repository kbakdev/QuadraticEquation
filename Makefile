#!make
SHELL := /bin/bash

.PHONY: run
run: ##@run Runs QuadraticEquation.
	go run main.go

.PHONY: test
test: ##@test Runs tests.
	go test -cover -v

.PHONY: test-cover
test-cover: ##@test Runs tests, this will display the HTML representation of the source file where you can clearly see which lines are or are not covered by the test.
	go test -coverprofile=coverage.out

.PHONY: test-cover-html
test-cover-html: ##@test convert to HTML, more human readable. Open it with Web Browser.
	go tool cover -html=coverage.out