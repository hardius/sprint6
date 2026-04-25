package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func ToConvert(data string) string {
	if len(data) == 0 {

	}
	splitData := strings.Fields(data)
	var isMorse = true

	for _, ch := range splitData[0] {
		if !(ch == '-' || ch == '.') {
			isMorse = false
			break
		}
	}

	if isMorse {
		return morse.ToText(data)
	} else {
		return morse.ToMorse(data)
	}
}
