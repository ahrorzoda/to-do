package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

func main() {
	content, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Ошибка чтения файла:", err)
		return
	}
	text := string(content)
	sliceText := strings.Split(text, " ")
	for i := 0; i < len(sliceText); i++ {
		ov, an, ich := sliceText[i], sliceText[i+1], sliceText[i+2]
		if ov[len(ov)-5:len(ov)] == "ов" && an[len(an)-5:len(an)] == "ан" && ich[len(ich)-5:len(ich)] == "ич" {
			log.Println("all ok")
			return
		}
	}
	length := utf8.RuneCountInString(text)
	suffix := subString(text, length-3, length)
	fmt.Println("suffix ==> ", suffix)
	fmt.Println(text)
}
func subString(str string, startPos, endPos int) string {
	if startPos < 0 || startPos >= len(str) || startPos > endPos {
		return ""
	}
	if endPos < 0 || endPos >= len(str) {
		endPos = len(str) - 1
	}
	return strings.Trim(str[startPos:endPos+1], " ")
}
