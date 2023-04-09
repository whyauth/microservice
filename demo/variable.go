package main

import "fmt"

func variableZeroValue() {
	var a int
	var b string
	fmt.Printf("ge shi hua %d %q\n", a, b)
}

func variableInitValue() {
	var a, c int = 3, 9
	var b string = "你好"
	fmt.Println(a, b, c)
}

func main() {
	variableZeroValue()
}
