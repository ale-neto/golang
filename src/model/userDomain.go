package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	GetJSONValue() (string, error)
	EncryptPassword()
}

type userDomain struct {
	Password string
	Email    string
	Name     string
	Age      int8
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling user domain:", err)
		return "", err
	}

	return string(b), nil
}

func (u *userDomain) GetEmail() string {
	return u.Email
}
func (u *userDomain) GetPassword() string {
	return u.Password
}
func (u *userDomain) GetName() string {
	return u.Name
}
func (u *userDomain) GetAge() int8 {
	return u.Age
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
	hash.Write([]byte(u.Password))
	u.Password = hex.EncodeToString(hash.Sum(nil))
}
