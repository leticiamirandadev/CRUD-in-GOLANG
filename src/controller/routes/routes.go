package routes
import (
	"github.com/gin-gonic/gin"
	"github.com/leticiamirandadev/CRUD-in-GOLANG/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface,
) {
     r.POST("/user", userController.CreateUser)
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.PUT("/user/:userId", userController.UpdateUser)
	r.DELETE("/user/:userId", userController.DeleteUser)
	r.POST("/login", userController.LoginUser)
}
