package main
import (
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/repository"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/model/service"
	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(database *mongo.Database) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	serviceInstance := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(serviceInstance)
}
