package storage

import (
	"fmt"
)

type StringStore struct {
	Value   string
	Channel chan string
}

func (s *StringStore) SetValue(value string) {
	s.Value = value
	s.Channel <- value
}

func (s *StringStore) Destroy() {
	close(s.Channel)
}

func CreateStringStore() *StringStore {
	stringChannel := make(chan string)
	go func() {
		for s := range stringChannel {
			fmt.Println(s)
		}
	}()
	return &StringStore{
		"",
		stringChannel,
	}
}
