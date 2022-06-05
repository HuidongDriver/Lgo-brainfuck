package main

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"
)

func Text2BrainFuck(text string) string {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`input=%s&do=Text+to+Brainfuck`, text))
	bodyText, err := Set(client, data)
	if err != nil {
		log.Fatal(err)
	}
	ret := regexp.MustCompile(`10">(?s:(.*?))</te`)
	res := ret.FindString(string(bodyText))
	trimStr := strings.Trim(res, "10\">")
	trimStr = trimStr[:len(trimStr)-4]
	return trimStr
}

func BrainFuck2Text(bf string) string {
	client := &http.Client{}
	var data = strings.NewReader(fmt.Sprintf(`input=%s&do=Brainfuck+to+Text`, bf))
	bodyText, err := Set(client, data)
	if err != nil {
		log.Fatal(err)
	}
	ret := regexp.MustCompile(`10">(?s:(.*?))</te`)
	res := ret.FindString(string(bodyText))
	trimStr := strings.Trim(res, "10\">")
	trimStr = trimStr[:len(trimStr)-4]
	return trimStr
}
