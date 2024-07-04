package state

import (
	. "json-translator/internal/jsonstore"
	. "json-translator/pkg/storage"
)

type ApplicationStateStore struct {
	DestinationFilePath *StringStore
	SourceFilePath      *StringStore
	TargetJSONKey       *StringStore
	TargetLanguage      *StringStore
	TranslatedJSONFile  *StringStore
	JSONFileStore       *JSONFileStore
}
