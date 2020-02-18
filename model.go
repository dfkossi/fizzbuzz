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
	if f.IsValid() != nil {
		return result, f.IsValid()
	}
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
	return result, nil
}

func (f *FizzBuzz) IsValid() error {
	var errs error

	if f.Int1 == 0 || f.Int2 == 0 || f.Limit == 0 {
		errs = errors.New("The value can't be 0")
	}
	if f.Str1 == "" || f.Str2 == "" {
		errs = errors.New("The field is required! Can't be empty")
	}
	if f.Int1 > f.Int2 {
		errs = errors.New("The value must be: Int2 > Int1 ")
	}
	if f.Int2 > f.Limit {
		errs = errors.New("The value must be: Int2 < Limit ")
	}

	return errs
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
