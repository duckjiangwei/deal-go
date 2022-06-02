package estateimg

import (
	"time"
)

// 屋苑图片表
type EstateImg struct {
	Id               uint      `gorm:"column:id" db:"id" json:"id"`                                              //序号
	Type             uint8     `gorm:"column:type" db:"type" json:"type"`                                        //图片类型（1外观图；2屋苑设施图）
	EstateId         uint      `gorm:"column:estate_id" db:"estate_id" json:"estate_id"`                         //屋苑id
	PhaseId          uint      `gorm:"column:phase_id" db:"phase_id" json:"phase_id"`                            //期id
	BlockId          uint      `gorm:"column:block_id" db:"block_id" json:"block_id"`                            //座id
	InfrastructureId uint      `gorm:"column:infrastructure_id" db:"infrastructure_id" json:"infrastructure_id"` //基础设施id
	FilePath         string    `gorm:"column:file_path" db:"file_path" json:"file_path"`                         //图片路径
	IsCover          uint8     `gorm:"column:is_cover" db:"is_cover" json:"is_cover"`                            //是否为封面（1是；0否）
	Posttime         time.Time `gorm:"column:posttime" db:"posttime" json:"posttime"`                            //添加时间
	ShowSign         uint8     `gorm:"column:show_sign" db:"show_sign" json:"show_sign"`                         //显示标记（1显示，0隐藏）
	EstatePhaseId    int       `gorm:"column:estate_phase_id" db:"estate_phase_id" json:"estate_phase_id"`       //屋苑ID或者期对应的屋苑ID
	IsTransfer       int8      `gorm:"column:is_transfer" db:"is_transfer" json:"is_transfer"`                   //是否中转,0否1是
}

//重写表名
func (EstateImg) TableName() string {
	return "hk591_data.estate_img"
}
