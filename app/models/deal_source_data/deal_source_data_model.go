package dealsourcedata

import (
	estatefloor "hk591_go/app/models/estate_floor"
	"hk591_go/pkg/database"

	"gorm.io/gorm"
)

type DealSourceData struct {
	Id              int    `gorm:"column:id" db:"id" json:"id"`
	EstateId        int    `gorm:"column:estate_id" db:"estate_id" json:"estate_id"`                            //屋苑id
	PhaseId         int    `gorm:"column:phase_id" db:"phase_id" json:"phase_id"`                               //期id
	BlockId         int    `gorm:"column:block_id" db:"block_id" json:"block_id"`                               //座id
	FloorId         int    `gorm:"column:floor_id" db:"floor_id" json:"floor_id"`                               //单元id
	DealDate        string `gorm:"column:deal_date" db:"deal_date" json:"deal_date"`                            //成交日期
	DealPrice       int    `gorm:"column:deal_price" db:"deal_price" json:"deal_price"`                         //成交价格(单位:万)
	BuildArea       string `gorm:"column:build_area" db:"build_area" json:"build_area"`                         //建筑面积
	UserArea        string `gorm:"column:user_area" db:"user_area" json:"user_area"`                            //实用面积
	BuildPrice      string `gorm:"column:build_price" db:"build_price" json:"build_price"`                      //建筑尺价
	UserPrice       string `gorm:"column:user_price" db:"user_price" json:"user_price"`                         //实用尺价
	HkjDetailId     int    `gorm:"column:hkj_detail_id" db:"hkj_detail_id" json:"hkj_detail_id"`                //deal_source_detail表的id
	NewHouseSalesId int    `gorm:"column:new_house_sales_id" db:"new_house_sales_id" json:"new_house_sales_id"` //hk591.new_house_sales表 id
	UnitAddress     string `json:"unit_address"`
}

//重写表名
func (DealSourceData) TableName() string {
	return "hk591_data.deal_source_data"
}

type Result struct {
	SerialZh string
	NameZh   string
}

func (deal *DealSourceData) AfterFind(tx *gorm.DB) (err error) {
	var resultPhase Result
	unitAddress := ""
	database.DB.Raw("select serial_zh,name_zh from hk591_data.estate_phase where id=?", deal.PhaseId).Scan(&resultPhase)
	unitAddress = resultPhase.SerialZh + "期" + resultPhase.NameZh

	var resultBlock Result
	database.DB.Raw("select serial_zh,name_zh from hk591_data.estate_block where id=?", deal.BlockId).Scan(&resultBlock)
	unitAddress = unitAddress + " " + resultBlock.SerialZh + "座" + resultBlock.NameZh

	var floor estatefloor.EstateFloor
	database.DB.Raw("select * from hk591_data.estate_floor where id=?", deal.FloorId).Scan(&floor)
	unitAddress = unitAddress + " " + floor.Floor + " " + floor.Unit + "室"

	deal.UnitAddress = unitAddress
	return
}
