package test

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/request"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIntegration_LoginUser(t *testing.T) {
	ctx := context.Background()
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "user"
	}

	t.Run("When_user_credentials_are_valid_Returns_success_and_token", func(t *testing.T) {
		randNum := strconv.Itoa(rand.Intn(100000))
		email := "login_test_" + randNum + "@test.com"
		rawPassword := "Test@123" + randNum

		// Calcular o hash MD5 da senha (exatamente como o service do projeto faz)
		hasher := md5.New()
		hasher.Write([]byte(rawPassword))
		hashedPassword := hex.EncodeToString(hasher.Sum(nil))

		// 1. Inserir usuário diretamente no banco com a senha já hacheada
		_, err := db.Collection(collectionName).InsertOne(ctx, bson.M{
			"_id":      primitive.NewObjectID(),
			"email":    email,
			"password": hashedPassword,
			"name":     "Login Test User",
			"age":      24,
		})
		assert.NoError(t, err)

		// 2. Tentar fazer login com as credenciais válidas (senha em texto plano)
		userLoginReq := request.UserLoginRequest{
			Email:    email,
			Password: rawPassword,
		}
		jsonLoginValue, _ := json.Marshal(userLoginReq)
		bodyLogin := bytes.NewBuffer(jsonLoginValue)

		recorderLogin := httptest.NewRecorder()
		ginCtxLogin := CreateTestContext(recorderLogin)
		MakeRequest(ginCtxLogin, gin.Params{}, url.Values{}, "POST", bodyLogin)

		userController.LoginUser(ginCtxLogin)

		// 3. Validar se o login teve sucesso (200 OK) e retornou o token no header
		assert.Equal(t, 200, recorderLogin.Code)
		authHeader := recorderLogin.Result().Header.Get("Authorization")
		assert.NotEmpty(t, authHeader)
	})

	t.Run("When_user_credentials_are_invalid_Returns_forbidden", func(t *testing.T) {
		userLoginReq := request.UserLoginRequest{
			Email:    "nonexistent_login@test.com",
			Password: "WrongPassword@123",
		}
		jsonLoginValue, _ := json.Marshal(userLoginReq)
		bodyLogin := bytes.NewBuffer(jsonLoginValue)

		recorderLogin := httptest.NewRecorder()
		ginCtxLogin := CreateTestContext(recorderLogin)
		MakeRequest(ginCtxLogin, gin.Params{}, url.Values{}, "POST", bodyLogin)

		userController.LoginUser(ginCtxLogin)

		// Validar que retorna erro de autorização/forbidden (403)
		assert.Equal(t, http.StatusForbidden, recorderLogin.Code)
	})
}
