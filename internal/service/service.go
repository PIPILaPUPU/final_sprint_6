package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Convertation(input string) string {
	input = strings.TrimSpace(input)

	if strings.ContainsAny(input, ".-") {
		return morse.ToText(input)
	}

	return morse.ToMorse(input)
}
