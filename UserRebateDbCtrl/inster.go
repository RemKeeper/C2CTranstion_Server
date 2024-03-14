package UserRebateDbCtrl

import (
	"log"
	"time"
)

func AddUserRebate(userID int, coinRebate int, balanceRebate float64) {
	var userRebate UserRebate
	if err := Db.Where("user_id = ?", userID).First(&userRebate).Error; err != nil {
		// Record not found, create a new one
		userRebate = UserRebate{
			UserID:        userID,
			CoinRebate:    coinRebate,
			BalanceRebate: balanceRebate,
			CreatedAt:     time.Now(),
		}
		if err := Db.Create(&userRebate).Error; err != nil {
			log.Printf("Failed to insert into user_rebates table: %v", err)
		} else {
			log.Printf("New entry added to user_rebates table: %+v", userRebate)
		}
	} else {
		// Record found, update it
		userRebate.CoinRebate += coinRebate
		userRebate.BalanceRebate += balanceRebate
		if err := Db.Save(&userRebate).Error; err != nil {
			log.Printf("Failed to update user_rebates table: %v", err)
		} else {
			log.Printf("Entry updated in user_rebates table: %+v", userRebate)
		}
	}
}
