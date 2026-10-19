package test

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/request"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func TestIntegration_CreateUser(t *testing.T) {
	ctx := context.Background()
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "user"
	}

	t.Run("When_user_already_exists_Returns_bad_request", func(t *testing.T) {
		email := "existing_" + strconv.Itoa(rand.Intn(100000)) + "@test.com"

		// 1. Inserir usuário prévio no banco
		_, err := db.Collection(collectionName).InsertOne(ctx, bson.M{
			"email":    email,
			"password": "test123@",
			"name":     "Existing User",
			"age":      25,
		})
		assert.NoError(t, err)

		// 2. Tentar criar o usuário com o mesmo e-mail via Controller
		userReq := request.UserRequest{
			Email:    email,
			Password: "test123@",
			Name:     "New Attempt",
			Age:      25,
		}
		jsonValue, _ := json.Marshal(userReq)
		body := bytes.NewBuffer(jsonValue)

		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)
		MakeRequest(ginCtx, gin.Params{}, url.Values{}, "POST", body)

		userController.CreateUser(ginCtx)

		// 3. Validar se retornou Bad Request (400)
		assert.Equal(t, 400, recorder.Code)
	})

	t.Run("When_user_does_not_exist_Creates_successfully", func(t *testing.T) {
		email := "new_" + strconv.Itoa(rand.Intn(100000)) + "@test.com"

		userReq := request.UserRequest{
			Email:    email,
			Password: "test123@",
			Name:     "Brand New User",
			Age:      28,
		}
		jsonValue, _ := json.Marshal(userReq)
		body := bytes.NewBuffer(jsonValue)

		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)
		MakeRequest(ginCtx, gin.Params{}, url.Values{}, "POST", body)

		userController.CreateUser(ginCtx)

		// 1. Validar se retornou Sucesso (200 OK)
		assert.Equal(t, 200, recorder.Code)

		// 2. Verificar diretamente no banco de dados se o usuário foi persistido
		var result bson.M
		err := db.Collection(collectionName).FindOne(ctx, bson.M{"email": email}).Decode(&result)
		assert.NoError(t, err)
		assert.NotNil(t, result)
		assert.Equal(t, email, result["email"])
	})
}
