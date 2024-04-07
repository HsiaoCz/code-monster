package store

import (
	"context"

	"github.com/HsiaoCz/code-monster/napis/types"
)

type UserStorer interface {
	CreateUser(context.Context, *types.User) (*types.User, error)
}
