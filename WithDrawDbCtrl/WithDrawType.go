package WithDrawDbCtrl

import "time"

type Withdraw struct {
	WithdrawID     uint      `gorm:"primary_key;column:withdraw_id;comment:'提现申请ID'" json:"withdraw_id"`
	ProductType    uint      `gorm:"column:product_type;comment:'产品类型'" json:"product_type"`                          // 商品类型
	UserID         uint      `gorm:"column:user_id;comment:'用户ID'" json:"user_id"`                                    // 用户ID
	Balance        float64   `gorm:"type:decimal(10,2);column:balance;comment:'用户余额'" json:"balance"`                 // 用户余额
	CashWithdrawal float64   `gorm:"type:decimal(10,2);column:cash_withdrawal;comment:'提现金额'" json:"cash_withdrawal"` // 提现金额
	Status         int       `gorm:"column:status;comment:'状态'" json:"status"`                                        // 状态
	Time           time.Time `gorm:"type:datetime;column:time;default:CURRENT_TIMESTAMP;comment:'时间'" json:"time"`    // 时间
}

func (Withdraw) TableName() string {
	return "withdraw"
}
