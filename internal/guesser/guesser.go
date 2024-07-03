package guesser

import (
	"strings"
	. "json-translator/pkg/storage"
)

func GuessTarget(sourceFilePathStore *StringStore,targetLanguageStore *StringStore) string {
	if sourceFilePathStore.Value == "" {
		return ""
	}
	if targetLanguageStore.Value == "" {
		return ""
	}

	p := sourceFilePathStore.Value
	split := strings.Split(p, ".")
	split[0] += ("-" + targetLanguageStore.Value)
	return strings.Join(split, ".")
}
