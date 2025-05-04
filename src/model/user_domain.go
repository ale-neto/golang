package model

import err_rest "github.com/ale-neto/golang/src/config/err"

type UserDomain struct {
	Password string
	Email    string
	Name     string
	Age      int8
}

type UserDomaimInterface interface {
	CreateUser(UserDomain) *err_rest.Err
	UpdateUser(string, UserDomain) *err_rest.Err
	FindUser(string) (*UserDomain, *err_rest.Err)
	DeleteUser(string) *err_rest.Err
}
