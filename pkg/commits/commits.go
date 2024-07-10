package commits

import (
	"encoding/json"
	"log"
	"net/http"
)

type PushEvent struct {
	Id      string           `json:"id"`
	Type    string           `json:"type"`
	Payload PushEventPayload `json:"payload"`
}

type PushEventPayload struct {
	Head      string                   `json:"head"`
	Commits   []PushEventPayloadCommit `json:"commits"`
	CreatedAt string                   `json:"created_at"`
}

type PushEventPayloadCommit struct {
	Message string `json:"message"`
	Url     string `json:"url"`
}

func filterCommits(commits []PushEvent) (filtered []PushEvent) {
	for _, commit := range commits {
		if commit.Type == "PushEvent" {
			filtered = append(filtered, commit)
		}
	}

	return
}

func fetchCommits(user, key string) []PushEvent {
	client := &http.Client{}
	endpoint := "https://api.github.com/users/" + user + "/events"

	request, err := http.NewRequest("GET", endpoint, nil)
	fallback := make([]PushEvent, 0)

	if err != nil {
		log.Print(err)

		return fallback
	}

	request.Header.Add("User-Agent", user)
	request.Header.Add("Authorization", "Bearer"+" "+key)

	response, err := client.Do(request)

	if err != nil {
		log.Print(err)

		return fallback
	}

	defer response.Body.Close()

	var commits []PushEvent

	err = json.NewDecoder(response.Body).Decode(&commits)

	if err != nil {
		log.Print(err)

		return fallback
	}

	filtered := filterCommits(commits)

	return filtered
}

func Retrieve(user, key string) []PushEvent {
	commits := fetchCommits(user, key)

	return commits
}
