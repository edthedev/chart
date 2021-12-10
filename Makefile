
# https://github.com/trstringer/cli-debugging-cheatsheets/blob/master/go.md
#

all: build

clean:
	rm chart.exe

build: setup
	go build chart.go

test:
	go test

run:
	go run chart.go

setup:
	go get -u github.com/guptarohit/asciigraph
