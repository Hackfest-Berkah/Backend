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

	r.GET("/qr", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User

		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get image", user.QRCode)
	})
}
