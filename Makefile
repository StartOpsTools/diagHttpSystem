PROJECTNAME="digHttpSystem"

all: clean deps linux

clean:
	rm -rf dist/

deps:
	go mod vendor

linux:
	mkdir -p dist/linux
	go build -o dist/linux/digHttpSystem cmd/main.go

