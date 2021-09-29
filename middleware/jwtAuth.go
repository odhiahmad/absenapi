package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/odhiahmad/absenapi/helper"
	"github.com/odhiahmad/absenapi/service"
)

func AuthorizeJWT(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response := helper.BuildErrorResponse("Failed to process request", "No token found", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		token, err := jwtService.ValidateToken(authHeader)
		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			log.Println("Claim[userId] ", claims["userId"])
			log.Println("Claim[issuer] :", claims["issuer"])
		} else {
			log.Println(err)
			response := helper.BuildErrorResponse("Token is not valid", err.Error(), nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
		}
	}
}
