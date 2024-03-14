package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	UserID          uint    `gorm:"primary_key;AUTO_INCREMENT;column:user_id;type:int(8) unsigned zerofill;not null;" json:"user_id"`
	UserName        string  `gorm:"unique;column:user_name;type:char(30);size:30;not null;" json:"user_name"`
	PriLabel        uint    `gorm:"column:pri_label;type:int unsigned;not null;" json:"pri_label"`
	IsFreeze        bool    `gorm:"column:is_freeze;type:tinyint;default:0;not null;" json:"is_freeze"`
	Balance         float64 `gorm:"column:balance;type:float;not null;" json:"balance"`
	CoinBalanceQuta int     `gorm:"column:coin_balance_quota;type:int;not null;" json:"coin_balance_quota"`
	BalanceQuota    float64 `gorm:"column:balance_quota;type:float;not null;" json:"balance_quota"`
	CoinBalance     int     `gorm:"column:coin_balance;type:int;not null;" json:"coin_balance"`
}

func GetAllUser(c *gin.Context) {
	var users []UserDbCtrl.User
	users, err := UserDbCtrl.GetAllUser()
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "数据库错误" + err.Error()})
		return
	}
	var userResponse []UserResponse
	for _, user := range users {
		userResponse = append(userResponse, UserResponse{
			UserID:          user.UserID,
			UserName:        user.UserName,
			PriLabel:        user.PriLabel,
			IsFreeze:        user.IsFreeze,
			Balance:         user.Balance,
			CoinBalanceQuta: user.CoinBalanceQuota,
			BalanceQuota:    user.BalanceQuota,
			CoinBalance:     user.CoinBalance,
		})
	}
	c.JSON(200, userResponse)
}
