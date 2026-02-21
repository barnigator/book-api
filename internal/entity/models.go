package entity

type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title,omitempty"`
	Author string `json:"author,omitempty"`
	Year   int    `json:"year,omitempty"`
}

var ID int
