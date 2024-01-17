package controller

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
	"time"
)

func Fleet(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.GET("/fleet/all", middleware.Authorization(), func(c *gin.Context) {
		var fleets []model.Fleet

		if err := db.Find(&fleets).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get all fleets", fleets)
	})

	r.GET("/fleet/:id", middleware.Authorization(), func(c *gin.Context) {
		id := c.Param("id")

		var fleet model.Fleet

		if err := db.Where("id = ?", id).First(&fleet).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success get fleet", fleet)
	})

	// only for arduino
	r.GET("/fleet/loc/:fleet_id", func(c *gin.Context) {
		ID := c.Param("fleet_id")
		lat := utils.StringToFloat(c.Query("lat"), c)
		lng := utils.StringToFloat(c.Query("lng"), c)

		var fleet model.Fleet
		if err := db.Where("id = ?", ID).First(&fleet).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		fleet.Latitude = lat
		fleet.Longitude = lng
		fleet.UpdatedAt = time.Now()

		if err := db.Save(&fleet).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success update fleet location", fleet)
	})
}
