package model

type Comment struct {
	Model
	ArticleId    int
	UserId       int
	UserNickName string
	Content      string
	Ip           string
}
