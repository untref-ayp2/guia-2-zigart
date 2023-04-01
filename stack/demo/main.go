package main

import (
	"fmt"
	"guia2/stack"
)

func main() {
	s := stack.NewStack[string](10)
	cadena := "Hola Mundo"
	for _, v := range cadena {
		s.Push(string(v))
	}

	c, err := s.Pop()
	for err == nil {
		fmt.Printf("%s", c)
		c, err = s.Pop()
	}
	fmt.Println()

}
