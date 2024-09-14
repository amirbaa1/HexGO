package model

type RequestBook struct {
	title  string `json:"title"`
	author Author `json:"author"`
}
