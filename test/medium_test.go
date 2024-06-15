package test

import (
	"testing"

	d "github.com/karincake/dodol"
	s "github.com/karincake/serabi"
)

func init() {
	s.CacheEnabled = true
}

func TestMediumSimpleSuccess(t *testing.T) {
	data := DataMediumSimple{
		Name:    "Santo Sembodo",
		Address: "Jl Localhost",
		Age:     17,
		Email:   "email@test.com",
		Phone:   "08312312",
	}

	if err := s.Validate(data); err != nil {
		t.Error("failed to parse request: ", err)
	}
}

func TestMediumSimpleFail(t *testing.T) {
	// should gives 5 errors
	data := DataMediumSimple{
		Name:    "1234",
		Address: "Jl Lo",
		Age:     0,
		Email:   "",
		Phone:   "abcd",
	}

	err := s.Validate(data)
	xerr := err.(d.FieldErrors)
	if len(xerr) != 5 {
		t.Error("failed to validate request")
	}
}

func TestMediumSimpleFieldSuccess(t *testing.T) {
	data := DataMediumSimpleField{
		Name:        "santosembodo",
		Password:    "password12345",
		RePassword:  "password12345",
		Age:         15,
		YearsActive: 14,
	}

	if err := s.Validate(data); err != nil {
		t.Error("failed to parse request: \n", err)
	}
}

func TestSmallComplexArrayofStringSuccess(t *testing.T) {
	// should gives 5 errors
	data := DataSmallComplexArrayofString{
		Data: []string{
			"Santo",
			"Sembodo",
			"Beras",
		},
	}

	err := s.Validate(data)
	if err != nil {
		xerr := err.(d.FieldErrors)
		if len(xerr) != 0 {
			t.Error("failed to parse request: \n", err)
		}
	}

	// used tot test cache if it's enabled
	err = s.Validate(data)
	if err != nil {
		xerr := err.(d.FieldErrors)
		if len(xerr) != 0 {
			t.Error("failed to parse request: \n", err)
		}
	}
}

func TestMediumSimpleArrayofStructSuccess(t *testing.T) {
	// should gives 5 errors
	data := DataMediumSimpleArrayofStruct{
		Data: []DataMediumSimple{
			{
				Name:    "Santo Sembodo Beras",
				Address: "Jl Localhost 2023",
				Age:     19,
				Email:   "test@example.com",
			},
			{
				Name:    "Santo Sembodo Beras",
				Address: "Jl Localhost 2023",
				Age:     19,
				Email:   "test@example.com",
			},
		},
	}

	err := s.Validate(data)
	if err != nil {
		xerr := err.(d.FieldErrors)
		if len(xerr) != 0 {
			t.Error("failed to parse request: \n", err)
		}
	}

	// used tot test cache if it's enabled
	if err != nil {
		xerr := err.(d.FieldErrors)
		if len(xerr) != 0 {
			t.Error("failed to parse request: \n", err)
		}
	}
}

func TestMediumSimpleFieldFail(t *testing.T) {
	// should gives 5 errors
	data := DataMediumSimpleField{
		Name:        "santo sembodo",
		Password:    "passw",
		RePassword:  "password1234",
		Age:         9,
		YearsActive: 12,
	}

	err := s.Validate(data)
	xerr := err.(d.FieldErrors)
	if len(xerr) != 5 {
		t.Error("failed to parse request: \n", err)
	}
}
