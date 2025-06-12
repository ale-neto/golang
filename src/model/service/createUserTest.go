package service

import (
	"testing"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
	"github.com/ale-neto/golang/src/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserDomainService_CreateUserService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_user_already_exists_returns_error", func(t *testing.T) {
		id := bson.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "Email is already registered in another account")
	})

	t.Run("when_user_is_not_registered_returns_error", func(t *testing.T) {
		id := bson.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(
			nil, err_rest.NewInternalServerErr("error trying to create user"))

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, user)
		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to create user")
	})

	t.Run("when_user_is_not_registered_returns_success", func(t *testing.T) {
		id := bson.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().FindUserByEmail(userDomain.GetEmail()).Return(
			nil, nil)

		repository.EXPECT().CreateUser(userDomain).Return(
			userDomain, nil)

		user, err := service.CreateUserService(userDomain)

		assert.Nil(t, err)
		assert.EqualValues(t, user.GetName(), userDomain.GetName())
		assert.EqualValues(t, user.GetEmail(), userDomain.GetEmail())
		assert.EqualValues(t, user.GetAge(), userDomain.GetAge())
		assert.EqualValues(t, user.GetID(), userDomain.GetID())
		assert.EqualValues(t, user.GetPassword(), userDomain.GetPassword())
	})
}
