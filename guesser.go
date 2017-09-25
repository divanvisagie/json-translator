package main

import (
	"strings"
)

func GuessTarget(targetLanguageStore *StringStore) string {
	if sourceFilePathStore.value == "" {
		return ""
	}
	if targetLanguageStore.value == "" {
		return ""
	}

	p := sourceFilePathStore.value
	split := strings.Split(p, ".")
	split[0] += ("-" + targetLanguageStore.value)
	return strings.Join(split, ".")
}
