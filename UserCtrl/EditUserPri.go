package UserCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"strconv"
)

type EditUserPriMessage struct {
	UserID   int `json:"user_id"`
	PriLabel int `json:"pri_label"`
}

func EditUserPri(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法读取请求数据",
		})
		return
	}
	var UserData EditUserPriMessage

	err = json.Unmarshal(data, &UserData)
	if err != nil {
		c.JSON(400, CommunicationStructure.Message{
			Code: 400,
			Msg:  "无法解析请求数据",
		})
		return
	}
	user, err := UserDbCtrl.EditUserPri(UserDbCtrl.User{UserID: uint(UserData.UserID), PriLabel: uint(UserData.PriLabel)})
	if err != nil {
		return
	}
	c.JSON(200, CommunicationStructure.Message{
		Code: 200,
		Msg:  "修改成功" + strconv.Itoa(int(user)),
	})

}
