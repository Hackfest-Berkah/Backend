package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
	"time"
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

	r.GET("/qr/helper", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		utils.HttpRespSuccess(c, http.StatusOK, "id", ID)
	})

	r.POST("/qr/:user_id/:amount", middleware.Authorization(), func(c *gin.Context) {
		userIDStr := c.Param("user_id")
		userID, err := uuid.Parse(userIDStr)
		if err != nil {
			fmt.Println("Error parsing UUID:", err)
			return
		}

		amountStr := c.Param("amount")
		amount := utils.StringToFloat(amountStr, c)

		var status model.Status
		if err := db.Where("user_id = ?", userID).First(&status).Error; err != nil {
			//	not found and not there
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		if status.Status == false {
			status.Status = true

			if err := db.Where("user_id = ?", userID).Save(&status).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "success", "User success tapped in!")
			return
		} else {
			var user model.User

			if err := db.Where("id = ?", userID).First(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
				return
			}

			if user.KiriBalance < amount {
				utils.HttpRespFailed(c, http.StatusNotAcceptable, "Insufficient Balance")
				return
			}

			user.KiriBalance -= amount

			if err := db.Save(&user).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			status.Status = false
			status.End = time.Now()
			status.UpdatedAt = time.Now()

			if err := db.Where("user_id = ?", userID).Save(&status).Error; err != nil {
				utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
				return
			}

			utils.HttpRespSuccess(c, http.StatusOK, "Success Tapped Out", nil)
			return
		}
	})
}
