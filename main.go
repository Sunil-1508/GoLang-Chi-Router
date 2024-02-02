package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := &http.Server{
		Addr:    ":3000",
		Handler: http.HandlerFunc(basichandler),
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("failed to Listen to the port", err)
	}
}

func basichandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world !!!!!"))
}
