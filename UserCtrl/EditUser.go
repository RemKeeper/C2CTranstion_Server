package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	"strconv"
)

type EditUserRequest struct {
	UserId           uint    `json:"user_id"`
	PriLabel         uint    `json:"pri_label"`
	CoinBalanceQuota int     `json:"coin_balance_quota"`
	BalanceQuota     float64 `json:"balance_quota"`
}

func EditUser(c *gin.Context) {
	var editUserRequest EditUserRequest
	err := c.BindJSON(&editUserRequest)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "解析数据失败 " + err.Error(),
		})
		return
	}
	user, err := UserDbCtrl.QueryUserByID(editUserRequest.UserId)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "查询用户失败 " + err.Error(),
		})
		return
	}
	if user.PriLabel > 4 {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "管理员不能设置其他管理员，有编辑需要请联系后台管理员",
		})
		return
	}

	_, err = UserDbCtrl.EditQuota(UserDbCtrl.User{
		UserID:           editUserRequest.UserId,
		CoinBalanceQuota: editUserRequest.CoinBalanceQuota,
		BalanceQuota:     editUserRequest.BalanceQuota,
	})
	if editUserRequest.PriLabel > 4 {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "权限等级不合法,最大为4 高级担保人",
		})
		return
	}

	_, err = UserDbCtrl.EditUserPri(UserDbCtrl.User{UserID: user.UserID, PriLabel: editUserRequest.PriLabel})
	if err != nil {
		return
	}
	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  "修改成功" + strconv.Itoa(int(user.UserID)),
	})

}
