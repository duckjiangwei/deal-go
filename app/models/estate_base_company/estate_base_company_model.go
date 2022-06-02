package estatebasecompany

// 物业公司表
type EstateBaseCompany struct {
	Id   uint   `gorm:"column:id" db:"id" json:"id"`       //序号
	Name string `gorm:"column:name" db:"name" json:"name"` //名称
}

//重写表名
func (EstateBaseCompany) TableName() string {
	return "hk591_data.estate_base_company"
}
