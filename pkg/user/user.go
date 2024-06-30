package user

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type UserInfo struct {
	Login     string `json:"login"`
	AvatarUrl string `json:"avatar_url"`
	Name      string `json:"name"`
}

func GetInfo(key string) (info UserInfo, err error) {
	client := &http.Client{}
	endpoint := "https://api.github.com/user"

	request, err := http.NewRequest("GET", endpoint, nil)

	if err != nil {
		log.Print(err)

		return info, errors.New("Error making get user info request")
	}

	request.Header.Add("User-Agent", "emanusantos")
	request.Header.Add("Authorization", "Bearer"+" "+key)

	response, err := client.Do(request)

	if err != nil {
		log.Print(err)

		return info, errors.New("Error fetching user info")
	}

	defer response.Body.Close()

	var userInfo UserInfo

	err = json.NewDecoder(response.Body).Decode(&userInfo)

	if err != nil {
		log.Print(err)

		return info, errors.New("Error parsing user info to struct")
	}

	return userInfo, nil
}
