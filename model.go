package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
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

	if f.Int1 == 0 || f.Int2 == 0 {
		errs = errors.New("Invalid integer format!")
	}
	if f.Str1 == "" || f.Str2 == "" {
		errs = errors.New("The field is required!")
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

func checkInputType(f *FizzBuzz) bool {

	allTypes := []string{"int", "int", "int", "string", "string"}

	fmt.Println(strconv.Itoa(f.Int1))

	tstObj := FizzBuzz{f.Int1, f.Int2, f.Limit, f.Str1, f.Str2}
	val := reflect.ValueOf(&tstObj).Elem()
	tab := []string{}

	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Field(i)
		tab = append(tab, fieldType.Type().String())
	}
	result := reflect.DeepEqual(allTypes, tab)
	return result
}

/* func (f *FizzBuzz) checkInputType() bool {

	allTypes := []string{"int", "int", "int", "string", "string"}
	t := new(FizzBuzz)
	tstObj := FizzBuzz{t.Int1, t.Int2, t.Limit, t.Str1, t.Str2}
	val := reflect.ValueOf(&tstObj).Elem()
	tab := []string{}

	for i := 0; i < val.NumField(); i++ {
		fieldType := val.Field(i)
		tab = append(tab, fieldType.Type().String())
	}
	result := reflect.DeepEqual(allTypes, tab)
	return result
} */
