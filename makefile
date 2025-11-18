
.PHONY: build clean

build: ./build/tomato

./build/tomato: $(shell find . -name "*.go" -type f)
	go build -o ./build/tomato

clean:
	rm -f ./build/tomato

rebuild: clean build

