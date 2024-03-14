package WithDrawCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/WithDrawDbCtrl"
	"github.com/gin-gonic/gin"
	json "github.com/json-iterator/go"
	"net/http"
)

func RefusalWithdraw(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 获取数据失败 " + err.Error(),
		})
		return
	}
	var refusalWithdrawRequest WithDrawRequest
	err = json.Unmarshal(data, &refusalWithdrawRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, CommunicationStructure.Message{
			Code: http.StatusBadRequest,
			Msg:  " 解析数据失败 " + err.Error(),
		})
		return
	}
	err = WithDrawDbCtrl.RefusalWithdraw(refusalWithdrawRequest.WithDrawID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  " 拒绝提现失败 " + err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, CommunicationStructure.Message{
			Code: http.StatusOK,
			Msg:  " 拒绝提现成功 ",
		})
	}
}
