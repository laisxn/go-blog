package model

type Article struct {
	Model
	Title        string
	ShortContent string
	Content      string
	ClickNum     uint
	Tag          string `json:"tag"`
}
