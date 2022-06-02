package estatefloor

import "time"

// 屋苑_楼层表
type EstateFloor struct {
	Id           uint      `gorm:"column:id" db:"id" json:"id"`                                  //序号
	BlockId      uint      `gorm:"column:block_id" db:"block_id" json:"block_id"`                //座id
	Floor        string    `gorm:"column:floor" db:"floor" json:"floor"`                         //楼层
	Unit         string    `gorm:"column:unit" db:"unit" json:"unit"`                            //单元
	BuildingArea uint      `gorm:"column:building_area" db:"building_area" json:"building_area"` //建筑面积
	UseArea      uint      `gorm:"column:use_area" db:"use_area" json:"use_area"`                //使用面积
	Posttime     time.Time `gorm:"column:posttime" db:"posttime" json:"posttime"`                //提交时间
	ShowSign     uint8     `gorm:"column:show_sign" db:"show_sign" json:"show_sign"`             //是否显示（1是，0否）
}
