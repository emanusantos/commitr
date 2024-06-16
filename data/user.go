package data

type User struct {
	Id            string `json:"id"`
	Username      string `json:"username"`
	Token         string `json:"token"`
	Refresh_token string `json:"refresh_token"`
}
