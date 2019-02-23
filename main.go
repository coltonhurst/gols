package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	files, err := ioutil.ReadDir(".") // get the files, folders, etc. in the current directory

	if err != nil { // if there's an error, log it
		log.Fatal(err)
	}

	if len(files) == 0 { // if there are no files, quit
		os.Exit(0)
	}

	for _, file := range files { // loop through the file array and print the files & directories
		fmt.Printf("%s\n", file.Name())
	}
}
