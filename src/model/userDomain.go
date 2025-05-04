package model

import (
	"crypto/md5"
	"encoding/hex"

	err_rest "github.com/ale-neto/golang/src/config/err"
)

type userDomain struct {
	Password string
	Email    string
	Name     string
	Age      int8
}

type UserDomainInterface interface {
	CreateUser(userDomain) *err_rest.Err
	UpdateUser(string) *err_rest.Err
	FindUser(string) (*userDomain, *err_rest.Err)
	DeleteUser(string) *err_rest.Err
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		Email:    email,
		Password: password,
		Name:     name,
		Age:      age,
	}
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
