build: 
	@go build -o bin/build-go

run: build
	@./build-go new -default 

test: 
	@go test -v ./...