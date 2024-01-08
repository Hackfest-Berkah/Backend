package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
)

func Profile(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.GET("/profile", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get user profile", user)
	})

	r.GET("/credits", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		userCredits := model.UserCredits{
			KiriBalance: user.KiriBalance,
			KiriPoint:   user.KiriPoint,
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get user credits", userCredits)
	})
}
