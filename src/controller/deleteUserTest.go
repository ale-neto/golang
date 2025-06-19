package controller

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/v2/bson"
)

func TestUserControllerInterface_DeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("userId_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "teste",
			},
		}

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := bson.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserService(id).Return(
			err_rest.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := bson.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().DeleteUserService(id).Return(nil)

		MakeRequest(context, param, url.Values{}, "DELETE", nil)
		controller.DeleteUser(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}
