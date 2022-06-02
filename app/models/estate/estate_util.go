// Package estate 存放屋苑相关逻辑
package estate

import (
	"hk591_go/pkg/app"
	"hk591_go/pkg/database"
	"hk591_go/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// All 获取所有用户数据
func All() (estate []Estate) {
	database.DB.Find(&estate)
	return
}

//分页
func Paginate(c *gin.Context, perPage int, tx *gorm.DB) (estate []Estate, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		tx,
		&estate,
		app.V1URL(database.TableName(&Estate{})),
		perPage,
	)
	return
}

//常用查询
func IsShowAndIsCheck(tx *gorm.DB) *gorm.DB {
	return tx.Where("show_sign = ? and check_status = ? and is_phase = ?", 1, 1, 0)
}
