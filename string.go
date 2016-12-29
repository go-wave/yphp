package yphp

import (
	"strings"
	"unicode"
)

// 单词首字母大写
func Ucwords(str string) string {
	return strings.Title(str)
}

// 首字母小写
func Lcfirst(str string) string {
	var n []rune
	for i, ch := range str {
		if i == 0 {
			n = append(n, unicode.ToLower(ch))
		} else {
			n = append(n, ch)
		}
	}
	return string(n)
}
