package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

type App struct {
	Router *mux.Router
}

func (a *App) Initialize() {
	a.Router = mux.NewRouter()
	a.initializeRoutes()
}
func (a *App) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, a.Router))
}

func (a *App) initializeRoutes() {
	a.Router.HandleFunc("/", a.handlePost).Methods("POST")
	a.Router.Handle("/metrics", promhttp.Handler())

}

func (a *App) handlePost(output http.ResponseWriter, request *http.Request) {
	var fizzbuzz FizzBuzz
	decoder := json.NewDecoder(request.Body)
	if err := decoder.Decode(&fizzbuzz); err != nil {
		fmt.Println("test1")
		respondWithError(output, http.StatusBadRequest, "Invalid request payload")
		return
	}
	defer request.Body.Close()
	result, err := fizzbuzz.createFizzBuzz()
	if err == nil {
		respondWithError(output, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJSON(output, http.StatusCreated, result)
}

func respondWithError(output http.ResponseWriter, code int, message string) {
	respondWithJSON(output, code, map[string]string{"error": message})
}

func respondWithJSON(output http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	output.Header().Set("Content-Type", "application/json")
	output.WriteHeader(code)
	output.Write(response)
}
