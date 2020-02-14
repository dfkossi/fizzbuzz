package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCreateFizzBuzz(t *testing.T) {

	payload := []byte(`{"int1":3,"int2":5,"limit":15,"str1":"fizz","str2":"buzz"}`)

	req, _ := http.NewRequest("POST", "/", bytes.NewBuffer(payload))
	response := executeRequest(req)

	expected := `"1,2,fizz,4,buzz,fizz,7,8,fizz,buzz,11,fizz,13,14,fizzbuzz,"`
	checkResponseCode(t, http.StatusCreated, response.Code)
	checkResponseBody(t, response.Body.String(), expected)

}

func executeRequest(request *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()

	a.Router.ServeHTTP(response, request)

	return response
}

func checkResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code: Got %d expect %v\n", expected, actual)
	}
}

func checkResponseBody(t *testing.T, expected, actual string) {
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			actual, expected)
	}
}
