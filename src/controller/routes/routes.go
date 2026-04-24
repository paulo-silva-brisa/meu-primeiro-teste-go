package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/paulo-silva-brisa/meu-primeiro-teste-go/src/controller"
)

func InitRouts(r *gin.RouterGroup, userController controller.UserControllerInterface) {
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userId", userController.FindUserByEmail)
	r.POST("/User", userController.CreateUser)
	r.PUT("/User/:userId", userController.UpdateUser)
	r.DELETE("/User/:userId", userController.DeleteUser)

}
