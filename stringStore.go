package main

type StringStore struct {
	value   string
	channel chan string
}

func (s StringStore) SetValue(value string) {
	s.value = value
}

func (s StringStore) Destroy() {
	close(s.channel)
}

func CreateStringStore() *StringStore {
	stringChannel := make(chan string)
	return &StringStore{
		"",
		stringChannel,
	}
}
