package model

import (
	"crypto/md5"
	"encoding/hex"
)

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	EncryptPassword()
}

type userDomain struct {
	password string
	email    string
	name     string
	age      int8
}

func (u *userDomain) GetEmail() string {
	return u.email
}
func (u *userDomain) GetPassword() string {
	return u.password
}
func (u *userDomain) GetName() string {
	return u.name
}
func (u *userDomain) GetAge() int8 {
	return u.age
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		email,
		password,
		name,
		age,
	}
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
}
