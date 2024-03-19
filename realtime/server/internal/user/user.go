package user

import "context"

type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type Repository interface {
	CreateUser(context.Context, *User) (*User, error)
}

type Service interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
	UpdateUser(context.Context, *CreateUserReq) (*CreateUserResp, error)
}

type CreateUserReq struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
type CreateUserResp struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}
