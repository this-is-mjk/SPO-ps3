package models

type User struct {
	UserId   string `json:"userId" binding:"required"`
	Password string `json:"password" binding:"required"`
}
