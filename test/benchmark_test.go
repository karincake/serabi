package test

import (
	"testing"

	s "github.com/karincake/serabi"
)

func BenchmarkSmallStringSuccess(b *testing.B) {
	instance := DataSmallSimpleString{
		Name: "Santo Sembodo",
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}

func BenchmarkSmallNumberSuccess(b *testing.B) {
	instance := DataSmallSimpleNumber{
		Age: 30,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}

func BenchmarkSmallBoolSuccess(b *testing.B) {
	instance := DataSmallSimpleBoolean{
		Married: true,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}

func BenchmarkSmallStringFail(b *testing.B) {
	instance := DataSmallSimpleString{
		Name: "1234",
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}

func BenchmarkSmallNumberFail(b *testing.B) {
	instance := DataSmallSimpleNumber{
		Age: 5,
	}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}

func BenchmarkSmallBoolFail(b *testing.B) {
	instance := DataSmallSimpleBoolean{}
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		_ = s.Validate(&instance)
	}
}
