package data

type Repository struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Organization string `json:"organization"`
	UserId       string `json:"user_id"`
}
