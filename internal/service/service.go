package service

import (
	"errors"
	"strings"
	"unicode"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func MorseConverter(s string) (string, error) {
	var result string

	if isText(s) {
		result = morse.ToMorse(s)
		return result, nil
	}

	if isMorse(s) {
		result = morse.ToText(s)
		return result, nil
	}

	return result, errors.New("faild to convert")
}

func isText(s string) bool {
	isText := true
	for _, r := range s {
		_, ok := morse.DefaultMorse[unicode.ToUpper(r)]
		if !ok {
			isText = false
			break
		}
	}
	return isText
}

func isMorse(s string) bool {
	f := func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	}
	return !strings.ContainsFunc(s, f)
}
