package main

import (
	"fmt"
	"net/url"

	"github.com/LgoLgo/Lgo-brainfuck/brainfuck"
)

func main() {
	str := "Hello World"

	bf := brainfuck.Text2BrainFuck(str)
	fmt.Println(bf)

	// If you want to translate Brainfuck encoding,
	// you need to URL encode it first
	text := brainfuck.BrainFuck2Text(url.QueryEscape(bf))
	fmt.Println(text)
}
