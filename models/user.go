package models

type User struct {
	Id       uint   `json:"id" binding:"required"`
	Name     string `json:"username"`
	Password string `json:"omit"`
}
