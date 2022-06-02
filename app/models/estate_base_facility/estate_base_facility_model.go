package estatebasefacility

import (
	"time"
)

// 周边设施表
type EstateBaseFacility struct {
	Id       uint      `gorm:"column:id" db:"id" json:"id"`                      //序号
	Type     uint8     `gorm:"column:type" db:"type" json:"type"`                //类型（1港铁站，2医院，3公园，4商场）
	Name     string    `gorm:"column:name" db:"name" json:"name"`                //名称
	Posttime time.Time `gorm:"column:posttime" db:"posttime" json:"posttime"`    //添加时间
	ShowSign uint8     `gorm:"column:show_sign" db:"show_sign" json:"show_sign"` //显示标记（1显示，0隐藏）
}

//重写表名
func (EstateBaseFacility) TableName() string {
	return "hk591_data.estate_base_facility"
}
