package store

import (
	"context"

	"github.com/HsiaoCz/code-monster/grpcs/internal/types"
)

type UserStorer interface {
	GetUsers(context.Context, *types.User) (*types.User, error)
}
