package estatefacilityrel

// 屋苑_周边设施联系
type EstateFacilityRel struct {
	Id         uint `gorm:"column:id" db:"id" json:"id"`                            //序号
	EstateId   uint `gorm:"column:estate_id" db:"estate_id" json:"estate_id"`       //屋苑id
	FacilityId uint `gorm:"column:facility_id" db:"facility_id" json:"facility_id"` //周边设施id
}

//重写表名
func (EstateFacilityRel) TableName() string {
	return "hk591_data.estate_facility_rel"
}
