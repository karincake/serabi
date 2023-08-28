package main

import (
	"fmt"

	s "github.com/karincake/serabi"
	t "github.com/karincake/serabi/test"
)

func main() {
	// s.CacheEnabled = true
	data := &t.DataMediumComplex{
		Name:    "Sa",
		Address: "2",
		Age:     0,
		Email:   "",
		Phone:   "",
	}
	err := s.Validate(data)
	if err != nil {
		fmt.Println(err.Error())
	}

	data = &t.DataMediumComplex{
		Name:    "Santoso",
		Address: "Jl Localhost",
		Age:     0,
		Email:   "",
		Phone:   "",
	}
	err = s.Validate(data)
	if err != nil {
		fmt.Println(err.Error())
	}
}
