package v1

import (
	"hk591_go/config"
	"hk591_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type ConfigController struct {
	BaseAPIController
}

//屋苑列表
func (ec *ConfigController) GetConfig(c *gin.Context) {
	//区域配置
	areaConfigData := map[string]interface{}{
		"area":     config.AreaData,     //大区域
		"district": config.DistrictData, //小区域
		"relation": config.Relation,     //区域关联
	}
	//楼龄
	age := config.Age
	//尺价
	price := config.Price

	response.Data(c, 200, gin.H{
		"area_config":  areaConfigData,
		"age_config":   age,
		"price_config": price,
	})
}
