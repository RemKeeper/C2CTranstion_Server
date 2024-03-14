package UserRebateDbCtrl

import (
	"time"
)

type UserRebate struct {
	UserID        int       `gorm:"column:user_id;type:int;not null"`
	CoinRebate    int       `gorm:"column:coin_rebate;type:int;not null"`
	BalanceRebate float64   `gorm:"column:balance_rebate;type:float;not null"`
	CreatedAt     time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP"`
}

func (UserRebate) TableName() string {
	return "user_rebates"
}
