package service

import (
	"context"

	"github.com/HsiaoCz/code-monster/lost/internal/user/pb"
	"github.com/HsiaoCz/code-monster/napis/store"
	"github.com/HsiaoCz/code-monster/napis/types"
)

type UserService struct {
	pb.UnimplementedLostServer
	store store.Store
}

func (u *UserService) CreateUser(ctx context.Context, in *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := u.store.UserStore.CreateUser(ctx, &types.User{
		Firstname: in.Firstname,
		Lastname:  in.Lastname,
		Password:  in.Password,
		Email:     in.Email,
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		ID:        user.ID,
		Email:     user.Email,
	}, nil
}
