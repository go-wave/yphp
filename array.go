package yphp

import (
	"strings"
)

func Implode(glue string, pieces []string) string {
	return strings.Join(pieces, glue)
}

func Explode(delimiter string, str string) []string {
	return strings.Split(str, delimiter)
}

func In_array(needle string, haystack []string) bool {
	for _, val := range haystack {
		if needle == val {
			return true
		}
	}
	return false
}
