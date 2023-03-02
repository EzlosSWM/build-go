build: 
	@go build -o testFolder/build-go

run: build
	@cd testFolder && ./build-go new