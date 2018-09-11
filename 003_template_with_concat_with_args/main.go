package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {

	name := os.Args[1]
	fmt.Println(os.Args[0])
	fmt.Println(os.Args[1])

	template := fmt.Sprint(`
	<!DOCTYPE html>
	<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Sample Template Page 1</title>
		</head>
		<body>
			<h1>` + name + `</h1>
			<p>This is only a test</p>
		</body>
	</html>
	`)

	newFile, err := os.Create("index.html")
	if err != nil {
		log.Fatal("Error creating file", err)
	}
	defer newFile.Close()

	io.Copy(newFile, strings.NewReader(template))

}
