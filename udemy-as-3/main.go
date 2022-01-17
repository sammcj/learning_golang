package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	// Using ReadFile and Println
	// content, err := ioutil.ReadFile(os.Args[1])

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(content))

	//  Using Open and Stdout
	f, err := os.Open(os.Args[1])

	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
