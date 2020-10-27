.EXPORT_ALL_VARIABLES:

all: deps run

deps:
	@dep ensure

build:
	@go build -o bin  main.go

run:
	@go run main.go





