package handlers

import (
	cookies "commitr/auth"
	"commitr/pkg/commits"
	"commitr/pkg/user"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func serveTemplate(writer http.ResponseWriter, request *http.Request) {
	var isAuthenticated = true

	token, err := cookies.ReadSigned(request, "token")

	if err != nil {
		isAuthenticated = false

		templ, _ := template.ParseFiles("views/index.html")
		templ.Execute(writer, map[string]interface{}{
			"IsAuthenticated": false,
		})

		return
	}

	userCookie, err := cookies.ReadSigned(request, "user")

	if err != nil {
		log.Println(err)
	}

	var user user.UserInfo
	err = json.Unmarshal([]byte(userCookie), &user)
	commits := commits.Retrieve(user.Name, token)

	log.Print(commits)

	data := map[string]interface{}{
		"IsAuthenticated": isAuthenticated,
		"User":            user,
		"Commits":         commits,
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

		user, err := user.GetInfo(password)

		if err != nil {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}

		output, err := json.Marshal(user)

		log.Print(string(output))

		if err != nil {
			http.Redirect(writer, request, "/", http.StatusSeeOther)
			return
		}

		cookies.WriteSigned(writer, http.Cookie{
			Name:  "token",
			Value: password,
		})

		cookies.WriteSigned(writer, http.Cookie{
			Name:  "user",
			Value: string(output),
		})

		http.Redirect(writer, request, "/", http.StatusSeeOther)

	case http.MethodGet:
		serveTemplate(writer, request)

	default:
		http.Error(writer, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
