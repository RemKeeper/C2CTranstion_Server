package main

import (
	"C2CTranstion_Server/DividendPercentageICtrl"
	"C2CTranstion_Server/Image"
	"C2CTranstion_Server/Middleware"
	"C2CTranstion_Server/OrderCtrl"
	"C2CTranstion_Server/SettingCtrl"
	"C2CTranstion_Server/StatisticCtrl"
	"C2CTranstion_Server/Toast"
	"C2CTranstion_Server/UserCtrl"
	"C2CTranstion_Server/UserRebateCtrl"
	"C2CTranstion_Server/WithDrawCtrl"
	"C2CTranstion_Server/transaction"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

//var Db *gorm.DB

func main() {
	r := gin.Default()
	//请求数限制
	r.Use(Middleware.CheckIp)
	//10M 请求限制
	//r.Use(Middleware.CheckRequestSize)
	r.Use(Middleware.CheckPri)

	r.Use(Middleware.NoCache())

	//r.Use(Test.CORSMiddleware())
	//2M 请求限制
	r.Use(Middleware.SizeLimit)
	//go TimMer()
	r.GET("/GetPermission", UserCtrl.GetPermission)

	//注册登陆
	r.POST("/register", UserCtrl.Register)
	r.POST("/login", UserCtrl.Login)

	//静态文件交互

	r.Static("/image", "./images")

	r.POST("/upload", Image.Upload)
	r.POST("/getOrderImage", Image.GetOrderImage)
	//全局通知
	r.GET("/getToast", Toast.GetMessage)
	r.POST("/confirmedToast", Toast.ConfirmedMessage)

	//r.Static("/m", "./static")

	//Admin 部分
	adminGroup := r.Group("/admin")
	{
		adminGroup.POST("/deleteUser", UserCtrl.DeleteUser)
		//adminGroup.POST("/editUserPri", UserCtrl.EditUserPri)
		adminGroup.POST("/freezeUser", UserCtrl.FreezeUser)
		adminGroup.POST("/unFreezeUser", UserCtrl.UnfreezeUser)
		adminGroup.GET("/getAllUser", UserCtrl.GetAllUser)

		adminGroup.GET("/getAllOrder", OrderCtrl.GetAllOrder)
		adminGroup.POST("/deleteOrder", OrderCtrl.DeleteOrder)
		adminGroup.POST("/freezeOrder", OrderCtrl.FreezeOrder)
		adminGroup.POST("/unFreezeOrder", OrderCtrl.UnFreezeOrder)

		adminGroup.GET("/getAllWithDraw", WithDrawCtrl.GetAllWithDraw)
		adminGroup.POST("/deleteWithDraw", WithDrawCtrl.DeleteWithdraw)
		adminGroup.POST("/refusalWithDraw", WithDrawCtrl.RefusalWithdraw)

		adminGroup.GET("/getSettings", SettingCtrl.GetSetting)

		adminGroup.POST("/saveSettings", SettingCtrl.SaveSetting)

		adminGroup.GET("/getStatistic", StatisticCtrl.GetStatistic)

		//adminGroup.POST("/editUserQuota", UserCtrl.EditUserQuota)

		adminGroup.POST("/editUser", UserCtrl.EditUser)
	}

	//User 部分
	userGroup := r.Group("/user")
	{

		userGroup.GET("/logout", UserCtrl.Logout)
		userGroup.GET("/getUserInfo", UserCtrl.GetUserInfo)
		userGroup.GET("/getSelfInfo", UserCtrl.GetSelfInfo)
		userGroup.POST("/setSelfInfo", UserCtrl.SetSelfInfo)

		userGroup.POST("/transferCoin", UserCtrl.TransferCoin)
		userGroup.POST("/transferBalance", UserCtrl.TransferBalance)

		userGroup.GET("/getAllOrder", OrderCtrl.UserGetAllOrder)
		userGroup.POST("/createDispute", OrderCtrl.CreateDispute)
		userGroup.GET("/getWaitOrder", OrderCtrl.GetWaitOrder)
		userGroup.GET("/getFreezeOrder", OrderCtrl.GetFreezeOrder)

		userGroup.GET("/getInvitedUser", UserCtrl.GetInvitedUser)
		userGroup.GET("/getInvitedUserRecursive", UserCtrl.GetInvitedUserRecursive)

		userGroup.GET("/getUserRebate", UserRebateCtrl.GetUserRebate)

		userGroup.POST("/createWithDraw", WithDrawCtrl.CreateWithDraw)

		userGroup.GET("/getOrderById", UserCtrl.GetOrderById)

	}

	transactionGroup := r.Group("/transaction")

	{
		transactionGroup.POST("/createOrder", transaction.CreateOrder)
		transactionGroup.GET("/getAllNormalOrder", transaction.GetAllNormalOrder)

		transactionGroup.GET("/getOrderNeedPay", DividendPercentageICtrl.GetOrderNeedPay)

		transactionGroup.POST("/carryTransaction", transaction.CarryTransaction)
		transactionGroup.POST("/confirmTransaction", transaction.ConfirmTransaction)
	}

	//Surety 部分
	suretyGroup := r.Group("/surety")
	{
		suretyGroup.GET("/suretyGetSelfOrder", OrderCtrl.SuretyGetSelfOrder)
		suretyGroup.GET("/suretyGetWaitOrder", OrderCtrl.SuretyGetWaitOrder)
		suretyGroup.POST("/carrySurety", transaction.CarrySurety)
		suretyGroup.POST("/confirmSurety", transaction.ConfirmSurety)

	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code":    "0",
			"message": "pong",
		})
	})

	err := r.Run(":25621")
	if err != nil {
		log.Println(err)
		return
	}
}

func TimMer() {
	t := time.NewTicker(30 * time.Second)
	for {
		select {
		case <-t.C:
			r, err := http.Get("https://f100.rem.asia/")
			if err != nil {
				os.Exit(1)
			}
			b, _ := io.ReadAll(r.Body)
			if string(b) != "Hello World!" {
				os.Exit(1)
			}
		}
	}
}
