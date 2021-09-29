package controller

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/odhiahmad/absenapi/dto"
	"github.com/odhiahmad/absenapi/entity"
	"github.com/odhiahmad/absenapi/helper"
	"github.com/odhiahmad/absenapi/service"
)

type AuthController interface {
	Login(ctx *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService  service.JWTService
}

func NewAuthController(authService service.AuthService, jwtService service.JWTService) AuthController {
	return &authController{
		authService: authService,
		jwtService:  jwtService,
	}
}

func (c *authController) Login(ctx *gin.Context) {
	var loginDTO dto.LoginDTO
	errDTO := ctx.ShouldBind(&loginDTO)
	if errDTO != nil {
		response := helper.BuildErrorResponse("Failed to process request", errDTO.Error(), helper.EmptyObj{})
		ctx.AbortWithStatusJSON(http.StatusBadRequest, response)
		return

	}
	authResult := c.authService.VerifyCredential(loginDTO.Username, loginDTO.Password)
	if v, ok := authResult.(entity.User); ok {
		generatedToken := c.jwtService.GenerateToken(strconv.FormatUint(v.Id, 10))
		v.Token = generatedToken
		response := helper.BuildResponse(true, "Ok!", v)
		ctx.JSON(http.StatusOK, response)
		return

	}
	response := helper.BuildErrorResponse("Failed to process request", "Invalid credential", helper.EmptyObj{})
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)

}
