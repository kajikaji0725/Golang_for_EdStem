package model

import "time"

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

type Lessons struct {
	Lessons []Lesson `json:"lessons"`
}

type Lesson struct {
	Course_id                                int       `json:"course_id"`
	Created_at                               time.Time `json:"created_at"`
	First_viewed_at                          time.Time `json:"first_viewed_at"`
	Id                                       int       `json:"id"`
	Is_hidden                                bool      `json:"is_hidden"`
	Is_timed                                 bool      `json:"is_timed"`
	Is_unlisted                              bool      `json:"is_unlisted"`
	Last_viewed_slide_id                     int       `json:"last_viewed_slide_id"`
	Last_submissions                         bool      `json:"Last_submissions"`
	Module_id                                int       `json:"module_id"`
	Number                                   int       `json:"number"`
	Openable                                 bool      `json:"openable"`
	Original_id                              int       `json:"original_id"`
	Release_challenge_feedback               bool      `json:"Release_challenge_feedback    "`
	Release_challenge_feedback_while_active  bool      `json:"Release_challenge_feedback_while_active"`
	Release_challenge_solutions              bool      `json:"Release_challenge_solutions"`
	Release_challenge_solutions_while_active bool      `json:"Release_challenge_solutions_while_active"`
	Release_quiz_correctness_only            bool      `json:"Release_quiz_correctness_only"`
	Release_quiz_solutions                   bool      `json:"Release_quiz_solutions"`
	Setting                                  Settings
	Slide_count                              int    `json:"slide_count"`
	State                                    string `json:"state"`
	Status                                   string `json:"status"`
	Timer_duration                           int    `json:"timer_duration"`
	Timer_effective_duration                 int    `json:"timer_effective_duration"`
	Timer_expiration_access                  bool   `json:"timer_expiration_access"`
	Title                                    string `json:"title"`
	Type                                     string `json:"type"`
	User_id                                  int    `json:"user_id"`
}

type Settings struct {
	Quiz_active_status         string `json:"Quiz_active_status"`
	Quiz_mode                  string `json:"Quiz_mode"`
	Quiz_question_number_style string `json:"Quiz_question_number_style"`
}
