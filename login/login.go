package login

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
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
	jsonclient := model.JsonClient{Email: email, Password: password}

	jsonjson, _ := json.Marshal(jsonclient)

	resp, err := c.client.Post("https://edstem.org/api/token", "application/json", bytes.NewBuffer(jsonjson))
	if err != nil {
		return err
	}
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("Response status was %d(except %d)", resp.StatusCode, http.StatusOK)
	}

	defer resp.Body.Close()

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	c.token = string(res)
	fmt.Println(c.token)

	return nil
}
