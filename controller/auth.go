package controller

import (
	"bytes"
	"cloud.google.com/go/storage"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"hackfest/model"
	"hackfest/utils"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	projectID  = "kiri-410313" // FILL IN WITH YOURS
	bucketName = "image-kiri"  // FILL IN WITH YOURS
)

type ClientUploader struct {
	cl         *storage.Client
	projectID  string
	bucketName string
	uploadPath string
}

var uploader *ClientUploader

func init() {
	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "./kiri-410313-39ffa45bef0d.json") // FILL IN WITH YOUR FILE PATH
	client, err := storage.NewClient(context.Background())
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	uploader = &ClientUploader{
		cl:         client,
		bucketName: bucketName,
		projectID:  projectID,
		// uploadPath: "qr-files/",
	}
}

// UploadQRCode uploads a QR code image
func (c *ClientUploader) UploadQRCode(data *bytes.Buffer, object string) error {
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, time.Second*50)
	defer cancel()

	wc := c.cl.Bucket(c.bucketName).Object(c.uploadPath + object).NewWriter(ctx)
	if _, err := io.Copy(wc, data); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}
	if err := wc.Close(); err != nil {
		return fmt.Errorf("Writer.Close: %v", err)
	}

	return nil
}

func Auth(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.POST("/register", func(c *gin.Context) {
		var userRegister model.UserRegister

		if err := c.BindJSON(&userRegister); err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		validEmail := utils.EmailValidator(userRegister.Email)

		if !validEmail {
			utils.HttpRespFailed(c, http.StatusInternalServerError, "Invalid Email")
			return
		}

		err := utils.PasswordValidator(userRegister.Password)
		if err != nil {
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
			QRCode:      "",
			Name:        userRegister.Name,
			KiriBalance: 0,
			KiriPoint:   0,
			CreatedAt:   time.Now(),
		}

		qrContent := fmt.Sprintf("%s/api/v1/qr/%s", os.Getenv("BASE_URL"), newUser.ID.String())
		qrCodeBuffer, err := utils.GenerateQRCode(qrContent)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		err = uploader.UploadQRCode(qrCodeBuffer, fmt.Sprintf("qrcodes/%s_qr.png", newUser.ID.String()))
		if err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		publicURL := fmt.Sprintf("https://storage.googleapis.com/%s/%s/%s", bucketName, "qrcodes", fmt.Sprintf("%s_qr.png", newUser.ID.String()))

		newUser.QRCode = publicURL

		if err := db.Create(&newUser).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusInternalServerError, err.Error())
			return
		}

		userResponse := model.UserResponse{
			ID:          newUser.ID,
			Name:        newUser.Name,
			Phone:       newUser.Phone,
			Email:       newUser.Email,
			Password:    newUser.Password,
			KiriBalance: newUser.KiriBalance,
			KiriPoint:   newUser.KiriPoint,
			CreatedAt:   newUser.CreatedAt,
			UpdatedAt:   newUser.UpdatedAt,
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success Register", userResponse)
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
