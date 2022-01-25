package main

import (
	"fmt"
	"path"
)

func main() {
	_, file := path.Split("css/main.css")

	fmt.Println("file:", file)
}
