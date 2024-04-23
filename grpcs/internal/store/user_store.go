package store

import "context"

type UserStorer interface{
	GetUsers(context.Context)
}