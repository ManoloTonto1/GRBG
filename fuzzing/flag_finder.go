package fuzzing

import (
	"encoding/hex"
	"regexp"
	"strings"
)

func ExtractFlag(str string) string {
	pattern := `flag\{[a-zA-Z0-9]+\}`

	re := regexp.MustCompile(pattern)
	matches := re.FindString(str)
	return matches
}
func ExtractFlagFromData(data map[string]string) map[string][]string {
	dataMap := map[string][]string{}
	for specifier, output := range data {
		hexStrings := strings.Split(output, "_")
		for _, hexString := range hexStrings {
			str, err := hex.DecodeString(hexString)
			if err != nil {
				continue
			}
			dataMap[specifier] = append(dataMap[specifier], Reverse(string(str)))
		}

	}
	return dataMap
}
func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
