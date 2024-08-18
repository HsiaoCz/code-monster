package dao

import (
	"context"

	"github.com/HsiaoCz/code-monster/traceman/types"
)

type SessionCaser interface {
	CreateSessions(context.Context, *types.Sessions) (*types.Sessions, error)
}
