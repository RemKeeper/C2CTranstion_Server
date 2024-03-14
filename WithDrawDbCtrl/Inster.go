package WithDrawDbCtrl

import (
	"C2CTranstion_Server/UserDbCtrl"
	"time"
)

func CreateWithDraw(user UserDbCtrl.User, productType uint, amount float64) error {
	withdraw := Withdraw{
		ProductType:    productType,
		UserID:         user.UserID,
		Balance:        user.Balance,
		CashWithdrawal: amount,
		Status:         0,
		Time:           time.Now(),
	}
	err := Db.Create(&withdraw).Error
	if err != nil {
		return err
	}
	return nil
}
