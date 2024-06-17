package handlers

import (
	"fmt"
	"net/http"
)

func HandleAuthCallback(writer http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case http.MethodGet:
		{
			fmt.Println("Auth callback")
		}

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
