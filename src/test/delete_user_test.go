package test

import (
	"context"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestIntegration_DeleteUser(t *testing.T) {
	ctx := context.Background()
	collectionName := os.Getenv("MONGODB_USER_COLLECTION")
	if collectionName == "" {
		collectionName = "user"
	}

	t.Run("When_user_exists_Deletes_successfully", func(t *testing.T) {
		userID := primitive.NewObjectID()
		email := "delete_test@test.com"

		// 1. Inserir usuário diretamente no banco de teste
		_, err := db.Collection(collectionName).InsertOne(ctx, bson.M{
			"_id":      userID,
			"email":    email,
			"password": "test123@",
			"name":     "Delete Test Name",
			"age":      22,
		})
		assert.NoError(t, err)

		// 2. Configurar requisição para deletar o usuário via controller
		recorder := httptest.NewRecorder()
		ginCtx := CreateTestContext(recorder)

		params := gin.Params{
			{Key: "userId", Value: userID.Hex()},
		}
		MakeRequest(ginCtx, params, url.Values{}, "DELETE", nil)

		userController.DeleteUser(ginCtx)

		// 3. Validar se a rota retornou sucesso (200 OK)
		assert.Equal(t, 200, recorder.Code)

		// 4. Verificar no banco de dados se o usuário realmente foi removido
		var result bson.M
		err = db.Collection(collectionName).FindOne(ctx, bson.M{"_id": userID}).Decode(&result)
		assert.Error(t, err) // Deve retornar erro de documento não encontrado
	})
}
