package StatisticDbCtrl

func GetStatistic() Statistic {
	var statistic Statistic
	Db.First(&statistic)
	return statistic
}
