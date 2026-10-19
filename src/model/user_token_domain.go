package model

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/logger"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *userDomain) GenerateToken() (string, *rest_err.RestErr) {
	secret := os.Getenv("JWT_SECRET")

	claims := jwt.MapClaims{
		"id":    ud.id,
		"email": ud.email,
		"name":  ud.name,
		"age":   ud.age,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	
}
token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		logger.Error("Error trying to generate JWT token", err, zap.String("journey", "generateToken"))
		return "", rest_err.NewInternalServerError("Error trying to generate JWT token")
	}

	return tokenString, nil
}
func VerifyTokenMiddleware(c *gin.Context) {
	secret := os.Getenv("JWT_SECRET")
	tokenValue := c.GetHeader("Authorization")

	tokenTrimmed := removeBearerPrefix(tokenValue)

	token, err := jwt.Parse(tokenTrimmed, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, rest_err.NewBadRequest(fmt.Sprintf("Unexpected signing method: %v", token.Header["alg"]))
		}
		return []byte(secret), nil
	})
      if err != nil {
		errRest := rest_err.NewUnauthorized("Invalid token")
		c.AbortWithStatusJSON(errRest.Code, errRest)
		return
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		errRest := rest_err.NewUnauthorized("Invalid token")
		c.AbortWithStatusJSON(errRest.Code, errRest)
		return
	}

	userId, _ := claims["id"].(string)
	email, _ := claims["email"].(string)
	name, _ := claims["name"].(string)
	ageFloat, _ := claims["age"].(float64)
      logger.Info(
		"User authenticated successfully via token",
		zap.String("journey", "verifyTokenMiddleware"),
		zap.String("userId", userId),
		zap.String("email", email),
	)

	// Opcional: injetar dados no contexto do Gin se precisar usar no controller depois
	c.Set("user_id", userId)
	c.Set("user_email", email)
	c.Set("user_name", name)
	c.Set("user_age", int8(ageFloat))
}
func removeBearerPrefix(token string) string {
	if strings.HasPrefix(token, "Bearer ") {
		return strings.TrimPrefix(token, "Bearer ")
	}
	return token
}
