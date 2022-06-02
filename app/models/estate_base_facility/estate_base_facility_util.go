package estatebasefacility

import "hk591_go/pkg/database"

// 查询所以周边设施
func All() (facilitys []EstateBaseFacility) {
	database.DB.Find(&facilitys)
	return
}
