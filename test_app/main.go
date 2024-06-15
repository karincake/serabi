package main

import (
	"fmt"

	s "github.com/karincake/serabi"
	t "github.com/karincake/serabi/test"
)

func main() {
	smallData := &t.DataSmallComplexNumber{
		Age: 10,
		Exp: 9,
	}
	err := s.Validate(smallData)
	if err != nil {
		fmt.Println(err.Error())
	}

	mediumData := &t.DataMediumComplex{
		Name:    "Sa",
		Address: "2",
		Age:     0,
		Email:   "",
		Phone:   "",
	}
	err = s.Validate(mediumData)
	if err != nil {
		fmt.Println(err.Error())
	}

	mediumData = &t.DataMediumComplex{
		Name:    "Santoso",
		Address: "Jl Localhost",
		Age:     0,
		Email:   "",
		Phone:   "",
	}
	err = s.Validate(mediumData)
	if err != nil {
		fmt.Println(err.Error())
	}
}
