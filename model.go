package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

//FizzBuzz is the struct representation of the input params
type FizzBuzz struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
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
