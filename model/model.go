package model

type JsonClient struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}
