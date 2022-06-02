package estatedeveloperrel

// 发展商表
type EstateBaseDeveloper struct {
	Id   uint   `gorm:"column:id" db:"id" json:"id"`       //序号
	Name string `gorm:"column:name" db:"name" json:"name"` //名称
}

func (EstateBaseDeveloper) TableName() string {
	return "hk591_data.estate_base_developer"
}
