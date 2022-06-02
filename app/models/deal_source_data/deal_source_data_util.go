package dealsourcedata

import (
	"hk591_go/pkg/app"
	"hk591_go/pkg/database"
	"hk591_go/pkg/paginator"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//分页
func Paginate(c *gin.Context, perPage int, tx *gorm.DB) (deal []DealSourceData, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		tx,
		&deal,
		app.V1URL(database.TableName(&DealSourceData{})),
		perPage,
	)
	return
}
