package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"strconv"
)

type DeleteUserMessage struct {
	UserID int `json:"user_id"`
}

func DeleteUser(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法读取请求数据",
		})
		return
	}
	var UserData DeleteUserMessage

	err = json.Unmarshal(data, &UserData)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法解析请求数据",
		})
		return
	}
	user, err := UserDbCtrl.DeleteUser(UserDbCtrl.User{UserID: uint(UserData.UserID)})
	if err != nil {
		return
	}
	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  "删除成功" + strconv.Itoa(int(user)),
	})
}
