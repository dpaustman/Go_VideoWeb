package model

// request
type UserCredential struct {
	Username string `json:"user_name"`
	Pwd      string `json:"pwd"`
}

// Data model
type VideoInfo struct {
	Id           string
	AuthorId     int
	Name         string
	DisplayCtiem string
}

// 评论实体
type Comments struct {
	Id      string
	VideoId string
	Author  string
	Content string
}
