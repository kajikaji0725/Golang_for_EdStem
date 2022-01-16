package edstem

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/kajikaji0725/Golang_for_Edstem/model"
)

func (c *Client) JsonParse() error {
	res, err := c.NewRequest(http.MethodGet, "https://edstem.org/api/user")
	if err != nil {
		return err
	}
	courses := model.Courses{}
	err = json.Unmarshal(res, &courses)
	if err != nil {
		return err
	}

	res, err = c.NewRequest(http.MethodGet, "https://edstem.org/api/courses/"+strconv.Itoa(courses.Courses[0].Course.Id)+"/lessons")

	if err != nil {
		return err
	}

	lessons := model.Lessons{}
	err = json.Unmarshal(res, &lessons)
	if err != nil {
		return err
	}

	for _, les := range lessons.Lessons {
		if strings.Compare(les.Status, "attempted") == 0 {
			res, err = c.NewRequest(http.MethodGet, "https://edstem.org/api/lessons/slides/"+strconv.Itoa(les.Last_viewed_slide_id)+"/questions")
			if err != nil {
				return err
			}
			if strings.EqualFold(string(res), model.ErNotfound) || strings.EqualFold(string(res), model.ErQuize) {
				continue
			} else {
				fmt.Println(les.Title)
			}
		}
	}
	return nil
}
