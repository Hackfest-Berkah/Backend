package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
	"os"
	"time"
)

func Auth(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.POST("/register", func(c *gin.Context) {
		var userRegister model.UserRegister

		if err := c.BindJSON(&userRegister); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		hashed, err := utils.Hash(userRegister.Password)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		newUser := model.User{
			ID:          uuid.New(),
			Email:       userRegister.Email,
			Password:    hashed,
			Name:        userRegister.Name,
			KiriBalance: 0,
			KiriPoint:   0,
			CreatedAt:   time.Now(),
		}

		if err := db.Create(&newUser).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success Register", newUser)
	})

	r.POST("/login", func(c *gin.Context) {
		var userLogin model.UserLogin

		if err := c.BindJSON(&userLogin); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		var user model.User
		if err := db.Where("email = ?", userLogin.Email).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !utils.CompareHash(userLogin.Password, user.Password) {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Wrong password!")
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"id":   user.ID.String(),
			"type": "user",
			"exp":  time.Now().Add(time.Hour).Unix(),
		})

		strToken, err := token.SignedString([]byte(os.Getenv("TOKEN")))
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Parsed token", gin.H{
			"user":  user,
			"token": strToken,
			"type":  "user",
		})
		return
	})
}
