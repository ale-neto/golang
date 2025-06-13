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

func TestUserDomainService_UpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repository := mocks.NewMockUserRepository(ctrl)
	service := NewUserDomainService(repository)

	t.Run("when_sending_a_valid_user_and_userId_returns_success", func(t *testing.T) {
		id := bson.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(nil)

		err := service.UpdateUserService(id, userDomain)

		assert.Nil(t, err)
	})

	t.Run("when_sending_a_invalid_user_and_userId_returns_error", func(t *testing.T) {
		id := bson.NewObjectID().Hex()

		userDomain := model.NewUserDomain("test@test.com", "test", "test", 50)
		userDomain.SetID(id)

		repository.EXPECT().UpdateUser(id, userDomain).Return(
			err_rest.NewInternalServerError("error trying to update user"))

		err := service.UpdateUserService(id, userDomain)

		assert.NotNil(t, err)
		assert.EqualValues(t, err.Message, "error trying to update user")
	})
}
