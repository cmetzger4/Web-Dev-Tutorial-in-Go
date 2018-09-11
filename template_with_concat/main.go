package main

import "fmt"

func main() {

	name := "Bruce Banner"

	template := `
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
	`
	fmt.Println(template)

}
