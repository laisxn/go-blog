package model

import "go-gin/localTime"

type Model struct {
	Id        int
	CreatedAt localTime.Time  `json:"created_at"`
	UpdatedAt localTime.Time  `json:"updated_at"`
	DeletedAt *localTime.Time `sql:"index" json:"deleted_at"`
}
