package main

import (
	"fmt"
)

type StringStore struct {
	value   string
	channel chan string
}

func (s *StringStore) SetValue(value string) {
	s.value = value
	s.channel <- value
}

func (s *StringStore) Destroy() {
	close(s.channel)
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
