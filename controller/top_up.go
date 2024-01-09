package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"gorm.io/gorm"
	"hackfest/middleware"
	"hackfest/model"
	"hackfest/utils"
	"net/http"
	"os"
)

func TopUp(db *gorm.DB, q *gin.Engine) {
	r := q.Group("/api/v1")

	r.POST("/topup", middleware.Authorization(), func(c *gin.Context) {
		methodStr := c.Query("method")
		method := utils.StringToInteger(methodStr, c)
		amountStr := c.Query("amount")
		amount := utils.StringToFloat(amountStr, c)

		ID, _ := c.Get("id")

		var user model.User
		if err := db.Where("id = ?", ID).First(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		midtransClient := coreapi.Client{}
		midtransClient.New(os.Getenv("MIDTRANS_SERVER_KEY"), midtrans.Sandbox)
		orderID := utils.RandomOrderID()

		req := &coreapi.ChargeReq{}

		if method == 1 {
			req = &coreapi.ChargeReq{
				PaymentType: "gopay",
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  orderID,
					GrossAmt: int64(amount),
				},
				Gopay: &coreapi.GopayDetails{
					EnableCallback: true,
					CallbackUrl:    "https://example.com/callback",
				},
				CustomerDetails: &midtrans.CustomerDetails{
					FName: user.Name,
					Email: user.Email,
					Phone: user.Phone,
				},
			}
		} else if method == 2 {
			req = &coreapi.ChargeReq{
				PaymentType: "shopeepay",
				TransactionDetails: midtrans.TransactionDetails{
					OrderID:  orderID,
					GrossAmt: int64(amount),
				},
				ShopeePay: &coreapi.ShopeePayDetails{
					CallbackUrl: "https://example.com/callback",
				},
				CustomerDetails: &midtrans.CustomerDetails{
					FName: user.Name,
					Email: user.Email,
					Phone: user.Phone,
				},
			}
		}

		resp, err := midtransClient.ChargeTransaction(req)
		if err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		user.KiriBalance += amount

		if err := db.Save(&user).Error; err != nil {
			utils.HttpRespFailed(c, http.StatusNotFound, err.Error())
			return
		}

		utils.HttpRespSuccess(c, http.StatusOK, "Success top up", resp)
	})
}
