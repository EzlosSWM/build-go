build: 
	@go build -o testFolder/build-go

run: build
	@./build-go new -default 

test: 
	@go test -v ./...