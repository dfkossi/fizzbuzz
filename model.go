package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type FizzBuzz struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func (f *FizzBuzz) makeFizzBuzzFromBytes(bytes []byte) FizzBuzz {
	fizzBuzz := FizzBuzz{}
	err := json.Unmarshal(bytes, &fizzBuzz)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return fizzBuzz
}
func (f *FizzBuzz) makeBytesFromFizzBuzz(fizzBuzz FizzBuzz) []byte {
	bytes, err := json.Marshal(fizzBuzz)
	panic(err)
	return bytes
}

func (f *FizzBuzz) createFizzBuzz() (string, error) {
	result := ""
	for i := 1; i <= f.Limit; i++ {
		if i%f.Int1 == 0 {
			result += f.Str1
		}
		if i%f.Int2 == 0 {
			result += f.Str2
		}
		if i%f.Int1 != 0 && i%f.Int2 != 0 {
			result += fmt.Sprint(i)
		}
		result += ","
	}
	if result == "" {
		return result, errors.New("Error")
	}
	return result, errors.New(" ")
}

var (
	opsProcessed = promauto.NewCounter(prometheus.CounterOpts{
		Name: "myapp_processed_ops_total",
		Help: "The total number of processed events",
	})
)

func recordMetrics() {
	go func() {
		for {
			opsProcessed.Inc()
			time.Sleep(2 * time.Second)
		}
	}()
}
