package sale

import "gorm.io/gorm"

//在线且为新屋苑
func OnlineAndIsNewEstateRel(tx *gorm.DB) *gorm.DB {
	return tx.Where("status = ? and closed = ? and is_new_estate = ?", 2, 0, 1)
}
