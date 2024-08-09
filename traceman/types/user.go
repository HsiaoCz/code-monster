package types

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username         string `json:"username"`
	Password         string `json:"-"`
	Email            string `json:"email"`
	Avatar           string `json:"avatar"`
	Background_image string `json:"background_image"`
}
