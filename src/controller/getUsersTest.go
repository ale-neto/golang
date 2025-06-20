package controller

import (
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	err_rest "github.com/ale-neto/golang/src/config/err"
	"github.com/ale-neto/golang/src/model"
	"github.com/ale-neto/golang/src/test/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUserControllerInterface_FindUserByEmail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("email_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "TEST_ERROR",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			nil, err_rest.NewInternalServerErr("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		service.EXPECT().FindUserByEmailService("test@test.com").Return(
			model.NewUserDomain(
				"test@test.com",
				"test",
				"test",
				20), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByEmail(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}

func TestUserControllerInterface_FindUserById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mocks.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("id_is_invalid_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userId",
				Value: "teste",
			},
		}

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusBadRequest, recorder.Code)
	})

	t.Run("id_is_valid_service_returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIDService(id).Return(
			nil, err_rest.NewInternalServerError("error test"))

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusInternalServerError, recorder.Code)
	})

	t.Run("email_is_valid_service_returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		context := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		service.EXPECT().FindUserByIDService(id).Return(
			model.NewUserDomain(
				"test@test.com",
				"test",
				"test",
				20), nil)

		MakeRequest(context, param, url.Values{}, "GET", nil)
		controller.FindUserByID(context)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})

}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
