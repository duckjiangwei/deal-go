package estate

import (
	estatebasecompany "hk591_go/app/models/estate_base_company"
	estatedeveloperrel "hk591_go/app/models/estate_base_developer"
	estatebasefacility "hk591_go/app/models/estate_base_facility"
	estatebaseinfrastructure "hk591_go/app/models/estate_base_infrastructure"
	estateimg "hk591_go/app/models/estate_img"
	"hk591_go/app/models/sale"
	config_data "hk591_go/config"
	"hk591_go/pkg/config"
	"hk591_go/pkg/database"
	"hk591_go/pkg/redis"
	"time"

	"github.com/spf13/cast"
)

func EstateListFormat(estate []Estate, facilitys []estatebasefacility.EstateBaseFacility) (results []map[string]interface{}) {
	for _, e := range estate {
		result := make(map[string]interface{})
		//基础信息
		result["id"] = e.Id
		result["name"] = e.NameZh
		result["address"] = e.AddressZh
		result["join_time"] = e.JoinTime
		result["phase_number"] = e.PhaseNumber
		result["block_number"] = e.BlockNumber
		result["unit_number"] = e.UnitNumber

		//近期尺价
		if e.DealSource.ThisMonthPriceUser != "" {
			result["price"] = e.DealSource.ThisMonthPriceUser
		} else if e.DealSource.ThreeMonthPriceUser != "" {
			result["price"] = e.DealSource.ThreeMonthPriceUser
		} else {
			result["price"] = ""
		}

		//历史尺价
		result["historical_price"] = e.DealSource.ThreeMonthPriceUser

		//区域
		result["area"] = config_data.AreaData[e.AreaId]
		result["district"] = config_data.DistrictData[e.DistrictId]

		//处理封面图
		if len(e.EstateImg) != 0 {
			result["img"] = estateimg.PathFormat(e.EstateImg[0].FilePath)["list"]
		} else {
			result["img"] = config.Get("estate.no_photo_list")
		}

		//3个月内成交数
		now := time.Now()
		startYear := now.AddDate(0, -3, 0).Format("20060102")
		dealCount := database.DB.Model(e).Where("deal_date > ?", startYear).Association("DealData").Count()
		result["deal_count"] = dealCount
		results = append(results, result)
	}

	return
}

func SingleEstateFormat(estate Estate) (result map[string]interface{}) {
	result = make(map[string]interface{})
	//基础信息
	result["id"] = estate.Id
	result["name"] = estate.NameZh
	result["area"] = config_data.AreaData[estate.AreaId]
	result["district"] = config_data.DistrictData[estate.DistrictId]
	result["address"] = estate.AddressZh
	result["longitude"] = estate.Longitude //经度
	result["latitude"] = estate.Latitude   //维度

	//楼盘信息
	result["deal_data"] = estate.DealSource
	result["phase_number"] = estate.PhaseNumber
	result["block_number"] = estate.BlockNumber
	result["unit_number"] = estate.UnitNumber
	result["join_date"] = estate.JoinTime
	//处理图片
	imgFormat := []interface{}{}
	if len(estate.EstateImg) != 0 {
		for _, img := range estate.EstateImg {
			imgFormat = append(imgFormat, estateimg.PathFormat(img.FilePath))
		}
	} else {
		img := map[string]interface{}{}
		img["large"] = config.Get("estate.no_photo_detail_large")
		img["middle"] = config.Get("estate.no_photo_detail_middle")
		img["small"] = config.Get("estate.no_photo_detail_small")
		imgFormat = append(imgFormat, img)
	}
	result["estate_img"] = imgFormat

	//发展商
	developer := []estatedeveloperrel.EstateBaseDeveloper{}
	database.DB.Model(estate).Association("EstateBaseDeveloper").Find(&developer)
	result["developer"] = developer

	//物业设施
	infrastructure := []estatebaseinfrastructure.EstateBaseInfrastructure{}
	database.DB.Model(estate).Association("EstateBaseInfrastructure").Find(&infrastructure)
	result["infrastructure"] = infrastructure

	//物业公司
	company := []estatebasecompany.EstateBaseCompany{}
	database.DB.Model(estate).Association("Estatebasecompany").Find(&company)
	result["property_company"] = company

	//从 redis 获取浏览数
	key := "estate_browse_count"
	browse := redis.Redis.Hget(key, cast.ToString(estate.Id))
	result["browse"] = browse

	//相关售物件
	result["sale"] = estate.Sale
	return
}

func EstateCollectionFormat(list []Estate) (results []map[string]interface{}) {
	for _, e := range list {
		result := make(map[string]interface{})
		//基础信息
		result["id"] = e.Id
		result["name"] = e.NameZh
		result["address"] = e.AddressZh

		//近期尺价
		if e.DealSource.ThisMonthPriceUser != "" {
			result["price"] = e.DealSource.ThisMonthPriceUser
		} else if e.DealSource.ThreeMonthPriceUser != "" {
			result["price"] = e.DealSource.ThreeMonthPriceUser
		} else {
			result["price"] = ""
		}

		//区域
		result["area"] = config_data.AreaData[e.AreaId]
		result["district"] = config_data.DistrictData[e.DistrictId]

		now := time.Now()

		//近 1 个月上架
		startMonthStr := now.AddDate(0, -1, 0).Format("2006-01-02 15:04:05")
		t, _ := time.Parse("2006-01-02 15:04:05", startMonthStr)
		startMonth := t.Unix()
		var saleCount int64
		database.DB.Model(sale.Sale{}).Where("is_new_estate = ? and cid = ? and posttime > ?", 1, e.Id, startMonth).Count(&saleCount)
		result["sale_count"] = saleCount

		//近半年成交
		startYear := now.AddDate(0, -6, 0).Format("20060102")
		dealCount := database.DB.Model(e).Where("deal_date > ?", startYear).Association("DealData").Count()
		result["deal_count"] = dealCount
		results = append(results, result)
	}

	return
}
