.PHONY: build
build:
	go build -o ./build/

.PHONY: install
install:
	go get -d

.PHONY: run
run:
	./app