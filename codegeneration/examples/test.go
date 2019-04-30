package main

import (
	"github.com/svenwiltink/go-playground/codegeneration"
	"fmt"
)
func main() {

	slice := []string{"banana", "koekjes"}
	fmt.Println(codegeneration.ContainsString(slice, "banana"))
}
