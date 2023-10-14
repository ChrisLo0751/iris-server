package models

import "time"

type Product struct {
	ID      int64     `json:"id" sql:"id"`
	Name    string    `json:"name sql:"name"`
	Number  int       `json:"number" sql:"number"`
	Image   string    `json:"image" sql:"image"`
	URL     string    `json:"url" sql:"url"`
	Created time.Time `json:"created" sql:"created"`
	Updated time.Time `json:"updated" sql:"updated"`
}
