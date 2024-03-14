package WithDrawDbCtrl

import (
	"C2CTranstion_Server/Toast"
	"C2CTranstion_Server/UserDbCtrl"
)

func CompleteWithdraw(withdrawId int) error {
	var withdraw Withdraw
	err := Db.Model(&Withdraw{}).Where("withdraw_id = ?", withdrawId).Update("status", 1).Find(&withdraw).Error
	switch withdraw.ProductType {
	case 0:
		err := UserDbCtrl.UserTransactionEditCoinBalance(false, withdraw.UserID, int(withdraw.CashWithdrawal))
		if err != nil {
			return err
		}
		Toast.AddMessage(Toast.Message{
			Userid: withdraw.UserID,
			Msg:    "您的提现已经完成，金币已经到账",
		})
	case 1:
		err := UserDbCtrl.UserTransactionEditBalance(false, withdraw.UserID, withdraw.CashWithdrawal)
		if err != nil {
			return err
		}
		Toast.AddMessage(Toast.Message{
			Userid: withdraw.UserID,
			Msg:    "您的提现已经完成，余额已经到账",
		})
	}

	if err != nil {
		return err
	}
	return nil
}

func RefusalWithdraw(withdrawId int) error {
	var withdraw Withdraw
	err := Db.Model(&Withdraw{}).Where("withdraw_id = ?", withdrawId).Update("status", 2).Find(&withdraw).Error
	if err != nil {
		return err
	}
	switch withdraw.ProductType {
	case 0:
		err := UserDbCtrl.UserTransactionEditCoinBalance(true, withdraw.UserID, int(withdraw.CashWithdrawal))
		if err != nil {
			return err
		}
	case 1:
		err := UserDbCtrl.UserTransactionEditBalance(true, withdraw.UserID, withdraw.CashWithdrawal)
		if err != nil {
			return err
		}
	}
	Toast.AddMessage(Toast.Message{
		Userid: withdraw.UserID,
		Msg:    "您的提现已经被拒绝",
	})
	return nil
}
