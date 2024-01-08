package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"hackfest/utils"
	"net/http"
	"os"
)

func Authorization() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header.Get("Authorization")
		header = header[len("Bearer "):]

		token, err := jwt.Parse(header, func(t *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("TOKEN")), nil
		})
		if err != nil {
			utils.HttpRespFailed(c, http.StatusBadRequest, err.Error())
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			idStr := claims["id"].(string)
			id, err := uuid.Parse(idStr)
			if err != nil {
				utils.HttpRespFailed(c, http.StatusBadRequest, err.Error())
				c.Abort()
				return
			}

			typeStr := claims["type"].(string)

			c.Set("id", id)
			c.Set("type", typeStr)

			c.Next()
			return
		} else {
			utils.HttpRespFailed(c, http.StatusForbidden, err.Error())
			c.Abort()
			return
		}

	}
}
