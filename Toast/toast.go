package Toast

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/Utils"
	"fmt"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"

	"net/http"
)

type ConfirmRequest struct {
	UUID string `json:"uuid"`
}

type MessageStructToast struct {
	ToastKey string `json:"toast_key"`
}

var MessageSlice = make(map[uint]map[string]Message, 500)

func AddMessage(msg Message) {
	uuid, err := Utils.GenerateCustomUUID()
	if err != nil {
		return
	}

	CopyMessage := MessageSlice[msg.Userid]

	if CopyMessage == nil {
		CopyMessage = make(map[string]Message, 500)
	}
	CopyMessage[uuid] = msg
	MessageSlice[msg.Userid] = CopyMessage

	fmt.Println(MessageSlice)
}

func ConfirmedMessage(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		return
	}
	var confirmRequest ConfirmRequest
	err = json.Unmarshal(data, &confirmRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "解析数据失败",
		})
		return
	}

	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取cookie失败",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取用户信息失败",
		})
		return
	}
	delete(MessageSlice[user.UserID], confirmRequest.UUID)
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  "确认成功",
	})
}

func GetMessage(c *gin.Context) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  "获取cookie失败",
		})
		return
	}
	user, err := UserDbCtrl.GetUserByCookie(cookie)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  "获取用户信息失败",
		})
		return
	}
	msg := MessageSlice[user.UserID]
	c.JSON(http.StatusOK, msg)
}
