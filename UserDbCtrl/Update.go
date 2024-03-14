package UserDbCtrl

import (
	"errors"
	"gorm.io/gorm"
)

func EditUserPri(user User) (uint, error) {
	if err := Db.Model(&user).Where("user_id = ?", user.UserID).Update("pri_label", user.PriLabel).Error; err != nil {
		return 0, err
	}
	return user.UserID, nil

}

func EditQuota(user User) (uint, error) {
	if err := Db.Model(&user).Where("user_id = ?", user.UserID).Update("coin_balance_quota", user.CoinBalanceQuota).Update("balance_quota", user.BalanceQuota).Error; err != nil {
		return 0, err
	}
	return user.UserID, nil

}

func FreezeUser(user User) (uint, error) {
	var EditUser User
	if err := Db.Model(&user).Where("user_id = ?", user.UserID).Update("is_freeze", true).Update("pri_label", 0).First(&EditUser).Error; err != nil {
		return 0, err
	}
	Db.Model(&Cookie{}).Where("user_id = ?", user.UserID).Update("cookie", "")
	return user.UserID, nil
}

func UnfreezeUser(user User) (uint, error) {
	if err := Db.Model(&user).Where("user_id = ?", user.UserID).Update("is_freeze", false).Update("pri_label", 2).Error; err != nil {
		return 0, err
	}
	return user.UserID, nil
}

func UserEditSelfInfo(cookie string, UserSet User) error {
	user, err := GetUserByCookie(cookie)
	if err != nil {
		return err
	}
	err = Db.Model(&user).Where("user_id = ?", user.UserID).Updates(UserSet).Error

	if err != nil {
		return err
	}
	return nil
}

func UserTransactionEditBalance(IsAdd bool, UserId uint, Balance float64) error {
	if Balance < 0 {
		return errors.New("金额不能为负数")
	}
	if IsAdd {
		err := Db.Model(&User{UserID: UserId}).Update("balance", gorm.Expr("balance + ?", Balance)).Error
		if err != nil {
			return err
		}
	} else {
		err := Db.Model(&User{UserID: UserId}).Update("balance", gorm.Expr("balance - ?", Balance)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func UserTransactionEditCoinBalance(IsAdd bool, UserId uint, CoinBalance int) error {
	if CoinBalance < 0 {
		return errors.New("金额不能为负数")
	}
	if IsAdd {
		err := Db.Model(&User{UserID: UserId}).Update("coin_balance", gorm.Expr("coin_balance + ?", CoinBalance)).Error
		if err != nil {
			return err
		}
	} else {
		err := Db.Model(&User{UserID: UserId}).Update("coin_balance", gorm.Expr("coin_balance - ?", CoinBalance)).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func TransferCoin(userId uint, toUserId uint, coinBalance int) error {
	err := UserTransactionEditCoinBalance(false, userId, coinBalance)
	if err != nil {
		return err
	}
	err = UserTransactionEditCoinBalance(true, toUserId, coinBalance)
	if err != nil {
		return err
	}
	return nil
}

func TransferBalance(userId uint, toUserId uint, balance float64) error {
	err := UserTransactionEditBalance(false, userId, balance)
	if err != nil {
		return err
	}
	err = UserTransactionEditBalance(true, toUserId, balance)
	if err != nil {
		return err
	}
	return nil
}
