package edstem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strconv"
	"time"

	"github.com/kajikaji0725/Golang_for_Edstem/model"
)

type Client struct {
	client *http.Client
	jar    *cookiejar.Jar
	token  string
}

func NewClient() *Client {
	jar, _ := cookiejar.New(
		nil,
	)
	httpClient := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
		Jar:     jar,
		Timeout: time.Second * 3,
	}
	return &Client{
		client: &httpClient,
		jar:    jar,
	}
}

func (c *Client) login(email, password string) error {
	var err error
	c.token, err = c.GetToken(email, password)
	if err != nil {
		return err
	}
	fmt.Println(c.token)

	res, err := c.NewRequest(http.MethodGet, "https://edstem.org/api/user")
	if err != nil {
		return err
	}
	courses := model.Courses{}
	err = json.Unmarshal(res, &courses)
	if err != nil {
		return err
	}

	fmt.Println(courses.Courses[0].Course.Id)

	res, err = c.NewRequest(http.MethodGet, "https://edstem.org/api/courses/"+strconv.Itoa(courses.Courses[0].Course.Id)+"/lessons")

	if err != nil {
		return err
	}


	lessons := model.Lessons{}
	err = json.Unmarshal(res, &lessons)
	if err != nil {
		return err
	}
	fmt.Println(lessons)
	return nil
}
