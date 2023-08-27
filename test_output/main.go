package main

import (
	"fmt"

	s "github.com/karincake/serabi"
	t "github.com/karincake/serabi/test"
)

func main() {
	data := t.DataMediumComplex{
		Name:    "123",
		Address: "",
		Age:     0,
		Email:   "",
		Phone:   "",
	}
	err := s.Validate(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
