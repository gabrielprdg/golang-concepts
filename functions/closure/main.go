package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/olamundo", marcaTempo(oiMundo))
	http.ListenAndServe(":3000", nil)
}

func marcaTempo(fn func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		inicio := time.Now()
		fn(w, r)
		fim := time.Now()
		fmt.Println("Tempo da requisição ", fim.Sub(inicio))
	}
}

func oiMundo(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Ola manos</h1>")
}
