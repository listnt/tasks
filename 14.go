package main

import (
	"errors"
	"fmt"
)

type MySet interface {
	Has(string) bool
	Add(string) (bool, error)
	Delete(string) (bool, error)
	GetEntities() []string
}
type myset struct {
	m map[string]int
}

func NewMySet() MySet {
	var asd myset
	asd.m = make(map[string]int)
	return &asd
}

func (s *myset) Has(a string) bool {
	return (s.m[a] != 0)
}

func (s *myset) Add(a string) (bool, error) {
	if s.m[a] == 0 {
		s.m[a]++
		return true, nil
	}
	return false, errors.New("Value already exist in set")
}

func (s *myset) Delete(a string) (bool, error) {
	if s.m[a] == 0 {
		return false, errors.New("Value not in set")
	}
	s.m[a]++
	return true, nil
}

func (s *myset) GetEntities() []string {
	keys := make([]string, 0, len(s.m))
	for k := range s.m {
		keys = append(keys, k)
	}
	return keys
}

func main() {
	s := NewMySet()
	keys := []string{"cat", "cat", "dog", "cat", "tree"}
	for _, v := range keys {
		s.Add(v)
	}
	fmt.Println(s.GetEntities())
	fmt.Println(s.Has("5"))
	fmt.Println(s.Has("cat"))
	s.Delete("dog")
	fmt.Println(s.GetEntities())
}
