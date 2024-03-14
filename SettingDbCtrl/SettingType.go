package SettingDbCtrl

type Config struct {
	EditHere                 int     `gorm:"column:edit_here;comment:'便于数据库Where的无意义主键'" json:"edit_here"`
	BountyPercentage         float64 `gorm:"type:decimal(3,2);column:bounty_percentage;comment:'赏金比例'" json:"bounty_percentage"`
	PlatformCommission       float64 `gorm:"type:decimal(3,2);column:platform_commission;comment:'平台抽成'" json:"platform_commission"`
	MinCoinTransaction       int     `gorm:"column:min_coin_transaction;comment:'最小金币交易数量'" json:"min_coin_transaction"`
	MinGemTransaction        int     `gorm:"column:min_gem_transaction;comment:'最小宝石交易数量'" json:"min_gem_transaction"`
	DividendPercentageI      float64 `gorm:"type:decimal(3,2);column:dividend_percentage_I;comment:'分成比例 一级'" json:"dividend_percentage_I"`
	DividendPercentageII     float64 `gorm:"type:decimal(3,2);column:dividend_percentage_II;comment:'分成比例 二级'" json:"dividend_percentage_II"`
	CoinDividendPercentageI  float64 `gorm:"type:decimal(3,2);column:coin_dividend_percentage_I;comment:'金币分成比例 一级'" json:"coin_dividend_percentage_I"`
	CoinDividendPercentageII float64 `gorm:"type:decimal(3,2);column:coin_dividend_percentage_II;comment:'金币分成比例 二级'" json:"coin_dividend_percentage_II"`
	Bulletin                 string  `gorm:"type:text;column:bulletin;comment:'全平台公告'" json:"bulletin"`
}

func (Config) TableName() string {
	return "config"
}
