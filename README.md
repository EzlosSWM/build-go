# build-go
#### `build-go` is CLI tool that generates boiler plate Go code for quick setup. 

I found myself always creating the same couple of files to start writing code so I decided to make that process quicker.

Here's a short list of files that is created:  
**Build:** 
#### 1. `main.go`
- Generates `hello world` code.
#### 2. `go.mod`
- Generates a `mod` file with the name of the folder
#### 3. `README.md`
#### 4. `Makefile`
- Generates a `Makefile` with a script with `build`, `run` & `test` cases.

## Usage 
1. Clone this repo  
`git clone https://github.com/EzlosSWM/build-go.git`
2. Navigate to the directory  
`cd build-go`
3. Build 
- `make build` 