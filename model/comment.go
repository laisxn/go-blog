package model

type Comment struct {
	Model
	Pid          int
	ArticleId    int
	UserId       int
	UserNickname string
	Content      string
	Ip           string
}
