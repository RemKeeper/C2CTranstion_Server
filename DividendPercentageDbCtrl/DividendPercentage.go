package DividendPercentageDbCtrl

import (
	"C2CTranstion_Server/OrderDbCtrl"
	"C2CTranstion_Server/SettingDbCtrl"
	"C2CTranstion_Server/StatisticDbCtrl"
	"C2CTranstion_Server/UserDbCtrl"
	"C2CTranstion_Server/UserRebateDbCtrl"
)

func GetDividendPercentage(order OrderDbCtrl.Order) error {
	var PayUserId uint32
	var DividendPercentageI float64
	var DividendPercentageII float64
	var AddCoinBalance float64
	var AddBalance float64

	switch order.ProductType {
	case 0:
		//增加统计
		StatisticDbCtrl.AddCoinNumberOfTransactions(uint32(order.Count))

		err := UserDbCtrl.UserTransactionEditCoinBalance(true, uint(order.SuretyID), order.Count+int(float64(order.Count)*SettingDbCtrl.GlobalConfig.BountyPercentage))
		if err != nil {
			return err
		}
		PayUserId = order.SellerID
		DividendPercentageI = SettingDbCtrl.GlobalConfig.CoinDividendPercentageI
		DividendPercentageII = SettingDbCtrl.GlobalConfig.CoinDividendPercentageII
		inviterI, inviterII, err := GetUserInviter(uint(PayUserId))
		if err != nil {
			return err
		}
		switch {
		case inviterI != 0 && inviterII == 0:
			AddCoinBalance = order.TotalPrice * DividendPercentageI
			UserRebateDbCtrl.AddUserRebate(int(inviterI), int(AddCoinBalance), 0)
			err := UserDbCtrl.UserTransactionEditCoinBalance(true, inviterI, int(AddCoinBalance))
			StatisticDbCtrl.AddCoinDividend(AddCoinBalance)
			if err != nil {
				return err
			}
		case inviterI != 0 && inviterII != 0:
			AddCoinBalance = order.TotalPrice * DividendPercentageI
			UserRebateDbCtrl.AddUserRebate(int(inviterI), int(AddCoinBalance), 0)
			err := UserDbCtrl.UserTransactionEditCoinBalance(true, inviterI, int(AddCoinBalance))
			StatisticDbCtrl.AddCoinDividend(AddCoinBalance)
			if err != nil {
				return err
			}
			AddCoinBalance = order.TotalPrice * DividendPercentageII
			UserRebateDbCtrl.AddUserRebate(int(inviterII), int(AddCoinBalance), 0)
			err = UserDbCtrl.UserTransactionEditCoinBalance(true, inviterII, int(AddCoinBalance))
			StatisticDbCtrl.AddCoinDividend(AddCoinBalance)
			if err != nil {
				return err
			}
		}
	case 1:
		StatisticDbCtrl.AddGemNumberOfTransactions(uint(order.Count))

		err := UserDbCtrl.UserTransactionEditBalance(true, uint(order.SuretyID), order.TotalPrice+order.TotalPrice*SettingDbCtrl.GlobalConfig.BountyPercentage)
		if err != nil {
			return err
		}

		PayUserId = order.PurchasersID
		DividendPercentageI = SettingDbCtrl.GlobalConfig.DividendPercentageI
		DividendPercentageII = SettingDbCtrl.GlobalConfig.DividendPercentageII
		inviterI, inviterII, err := GetUserInviter(uint(PayUserId))
		if err != nil {
			return err
		}
		switch {
		case inviterI != 0 && inviterII == 0:
			AddBalance = order.TotalPrice * DividendPercentageI
			UserRebateDbCtrl.AddUserRebate(int(inviterI), 0, AddBalance)
			err := UserDbCtrl.UserTransactionEditBalance(true, inviterI, AddBalance)
			if err != nil {
				return err
			}
			StatisticDbCtrl.AddTotalDividend(AddBalance)
		case inviterI != 0 && inviterII != 0:
			AddBalance = order.TotalPrice * DividendPercentageI
			UserRebateDbCtrl.AddUserRebate(int(inviterI), 0, AddBalance)
			err := UserDbCtrl.UserTransactionEditBalance(true, inviterI, AddBalance)
			if err != nil {
				return err
			}
			StatisticDbCtrl.AddTotalDividend(AddBalance)
			AddBalance = order.TotalPrice * DividendPercentageII
			UserRebateDbCtrl.AddUserRebate(int(inviterII), 0, AddBalance)
			err = UserDbCtrl.UserTransactionEditBalance(true, inviterII, AddBalance)
			if err != nil {
				return err
			}
			StatisticDbCtrl.AddTotalDividend(AddBalance)
		}
	}
	return nil
}

func GetUserInviter(userId uint) (uint, uint, error) {
	user, err := UserDbCtrl.QueryUserByID(userId)
	if err != nil {
		return 0, 0, err
	}
	user2, err := UserDbCtrl.QueryUserByID(user.InviterId)
	if err != nil {
		return 0, 0, err
	}
	return user.InviterId, user2.InviterId, nil

}

func OrderNeedPay(order OrderDbCtrl.Order) float64 {
	var PayUserId uint32
	var DividendPercentageI float64
	var DividendPercentageII float64
	switch order.ProductType {
	case 0:
		PayUserId = order.SellerID
		DividendPercentageI = SettingDbCtrl.GlobalConfig.CoinDividendPercentageI
		DividendPercentageII = SettingDbCtrl.GlobalConfig.CoinDividendPercentageII
	case 1:
		PayUserId = order.PurchasersID
		DividendPercentageI = SettingDbCtrl.GlobalConfig.DividendPercentageI
		DividendPercentageII = SettingDbCtrl.GlobalConfig.DividendPercentageII
	}
	inviterI, inviterII, err := GetUserInviter(uint(PayUserId))
	if err != nil {
		return 0
	}
	var Ratio float64
	Ratio += SettingDbCtrl.GlobalConfig.BountyPercentage
	Ratio += SettingDbCtrl.GlobalConfig.PlatformCommission
	StatisticDbCtrl.AddPlatformCumulativeProfit(order.TotalPrice * SettingDbCtrl.GlobalConfig.PlatformCommission)
	switch {
	case inviterI == 0 && inviterII == 0:
		Ratio = 1
	case inviterI != 0 && inviterII == 0:
		Ratio = 1 + DividendPercentageI
	case inviterI != 0 && inviterII != 0:
		Ratio = 1 + DividendPercentageI + DividendPercentageII
	}
	return order.TotalPrice * Ratio
}
