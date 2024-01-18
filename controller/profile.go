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

	r.POST("/edit-profile", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var userEditProfile model.UserEditProfile
		if err := c.BindJSON(&userEditProfile); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		validEmail := utils.EmailValidator(userEditProfile.Email)

		if !validEmail {
			utils.HttpRespFailed(c, http.StatusInternalServerError, "Invalid Email")
			return
		}

		user.Name = userEditProfile.Name
		user.Email = userEditProfile.Email
		user.Phone = userEditProfile.Phone

		if err := db.Save(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success edit user profile", nil)
	})

	r.POST("/change-password", middleware.Authorization(), func(c *gin.Context) {
		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		var userChangePassword model.UserChangePassword
		if err := c.BindJSON(&userChangePassword); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		if !utils.CompareHash(userChangePassword.OldPassword, user.Password) {
			utils.HttpRespFailed(c, http.StatusUnauthorized, "Wrong password!")
			return
		}

		if userChangePassword.NewPassword != userChangePassword.ConfirmPassword {
			utils.HttpRespFailed(c, http.StatusInternalServerError, "New password and confirm password not match!")
			return
		}

		err := utils.PasswordValidator(userChangePassword.NewPassword)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		hashed, err := utils.Hash(userChangePassword.NewPassword)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		user.Password = hashed

		if err := db.Save(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success change password", nil)

	})
}
