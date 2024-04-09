package service

import (
	"context"

	"github.com/HsiaoCz/code-monster/lost/internal/user/pb"
	"github.com/HsiaoCz/code-monster/lost/internal/user/store"
	"github.com/HsiaoCz/code-monster/lost/internal/user/types"
)

type UserService struct {
	pb.UnimplementedLostServer
	store *store.Store
}

func NewUserService(store *store.Store) *UserService {
	return &UserService{
		store: store,
	}
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

func (u *UserService) UserLogin(ctx context.Context, in *pb.UserLoginRequest) (*pb.UserLoginResponse, error) {
	return nil, nil
}

func (u *UserService) GetUserByID(ctx context.Context, in *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	return nil, nil
}
