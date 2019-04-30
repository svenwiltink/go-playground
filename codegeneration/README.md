A little experiment that shows how `go generate` can be used to generate repetitive code.
`go run tools/generateFunctions.go` is run to generate `functions.go`. 

Whenever the tools file is updated `go generate` should be run to update the `functions.go` file.
Under no circumstance should the generated file be changed manually