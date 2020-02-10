package main

import (
    "fmt"
)

func main() {
    
    banana := []string{"green banana", "brown banana", "yellow banana"}
    
    output := make([]*string, len(banana))
    
    for i, b := range banana {
        fmt.Printf("i: %d pointer: %p value: %s\n", i, &b, b)
        output[i] = &b
    }
    
    fmt.Println()
    
    for i, b := range output {
        fmt.Printf("i: %d pointer: %p value: %s\n", i, b, *b)
    }
}
