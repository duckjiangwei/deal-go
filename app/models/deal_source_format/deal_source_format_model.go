package dealsourceformat

type DealSourceFormat struct {
	Id                     int    `gorm:"column:id" db:"id" json:"id"`
	EstateId               int    `gorm:"column:estate_id" db:"estate_id" json:"estate_id"`                                     //屋苑id
	PhaseId                int    `gorm:"column:phase_id" db:"phase_id" json:"phase_id"`                                        //期id
	ThisMonthPriceUser     string `gorm:"column:this_month_price_user" db:"this_month_price_user" json:"this_month_price_user"` //当月平均尺价
	ThisMonthPriceBuild    string `gorm:"column:this_month_price_build" db:"this_month_price_build" json:"this_month_price_build"`
	ThreeMonthPriceUser    string `gorm:"column:three_month_price_user" db:"three_month_price_user" json:"three_month_price_user"` //最近3个月平均尺价
	ThreeMonthPriceBuild   string `gorm:"column:three_month_price_build" db:"three_month_price_build" json:"three_month_price_build"`
	NewPriceUser           string `gorm:"column:new_price_user" db:"new_price_user" json:"new_price_user"`                                     //最新成交尺价
	NewPriceUserDate       string `gorm:"column:new_price_user_date" db:"new_price_user_date" json:"new_price_user_date"`                      //最近成交月份
	YearHighPriceUser      string `gorm:"column:year_high_price_user" db:"year_high_price_user" json:"year_high_price_user"`                   //一年内最高尺价
	YearHighPriceUserDate  string `gorm:"column:year_high_price_user_date" db:"year_high_price_user_date" json:"year_high_price_user_date"`    //一年最高价月份
	YearLowPriceUser       string `gorm:"column:year_low_price_user" db:"year_low_price_user" json:"year_low_price_user"`                      //一年内最低尺价
	YearLowPriceUserDate   string `gorm:"column:year_low_price_user_date" db:"year_low_price_user_date" json:"year_low_price_user_date"`       //一年内最低尺价时间
	NewPriceBuild          string `gorm:"column:new_price_build" db:"new_price_build" json:"new_price_build"`                                  //最近成交价(建筑)
	YearHighPriceBuild     string `gorm:"column:year_high_price_build" db:"year_high_price_build" json:"year_high_price_build"`                //一年最高价(建筑)
	YearHighPriceBuildDate string `gorm:"column:year_high_price_build_date" db:"year_high_price_build_date" json:"year_high_price_build_date"` //一年最高价(建筑)时间
	YearLowPriceBuild      string `gorm:"column:year_low_price_build" db:"year_low_price_build" json:"year_low_price_build"`                   //一年最低价(建筑)
	YearLowPriceBuildDate  string `gorm:"column:year_low_price_build_date" db:"year_low_price_build_date" json:"year_low_price_build_date"`    //一年最低价(建筑)时间
}

//重写表名
func (DealSourceFormat) TableName() string {
	return "hk591_data.deal_source_format"
}
