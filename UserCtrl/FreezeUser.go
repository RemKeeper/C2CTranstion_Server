package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"strconv"
)

type FreezeUserStruct struct {
	UserID int `json:"user_id"`
}

func FreezeUser(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "Bad Request"})
	}

	var EditUser FreezeUserStruct

	err = json.Unmarshal(data, &EditUser)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "无法解析请求数据"})
		return
	}
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "获取cookie失败"})
		return
	}
	User, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "获取用户信息失败"})
		return
	}
	if User.UserID == uint(EditUser.UserID) {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "不能冻结自己"})
		return
	}

	EUser, err := UserDbCtrl.FreezeUser(UserDbCtrl.User{UserID: uint(EditUser.UserID)})
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "数据库错误" + err.Error()})
		return
	}
	c.JSON(200, CommunicationStructure.Message{Code: 200, Msg: "冻结成功 " + strconv.Itoa(int(EUser))})

}

func UnfreezeUser(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "Bad Request"})
	}

	var user FreezeUserStruct

	err = json.Unmarshal(data, &user)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{Code: 400, Msg: "无法解析请求数据"})
		return
	}
	EUser, err := UserDbCtrl.UnfreezeUser(UserDbCtrl.User{UserID: uint(user.UserID)})
	if err != nil {
		c.JSON(500, CommunicationStructure.Message{Code: 500, Msg: "数据库错误" + err.Error()})
		return
	}
	c.JSON(200, CommunicationStructure.Message{Code: 200, Msg: "解冻成功 " + strconv.Itoa(int(EUser))})
}
