package transaction

import (
	"C2CTranstion_Server/TransactionDbCtrl"
	"github.com/gin-gonic/gin"
)

func GetAllNormalOrder(c *gin.Context) {
	order, err := TransactionDbCtrl.GetNormalOrder()
	if err != nil {
		return
	}
	c.JSON(200, order)
}
