package login

import (
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Client struct {
	client *http.Client
	jar    *cookiejar.Jar
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

	req, err := http.NewRequest(http.MethodGet, "https://edstem.org/au/login?redirect=%2Fau%2Fdashboard&auth=1", nil)
	if err != nil {
		return err
	}

	req.Header.Set("email", email)
	req.Header.Set("password", password)

	resp, err := c.client.Do(req)
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

	fmt.Println(string(res))

	return nil
}
