package models

type User struct {
	Id       uint
	UserName string `bson:"username"`
	Password string `bson:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
