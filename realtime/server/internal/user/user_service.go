package user

import (
	"context"
	"strconv"
	"time"

	"github.com/HsiaoCz/code-monster/realtime/server/utils"
)

type service struct {
	Repository
	timeout time.Duration
}

func NewService(respitory Repository) Service {
	return &service{
		respitory,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// TODO hash password
	hashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Username: req.Username,
		Password: hashPassword,
		Email:    req.Email,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	resp := &CreateUserResp{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}
	return resp, nil
}

func (s *service) UpdateUser(c context.Context, req *CreateUserReq) (*CreateUserResp, error) {
	return nil, nil
}
