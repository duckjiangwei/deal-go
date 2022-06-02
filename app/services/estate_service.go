package services

import (
	dealsourcedata "hk591_go/app/models/deal_source_data"
	"hk591_go/app/models/estate"
	"hk591_go/app/models/sale"
	"hk591_go/app/requests"
	"hk591_go/pkg/config"
	"hk591_go/pkg/database"
	"hk591_go/pkg/paginator"
	"hk591_go/pkg/redis"
	"strings"
	"time"

	estatebasefacility "hk591_go/app/models/estate_base_facility"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//获取屋苑列表
func GetEstate(c *gin.Context, request requests.EstateIndexRequest) ([]map[string]interface{}, paginator.Paging) {
	//每页行数
	if request.PerPage == 0 {
		request.PerPage = 20
	}

	//获取屋苑数据
	tx := database.DB.Model(&estate.Estate{}).Preload("EstateImg", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("show_sign=?", 1).Order("is_cover DESC")
	}).Preload("DealSource").Where("show_sign", 1)
	// .Preload("EstateFacilityRel")
	//大区域
	if request.AreaId > 0 {
		tx = tx.Where("area_id=?", request.AreaId)
	}
	//小区域
	if request.DistrictId != "" {
		district_ids := strings.Split(request.DistrictId, ",")
		tx = tx.Where(map[string]interface{}{"district_id": district_ids})
	}
	//价格
	if request.Price > 0 {
		minPrice, maxPrice := conditionOfPrice(request.Price)
		tx = tx.Joins("left join hk591_data.deal_source_format as sf ON sf.estate_id=estate.id").Where("sf.this_month_price_user between ? and ?", minPrice, maxPrice)
	}
	//楼龄
	if request.Age > 0 {
		now := time.Now()
		startYear := now.AddDate(-request.Age, 0, 0).Format("2006-01")
		if request.Age == 31 {
			tx = tx.Where("join_time < ?", startYear)
		} else {
			tx = tx.Where("join_time > ?", startYear)
		}
	}
	tx = tx.Scopes(estate.IsShowAndIsCheck)
	estateDatas, paging := estate.Paginate(c, request.PerPage, tx)

	//获取周边设施数据
	facility := estatebasefacility.All()

	//格式化
	results := estate.EstateListFormat(estateDatas, facility)

	return results, paging
}

//屋苑詳情
func GetSingleEstate(id string) (result map[string]interface{}) {
	var estateData estate.Estate

	//显示的关联数
	relationCount := config.GetInt("estate.relation_count")

	database.DB.Preload("EstateImg", func(tx *gorm.DB) *gorm.DB {
		return tx.Where("show_sign", 1).Order("is_cover DESC")
	}).Preload("DealSource").Preload("Sale", func(tx *gorm.DB) *gorm.DB {
		return tx.Scopes(sale.OnlineAndIsNewEstateRel).Where("hasimg=?", 1).Limit(relationCount)
	}).Scopes(estate.IsShowAndIsCheck).First(&estateData, id)

	//数据未找到，返回 404
	if estateData.Id == 0 {
		return
	}

	//格式化
	result = estate.SingleEstateFormat(estateData)
	return result
}

func conditionOfPrice(price int) (minPrice, maxPrice int) {
	switch price {
	case 1:
		minPrice = 0
		maxPrice = 10000
	case 2:
		minPrice = 10000
		maxPrice = 13000
	case 3:
		minPrice = 13000
		maxPrice = 15000
	case 4:
		minPrice = 15000
		maxPrice = 18000
	case 5:
		minPrice = 18000
		maxPrice = 20000
	case 6:
		minPrice = 20000
		maxPrice = 100000000
	default:
		minPrice = 0
		maxPrice = 100000000
	}
	return
}

func GetDealRecord(c *gin.Context, request requests.DealRecordRequest) ([]map[string]interface{}, paginator.Paging) {
	//每页行数
	if request.PerPage == 0 {
		request.PerPage = 3
	}

	//获取屋苑数据
	tx := database.DB.Model(&dealsourcedata.DealSourceData{}).Where("estate_id = ? and floor_id > ?", request.Id, 0).Order("id desc")
	dealData, paging := dealsourcedata.Paginate(c, request.PerPage, tx)

	//格式化
	results := dealsourcedata.DealDataFormat(dealData)

	return results, paging
}

type Deal struct {
	Month    string `json:"month"`
	Count    string `json:"count"`
	AvgPrice string `json:"avg_price"`
}

func GetDealChartData(id string) (result []Deal) {
	//查近一年的成交
	now := time.Now()
	lastYear := now.AddDate(0, -11, 0).Format("20060102")

	dealGroup := []Deal{}
	database.DB.
		Raw("select DATE_FORMAT(deal_date, '%Y%m') AS month,COUNT(*) AS count,AVG(user_price) AS avg_price from hk591_data.deal_source_data where estate_id=? and deal_date>? group by month", id, lastYear).
		Scan(&dealGroup)

	result = dealGroup
	return
}

func Collection(idcode, estateId string) bool {
	key := "estate_collection:" + idcode
	res := redis.Redis.Zadd(key, float64(time.Now().Unix()), estateId)
	return res
}

func UnCollection(idcode, estateId string) bool {
	key := "estate_collection:" + idcode
	//转切片，批量删除

	res := redis.Redis.Zrem(key, estateId)
	return res
}

func GetCollectionIds(idcode string) []string {
	key := "estate_collection:" + idcode
	count := redis.Redis.Zcard(key)
	estatdIds := redis.Redis.Zrange(key, 0, count)
	return estatdIds
}

func GetAllEstateOfCollection(ids []string) (result []map[string]interface{}) {
	list := []estate.Estate{}
	database.DB.Preload("DealSource").Where("id in ?", ids).Clauses(clause.OrderBy{
		Expression: clause.Expr{SQL: "FIELD(id,?)", Vars: []interface{}{ids}, WithoutParentheses: true},
	}).Find(&list)

	//格式化
	result = estate.EstateCollectionFormat(list)
	return
}
