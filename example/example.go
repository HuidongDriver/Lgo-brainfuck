package main

import (
	"fmt"
	"github.com/L2ncE/go-brainfuck/brainfuck"
	"net/url"
)

func main() {
	str := "Hello World"

	bf := brainfuck.Text2BrainFuck(str)
	fmt.Println(bf)

	//If you want to translate Brainfuck encoding,
	//you need to URL encode it first
	text := brainfuck.BrainFuck2Text(url.QueryEscape(bf))
	fmt.Println(text)
}
