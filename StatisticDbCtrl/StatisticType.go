package StatisticDbCtrl

type Statistic struct {
	TotalDividend            float64 `json:"total_dividend" gorm:"column:TotalDividend"`
	CoinDividend             float64 `json:"coin_dividend" gorm:"column:CoinDividend"`
	CoinNumberOfTransactions uint32  `json:"coin_number_of_transactions" gorm:"column:CoinNumberOfTransactions"`
	GemNumberOfTransactions  uint    `json:"gem_number_of_transactions" gorm:"column:GemNumberOfTransactions"`
	PlatformCumulativeProfit float64 `json:"platform_cumulative_profit" gorm:"column:PlatformCumulativeProfit"`
}

func (Statistic) TableName() string {
	return "Statistic"
}
