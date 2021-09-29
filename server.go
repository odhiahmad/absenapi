package main

import (
	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/absenapi/config"
	"github.com/odhiahmad/absenapi/controller"
	"github.com/odhiahmad/absenapi/repository"
	"github.com/odhiahmad/absenapi/service"
	"gorm.io/gorm"
)

var (
	db               *gorm.DB                    = config.SetupDatabaseConnection()
	userRepository   repository.UserRepository   = repository.NewUserRepository(db)


	jwtService    service.JWTService    = service.NewJwtService()
	authService   service.AuthService   = service.NewAuthService(userRepository)
	userService   service.UserService   = service.NewUserService(userRepository)

	authController   controller.AuthController   = controller.NewAuthController(authService, jwtService)
	userController   controller.UserController   = controller.NewUserController(userService, jwtService)

)

func main() {
	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	authRoutes := r.Group("api/auth")
	{
		authRoutes.POST("/login", authController.Login)
	}
	// middleware.AuthorizeJWT(jwtService)

	userRoutes := r.Group("api/user")
	{
		userRoutes.POST("/create", userController.CreateUser)
		userRoutes.PUT("/update", userController.UpdateUser)
	}

	r.Run()
}
