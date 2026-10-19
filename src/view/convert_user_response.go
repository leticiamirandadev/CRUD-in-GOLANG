package view
import (
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller/model/response"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
