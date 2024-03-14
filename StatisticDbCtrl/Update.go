package StatisticDbCtrl

// AddTotalDividend 添加总分红统计
func AddTotalDividend(addNum float64) float64 {
	var statistic Statistic
	Db.First(&statistic)
	statistic.TotalDividend += addNum
	Db.Save(&statistic)
	return statistic.TotalDividend
}

// AddCoinDividend 添加金币分红统计
func AddCoinDividend(addNum float64) float64 {
	var statistic Statistic
	Db.First(&statistic)
	statistic.CoinDividend += addNum
	Db.Save(&statistic)
	return statistic.CoinDividend
}

// AddCoinNumberOfTransactions 添加金币交易数量统计
func AddCoinNumberOfTransactions(addNum uint32) uint32 {
	var statistic Statistic
	Db.First(&statistic)
	statistic.CoinNumberOfTransactions += addNum
	Db.Save(&statistic)
	return statistic.CoinNumberOfTransactions
}

// AddGemNumberOfTransactions 添加宝石交易数量统计
func AddGemNumberOfTransactions(addNum uint) uint {
	var statistic Statistic
	Db.First(&statistic)
	statistic.GemNumberOfTransactions += addNum
	Db.Save(&statistic)
	return statistic.GemNumberOfTransactions
}

// AddPlatformCumulativeProfit 添加平台累计利润统计
func AddPlatformCumulativeProfit(addNum float64) float64 {
	var statistic Statistic
	Db.First(&statistic)
	statistic.PlatformCumulativeProfit += addNum
	Db.Save(&statistic)
	return statistic.PlatformCumulativeProfit
}
