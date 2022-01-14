package model

type JsonClient struct {
	Email    string `json:"login"`
	Password string `json:"password"`
}

type Token struct {
	Token string `json:"token"`
}

type Courses struct {
	Courses []*Course `json:"courses"`
}

type Course struct {
	Course Id `json:"course"`
}

type Id struct {
	Id int `json:"id"`
}
