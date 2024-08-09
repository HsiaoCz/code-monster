package dao

import "context"

type UserCaser interface{
	CreateUser(context.Context)
}

type UserDao struct{}