package WithDrawDbCtrl

func GetAllWithDraw(sortStr string) ([]Withdraw, error) {
	var withdraw []Withdraw
	err := Db.Order(sortStr).Where("status = ?", 0).Find(&withdraw).Error
	if err != nil {
		return nil, err
	}
	return withdraw, nil
}
