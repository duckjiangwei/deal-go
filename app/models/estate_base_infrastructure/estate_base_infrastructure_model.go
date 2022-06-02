package estatebaseinfrastructure

// 基础设施表
type EstateBaseInfrastructure struct {
	Id   uint   `gorm:"column:id" db:"id" json:"id"`       //序号
	Name string `gorm:"column:name" db:"name" json:"name"` //名称
}

//重写表名
func (EstateBaseInfrastructure) TableName() string {
	return "hk591_data.estate_base_infrastructure"
}
