package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"path"
)

func AllActions() error {
	basePath, err := dirName()
	if err != nil {
		return err
	}

	// create files
	createFiles()

	// exec appt
	exech(basePath)

	// write to files
	writeMain()         // write main
	writeMake(basePath) // write make

	// fmt main
	format()

	fmt.Println("Files successfully created.")
	return nil
}

// get the current path name
func dirName() (string, error) {
	// get path info
	pwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	return path.Base(pwd), nil
}

// create files - makefile readme
func createFiles() (*os.File, error) {
	var files = []string{"main.go", "Makefile", "README.md"}
	var err error

	for _, file := range files {
		f, err := os.Create(file)
		if err != nil {
			return f, nil
		}
	}

	return nil, err
}

// execute go mod init
func exech(basePath string) {

	// go mod init
	cmd, _ := exec.Command("go", "mod", "init", basePath).Output()

	fmt.Println(string(cmd))
}

// writes main.go
func writeMain() {
	initPKG := `
	package main

	import "fmt"

	func main() {
		fmt.Println("Hello World!")
	}
	`

	f, _ := os.OpenFile("main.go", os.O_RDWR|os.O_APPEND, 0644)
	// f, _ := os.Create("main.go")

	defer f.Close()

	_, err := f.WriteString(initPKG)
	if err != nil {
		fmt.Println(err)
	}

}

// formats main.go
func format() {
	cmd, _ := exec.Command("go", "fmt", "./main.go").Output()

	fmt.Println(string(cmd))
}

// write makefile
func writeMake(p string) {

	initMake := fmt.Sprintf("build:\n\t@go build -o bin/%s\n\nrun: build\n\t@./bin/%s", p, p)

	f, _ := os.OpenFile("Makefile", os.O_RDWR|os.O_APPEND, 0644)

	defer f.Close()

	_, err := f.WriteString(initMake)
	if err != nil {
		fmt.Println(err)
	}

}
