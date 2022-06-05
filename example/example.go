package main

import (
	"fmt"
	"net/url"
)

func main() {
	str := "Hello World"

	bf := brainfuck.Text2BrainFuck(str)
	fmt.Println(bf)

	//If you want to translate Brainfuck encoding,
	//you need to URL encode it first
	text := brainfuck.Text2BrainFuck(url.QueryEscape(bf))
	fmt.Println(text)
}
