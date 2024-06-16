package handlers

import (
	"html/template"
	"net/http"
)

func HandleHome(writer http.ResponseWriter, request *http.Request) {
	if request.URL.Path != "/" {
		http.NotFound(writer, request)
		return
	}

	switch request.Method {
	case http.MethodGet:
		{
			templ, _ := template.ParseFiles("views/index.html")
			templ.Execute(writer, request)
		}

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
