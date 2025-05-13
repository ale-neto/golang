package model

import (
	"encoding/json"
	"fmt"
)

type userDomain struct {
	ID       string
	Password string
	Email    string
	Name     string
	Age      int8
}

func (u *userDomain) SetID(id string) {
	u.ID = id
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
