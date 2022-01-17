package edstem

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/kajikaji0725/Golang_for_Edstem/model"
)

func (c *Client) JsonParse() ([]model.Announcement, error) {
	res, err := c.GetJson(http.MethodGet, "https://edstem.org/api/user")
	if err != nil {
		return nil, err
	}

	courses := model.Courses{}

	err = json.Unmarshal(res, &courses)
	if err != nil {
		return nil, err
	}

	res, err = c.GetJson(http.MethodGet, "https://edstem.org/api/courses/"+strconv.Itoa(courses.Courses[0].Course.Id)+"/lessons")
	if err != nil {
		return nil, err
	}

	lessons := model.Lessons{}
	err = json.Unmarshal(res, &lessons)
	if err != nil {
		return nil, err
	}

	announcement := []model.Announcement{}

	for _, les := range lessons.Lessons {
		if les.Status == "attempted" {
			res, err = c.GetJson(http.MethodGet, "https://edstem.org/api/lessons/slides/"+strconv.Itoa(les.LastViewedSlideId)+"/questions")
			if err != nil {
				return nil, err
			}
			if string(res) == model.ErNotfound || string(res) == model.ErQuize {
				continue
			} else {
				announcement = append(announcement, model.Announcement{Title: les.Title, Deadline: nil})
			}
		}
	}
	return announcement, nil
}
