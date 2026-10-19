package test

import (
	"bytes"
	"context"
	"encoding/json"
	"math/rand"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/request"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIntegration_UpdateUser(t *testing.T) {
	ctx := context.Background()
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "user"
	}

	t.Run("When_user_exists_Updates_successfully", func(t *testing.T) {
		userID := primitive.NewObjectID()
		email := "update_test@test.com"

		// 1. Inserir usuário inicial no banco
		_, err := db.Collection(collectionName).InsertOne(ctx, bson.M{
			"_id":      userID,
			"email":    email,
			"password": "test123@",
			"name":     "Old Name",
			"age":      20,
		})
		assert.NoError(t, err)

		// 2. Definir dados atualizados usando ponteiros
		newName := "New Updated Name"
		newAge := int8(25 + rand.Intn(50))

		userUpdateReq := request.UserUpdateRequest{
			Name: &newName,
			Age:  &newAge,
		}
		jsonValue, _ := json.Marshal(userUpdateReq)
		body := bytes.NewBuffer(jsonValue)

		// 3. Configurar requisição PUT via Controller
		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)

		params := gin.Params{
			{Key: "userId", Value: userID.Hex()},
		}
		MakeRequest(ginCtx, params, url.Values{}, "PUT", body)

		userController.UpdateUser(ginCtx)

		// 4. Validar se retornou sucesso (200 OK)
		assert.Equal(t, 200, recorder.Code)

		// 5. Verificar diretamente no banco se os dados foram atualizados
		var result bson.M
		err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
		assert.NoError(t, err)
		assert.Equal(t, newName, result["name"])
		assert.Equal(t, int32(newAge), result["age"])
	})
}
