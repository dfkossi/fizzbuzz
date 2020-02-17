package main

import (
	"testing"
)

func TestCreateFizzBuzz(t *testing.T) {
	fizzBuzz := FizzBuzz{3, 5, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()

	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,"
	if result != expected || result == "" {
		t.Errorf(err.Error())
	}
}
