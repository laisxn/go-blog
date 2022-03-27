package model

import "go-gin/jsonTime"

type Model struct {
	Id        int
	CreatedAt jsonTime.Time  `json:"createdAt"`
	UpdatedAt jsonTime.Time  `json:"updatedAt"`
	DeletedAt *jsonTime.Time `sql:"index" json:"deletedAt"`
}
