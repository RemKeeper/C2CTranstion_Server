package UserRebateDbCtrl

func GetUserRebate(userId uint) (UserRebate, error) {
	var userRebate UserRebate
	err := Db.Where("user_id = ?", userId).First(&userRebate).Error
	if err != nil {
		return UserRebate{}, err
	}
	return userRebate, nil
}
