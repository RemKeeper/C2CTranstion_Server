package WithDrawCtrl

import (
	"C2CTranstion_Server/CommunicationStructure"
	"C2CTranstion_Server/WithDrawDbCtrl"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllWithDraw(c *gin.Context) {
	sortStr := c.Query("sort")
	withdraw, err := WithDrawDbCtrl.GetAllWithDraw(sortStr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, CommunicationStructure.Message{
			Code: http.StatusInternalServerError,
			Msg:  " 获取提现记录失败 " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, withdraw)

}
