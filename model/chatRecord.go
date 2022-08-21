package model

type ChatRecord struct {
	Model
	Pid          int
	UserId       int
	UserNickname string
	Content      string
	Ip           string
}
