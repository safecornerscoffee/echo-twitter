package model

type (
	User struct {
		ID        string   `json:"id"`
		Email     string   `json:"email"`
		Password  string   `json:"password"`
		Token     string   `json:"token"`
		Followers []string `json:"followers"`
	}
)
