all: build run

build:
	go build -o ./bin/app .
	chmod +x ./bin/app

run:
	./bin/app
