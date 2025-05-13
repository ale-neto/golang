package model

type UserDomainInterface interface {
	GetAge() int8
	GetEmail() string
	GetName() string
	GetPassword() string
	GetJSONValue() (string, error)
	EncryptPassword()
	SetID(string)
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
