package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)

	err := http.ListenAndServe(":3333", mux)
	log.Fatal(err)
}

func index(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	switch request.Method {
	case http.MethodGet:
		fmt.Println("GET /")

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}

}
