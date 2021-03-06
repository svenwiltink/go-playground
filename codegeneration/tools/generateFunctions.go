package main

import (
	"bytes"
	"go/format"
	"io/ioutil"
	"log"
	"strings"
	"text/template"
)

const functionTemplate = `
{{ range $type := .Types }}
func Contains{{ $type | Title }}(slice []{{ $type }}, n {{ $type }}) bool {
	for _, item := range slice {
		if item == n {
			return true
		}
	}
	return false
}
{{ end }}

{{ range $type := .Types }}
func Count{{ $type | Title }}(slice []{{ $type }}, n {{ $type }}) int {
	instances := 0
	for _, item := range slice {
		if item == n {
			instances++
		}
	}
	return instances
}
{{ end }}

{{ range $type := .Types }}
func Delete{{ $type | Title }}(slice []{{ $type }}, i int) []{{ $type }} { 
	return append(slice[:i], slice[i+1:]...)
}
{{ end }}

{{ range $type := .Types }}
func Index{{ $type | Title }}(slice []{{ $type }}, n {{ $type }}) int {
	for index, item := range slice {
		if item == n {
			return index
		}
	}
	return -1
}
{{ end }}

{{ range $type := .Types }}
func LastIndex{{ $type | Title }}(slice []{{ $type }}, n {{ $type }}) int {
	lastIndex := -1
	for index, item := range slice {
		if item == n {
			lastIndex = index
		}
	}
	return lastIndex
}
{{ end }}

{{ range $type := .Types }}
func Replace{{ $type | Title }}(slice []{{ $type }}, old, new {{ $type }}) []{{ $type }} {
	for index, item := range slice {
		if item == old {
			slice[index] = new
			return slice
		}
	}
	return slice
}
{{ end }}

{{ range $type := .Types }}
func ReplaceAll{{ $type | Title }}(slice []{{ $type }}, old, new {{ $type }}) []{{ $type }} {
	for index, item := range slice {
		if item == old {
			slice[index] = new
		}
	}
	return slice
}
{{ end }}
`

const header = `
// CODE GENERATED BY "generateFunctions". DO NOT EDIT MANUALLY
package codegeneration
`

var types = []string{"byte", "bool", "int", "float32", "float64", "string", "int8", "int16"}

func main() {

	funcMap := template.FuncMap{
		"Title": strings.Title,
	}
	templ, err := template.New("sliceFuncs").Funcs(funcMap).Parse(functionTemplate)
	if err != nil {
		panic(err)
	}

	var output bytes.Buffer

	output.Write([]byte(header))

	err = templ.Execute(&output, struct {
		Types []string
	}{Types: types})

	if err != nil {
		panic(err)
	}

	src, err := format.Source(output.Bytes())
	if err != nil {
		log.Println("warning: internal error: invalid Go generated:", err)
		src = output.Bytes()
	}

	err = ioutil.WriteFile("functions.go", src, 0644)
	if err != nil {
		panic(err)
	}
}
