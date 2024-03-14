package WithDrawCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/WithDrawDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

type WithDrawRequest struct {
	WithDrawID int `json:"withdraw_id"`
}

func DeleteWithdraw(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 获取数据失败 " + err.Error(),
		})
		return
	}
	var completeWithdrawRequest WithDrawRequest
	err = json.Unmarshal(data, &completeWithdrawRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 解析数据失败 " + err.Error(),
		})
		return
	}
	err = WithDrawDbCtrl.CompleteWithdraw(completeWithdrawRequest.WithDrawID)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 完成提现失败 " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, CommunicationStructure.Message{
		Code: http.StatusOK,
		Msg:  " 完成提现成功 ",
	})
}
