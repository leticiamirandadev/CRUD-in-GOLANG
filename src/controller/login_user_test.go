package controller

import (
	"bytes"
	"encoding/json"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/request"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/test/mock"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/mock/gomock"
)

func TestUserController_LoginUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	service := mock.NewMockUserDomainService(ctrl)
	controller := NewUserControllerInterface(service)

	t.Run("When_request_is_invalid_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		loginRequest := request.UserLoginRequest{
			Email:    "invalid_email",
			Password: "123",
		}
		jsonValue, _ := json.Marshal(loginRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		controller.LoginUser(ctx)

		assert.Equal(t, 400, recorder.Code)
	})

	t.Run("When_service_returns_error_Returns_error", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		loginRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "test123@",
		}
		jsonValue, _ := json.Marshal(loginRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		service.EXPECT().LoginUserServices(gomock.Any()).Return(nil, "", rest_err.NewInternalServerError("error trying to login user"))

		controller.LoginUser(ctx)

		assert.Equal(t, 500, recorder.Code)
	})

	t.Run("When_user_is_logged_successfully_Returns_success", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := TestContext(recorder)

		loginRequest := request.UserLoginRequest{
			Email:    "test@test.com",
			Password: "test123@",
		}
		jsonValue, _ := json.Marshal(loginRequest)
		body := bytes.NewBuffer(jsonValue)

		MakeRequest(ctx, gin.Params{}, url.Values{}, "POST", body)

		id := primitive.NewObjectID().Hex()
		userDomain := model.NewUserDomain("test@test.com", "test123@", "Test Name", 20)
		userDomain.SetID(id)

		token := "mock_jwt_token"

		service.EXPECT().LoginUserServices(gomock.Any()).Return(userDomain, token, nil)

		controller.LoginUser(ctx)

		assert.Equal(t, 200, recorder.Code)
		assert.Equal(t, token, recorder.Header().Get("Authorization"))
	})
}
