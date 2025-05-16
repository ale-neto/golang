package model

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
)

type userDomain struct {
	id       string
	password string
	email    string
	name     string
	age      int8
}

func (u *userDomain) SetID(id string) {
	u.id = id
}

func (u *userDomain) GetJSONValue() (string, error) {
	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("Error marshalling user domain:", err)
		return "", err
	}

	return string(b), nil
}

func (u *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(u.password))
	u.password = hex.EncodeToString(hash.Sum(nil))
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
