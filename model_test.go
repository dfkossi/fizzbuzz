package main

import (
	"testing"
)

func TestCreateFizzBuzz(t *testing.T) {
	fizzBuzz := FizzBuzz{3, 5, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()

	expected := "1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,"
	if result != expected || err != nil {
		t.Errorf(err.Error())
	}
}

func TestStringNotEmpty(t *testing.T) {

	fizzBuzz := FizzBuzz{3, 5, 15, "", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if err != nil && result == "" {
		t.Errorf("\n The error is: %v", err.Error())
	}
}

func TestCorrectIntegerInput(t *testing.T) {

	fizzBuzz := FizzBuzz{5, 3, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if err != nil && result == "" {
		t.Errorf("\n The error is: %v", err.Error())
	}
}

func TestCorrectLimit(t *testing.T) {

	fizzBuzz := FizzBuzz{3, 50, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if err != nil && result == "" {
		t.Errorf("\n The error is: %v", err.Error())
	}
}
func TestCreateFizzBuzzInvalidInput(t *testing.T) {

	fizzBuzz := FizzBuzz{0, 5, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if err != nil || result == "" {
		t.Errorf("\n Test return: %v", err.Error())
	}
}
