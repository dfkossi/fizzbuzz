package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Values struct {
	Int1  int
	Int2  int
	Limit int
	Str1  string
	Str2  string
}

func makeValuesFromBytes(bytes []byte) Values {
	values := Values{}
	err := json.Unmarshal(bytes, &values)

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	return values
}
func handlePost(w http.ResponseWriter, req *http.Request) {
	b, err := ioutil.ReadAll(req.Body)
	defer req.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	values := makeValuesFromBytes(b)
	result := FizzBuzz2(values)
	fmt.Printf("Values: %v", result)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(result))

}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlePost)
	log.Fatal(http.ListenAndServe(":8080", router))
}

func FizzBuzz2(val Values) string {
	result := ""
	for i := 1; i <= val.Limit; i++ {
		if i%val.Int1 == 0 {
			result += val.Str1
		}
		if i%val.Int2 == 0 {
			result += val.Str2
		}
		if i%val.Int1 != 0 && i%val.Int2 != 0 {
			result += fmt.Sprint(i)
		}
		result += ","
	}
	return result
}
