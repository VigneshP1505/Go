package main

import (
	"fmt"
	"os"
)

// interfaces
type Store interface {
	Get(id string) (string, error)
}

type MemoryStore struct{}

func (m MemoryStore) Get(id string) (string, error) {
	return "value", nil
}

// file handling
func readFile() error {
	f, err := os.Open("data.txt")
	if err != nil {
		return fmt.Errorf("Open file:%w", err)
	}
	defer f.Close()
	return nil
}

func ReadStore(s Store) {
	value, err := s.Get("123")
	if err != nil {
		return
	}
	fmt.Println(value)
}

func Fundamentals() {
	nums := []int{10, 20, 30}
	nums = append(nums, 40)
	count := map[string]int{}
	count["go"]++
	fmt.Println(count["go"])

	type User struct {
		ID   int
		Name string
	}

	data := User{ID: 1, Name: "Alex"}
	data.ID = 2
	x := 0
	update(&x)
	fmt.Println(x)
}

func update(x *int) {
	*x = 10
}

func main() {
	Fundamentals()
}
