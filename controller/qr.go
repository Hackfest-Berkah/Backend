package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
)

func QR(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.GET("/qr/:userID", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User

		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get image", user.QRCode)
	})

	r.POST("/qr/scan/:amount", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		amountStr := c.Param("amount")
		amount := utils.StringToFloat(amountStr, c)

		var user model.User

		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		user.KiriBalance += amount
		user.KiriPoint += amount

		if err := db.Save(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success scan QR", nil)
	})
}
