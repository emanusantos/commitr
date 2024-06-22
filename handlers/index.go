package handlers

import (
	cookies "commitr/auth"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func serveTemplate(writer http.ResponseWriter, request *http.Request) {
	var isAuthenticated = false

	_, err := cookies.ReadSigned(request, "token")

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
		password := request.FormValue("password")

		log.Println(password)

		if !strings.HasPrefix(password, "ghp_") {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}

		cookie := http.Cookie{
			Name:  "token",
			Value: password,
		}

		cookies.WriteSigned(writer, cookie)

		http.Redirect(writer, request, "/", http.StatusSeeOther)

	case http.MethodGet:
		token, err := cookies.ReadSigned(request, "token")

		if err != nil {
			log.Println(err)

			serveTemplate(writer, request)

			return
		}

		log.Println(token)
		serveTemplate(writer, request)

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
