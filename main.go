package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
	"time"
)

var files = []string{"main.go", "Makefile", "README.md"}

// make concurrent
func main() {
	start := time.Now() // speed tester

	// setup subcommand
	newCmd := flag.NewFlagSet("new", flag.ExitOnError)

	allFunctions()

	fmt.Println(time.Since(start)) // speed tester
}

func allFunctions() {
	// create files
	create()

	// exec appt
	exech()

	// write to files
	writef() // write main
	writeM() // write make

	// fmt main
	format()

}

// create files - makefile readme
func create() {
	for _, file := range files {
		os.Create(file)
	}

	log.Println("files created")
}

// execute go mod init
func exech() {
	// get path info
	pwd, _ := os.Getwd()

	basePath := path.Base(pwd)

	// go mod init
	cmd, _ := exec.Command("go", "mod", "init", basePath).Output()

	fmt.Println(string(cmd))
}

// writes main.go
func writef() {
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
func writeM() {
	initMake := `
	build: 
		@go build -o testFolder/go-Prepper

	run: build
		@./testFolder/go-Prepper
	`

	f, _ := os.OpenFile("Makefile", os.O_RDWR|os.O_APPEND, 0644)
	// f, _ := os.Create("main.go")

	defer f.Close()

	_, err := f.WriteString(initMake)
	if err != nil {
		fmt.Println(err)
	}

}
