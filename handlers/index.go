package handlers

import (
	cookies "commitr/auth"
	"fmt"
	"html/template"
	"net/http"
)

func serveTemplate(writer http.ResponseWriter, request *http.Request) {
	var isAuthenticated = false

	_, err := cookies.ReadSigned(request, "session")

	if err == nil {
		isAuthenticated = true
	}

	data := map[string]interface{}{
		"IsAuthenticated": isAuthenticated,
	}

	templ, _ := template.ParseFiles("views/index.html")
	templ.Execute(writer, data)
}

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	switch request.Method {
	case http.MethodPost:
		fmt.Println("TRIGGERED LOGIN REQUEST")
		serveTemplate(writer, request)

	case http.MethodGet:
		serveTemplate(writer, request)

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
