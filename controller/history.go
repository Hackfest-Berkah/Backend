package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
)

func History(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.GET("/history", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var histories []model.History
		if err := db.Where("user_id = ?", user.ID).Find(&histories).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success", histories)
	})

	r.GET("/history/kiripay", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var histories []model.History
		if err := db.Where("user_id = ?", user.ID).
			Where("type IN ?", []string{"Top Up", "Payment"}).
			Find(&histories).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success", histories)
	})

	r.GET("/history/fleet", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var histories []model.History
		if err := db.Where("user_id = ?", user.ID).
			Where("type IN ?", []string{"Shared Taxi", "City Bus"}).
			Find(&histories).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success", histories)
	})
}
