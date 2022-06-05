# go-brainfuck

Online conversion between `Go` and `Brainfuck`/Go语言与Brainfuck语言在线转换

## Catalog

- [go-brainfuck](#go-brainfuck)
    - [Catalog](#Catalog)
    - [Special statement](#Special-statement)
    - [Install](#Install)
    - [Quick start](#Quick-start)
    - [Other](#Other)

## Special statement

- The scripts released by this repository and any functions involved in them are only used for testing and learning and
  research. Commercial use is prohibited, and illegal use is prohibited. The accuracy, completeness and validity of the
  scripts cannot be guaranteed. Please make your own judgment according to the situation.

- All resource files in this project are prohibited from being reproduced or published in any form by any public account
  or self-media.

- I am not responsible for any scripting problems, including but not limited to any loss or damage caused by any
  scripting errors.

- Do not use any content of this warehouse for commercial or illegal purposes, otherwise you will be responsible for the
  consequences.

- If any unit or individual believes that the script of the project may be suspected of infringing their rights, they
  should promptly notify and provide proof of identity and ownership, and I will delete the relevant script after
  receiving the certification document.

- Anyone viewing this project in any way or using any of its scripts, directly or indirectly, should read this notice
  carefully.I reserve the right to change or supplement this disclaimer at any time.You are deemed to have accepted this
  disclaimer by using and copying the code of any related script or Go project.

**You must completely delete the above from your computer or phone within 24 hours of downloading**

> ***If you use or copy the script of this repository, it is deemed `accepted` This statement, please read it
carefully***

## Install

1. You need a development environment for [Go](https://golang.org/), You can install Gin with the following Go command

```sh
$ go get -u github.com/L2ncE/go-brainfuck
```

2. Import in your code

```go
import "github.com/L2ncE/go-brainfuck"
```

## Quick start

The following code can be found in example.go.

```go
package main

import (
  "fmt"
  "net/url"
)

func main() {
  str := "Hello World"

  bf := Text2BrainFuck(str)
  fmt.Println(bf)

  //If you want to translate Brainfuck encoding,
  //you need to URL encode it first
  text := BrainFuck2Text(url.QueryEscape(bf))
  fmt.Println(text)
}
```


## Other

If you have any other questions, please submit an issue. If you want to participate in the development or have any bugs, please submit a PR

Contact me ![Mail Badge](https://img.shields.io/badge/-llance_24@foxmail.com-c14438?style=flat&logo=Gmail&logoColor=white&link=mailto:llance_24@foxmail.com)
