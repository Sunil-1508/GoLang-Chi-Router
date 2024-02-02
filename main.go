package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	router := chi.NewRouter()
	router.Get("/",homehandler)
	router.Get("/hello", basichandler)
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to Listen to the port", err)
	}
}
func homehandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("home page"))
}
func basichandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World !"))
}
