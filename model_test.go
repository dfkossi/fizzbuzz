package main

import (
	"fmt"
	"strconv"
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

func TestCreateFizzBuzzEmptyInput(t *testing.T) {

	fizzBuzz := FizzBuzz{3, 5, 15, "", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if result == "" {
		t.Errorf("\n The error is: %v", err.Error())
	}
}
func TestCreateFizzBuzzInvalidInput(t *testing.T) {

	testValue, errs := strconv.Atoi("ssss")
	fizzBuzz := FizzBuzz{testValue, 5, 15, "fizz", "buzz"}
	result, err := fizzBuzz.createFizzBuzz()
	if errs != nil || result == "" {
		t.Errorf("\n Test return: %v \n The error is: %v", err.Error(), errs.Error())
	}
}

func panicErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
