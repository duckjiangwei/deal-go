//屋苑模型
package estate

import (
	dealsourcedata "hk591_go/app/models/deal_source_data"
	dealsourceformat "hk591_go/app/models/deal_source_format"
	estatebasecompany "hk591_go/app/models/estate_base_company"
	estateDeveloper "hk591_go/app/models/estate_base_developer"
	estatebaseinfrastructure "hk591_go/app/models/estate_base_infrastructure"
	estateimg "hk591_go/app/models/estate_img"
	"hk591_go/app/models/sale"
)

type Estate struct {
	Id                   uint    `gorm:"column:id" db:"id" json:"id"`                                                          //序号
	NameZh               string  `gorm:"column:name_zh" db:"name_zh" json:"name_zh"`                                           //中文名称
	NameEn               string  `gorm:"column:name_en" db:"name_en" json:"name_en"`                                           //英文名称
	AreaId               uint    `gorm:"column:area_id" db:"area_id" json:"area_id"`                                           //大区id
	DistrictId           uint    `gorm:"column:district_id" db:"district_id" json:"district_id"`                               //小区id
	AddressZh            string  `gorm:"column:address_zh" db:"address_zh" json:"address_zh"`                                  //中文地址
	AddressEn            string  `gorm:"column:address_en" db:"address_en" json:"address_en"`                                  //英文地址
	JoinTime             string  `gorm:"column:join_time" db:"join_time" json:"join_time"`                                     //入伙时间
	UnitNumber           uint    `gorm:"column:unit_number" db:"unit_number" json:"unit_number"`                               //单位总数
	PhaseNumber          uint    `gorm:"column:phase_number" db:"phase_number" json:"phase_number"`                            //期数
	BlockNumber          uint    `gorm:"column:block_number" db:"block_number" json:"block_number"`                            //座数
	SchoolnetIdB         uint    `gorm:"column:schoolnet_id_b" db:"schoolnet_id_b" json:"schoolnet_id_b"`                      //校网大区
	SchoolnetIdM         uint    `gorm:"column:schoolnet_id_m" db:"schoolnet_id_m" json:"schoolnet_id_m"`                      //校网中区
	SchoolnetIdS         uint    `gorm:"column:schoolnet_id_s" db:"schoolnet_id_s" json:"schoolnet_id_s"`                      //校网小区
	DeveloperSpider      string  `gorm:"column:developer_spider" db:"developer_spider" json:"developer_spider"`                //发展商（爬取）
	CompanySpider        string  `gorm:"column:company_spider" db:"company_spider" json:"company_spider"`                      //物业公司（爬取）
	InfrastructureSpider string  `gorm:"column:infrastructure_spider" db:"infrastructure_spider" json:"infrastructure_spider"` //物业设施（爬取）
	Longitude            float64 `gorm:"column:longitude" db:"longitude" json:"longitude"`                                     //经度
	Latitude             float64 `gorm:"column:latitude" db:"latitude" json:"latitude"`                                        //纬度
	CheckStatus          uint8   `gorm:"column:check_status" db:"check_status" json:"check_status"`                            //审核状态（0未审核；1已审核）
	DealUseNumber        uint    `gorm:"column:deal_use_number" db:"deal_use_number" json:"deal_use_number"`                   //成交有用次數
	DealUselessNumber    uint    `gorm:"column:deal_useless_number" db:"deal_useless_number" json:"deal_useless_number"`       //成交無用次數
	Posttime             string  `gorm:"column:posttime" db:"posttime" json:"posttime"`                                        //添加时间
	ShowSign             uint8   `gorm:"column:show_sign" db:"show_sign" json:"show_sign"`                                     //显示标记（1显示，0隐藏）
	IsPhase              int     `gorm:"column:is_phase" db:"is_phase" json:"is_phase"`                                        //是否为屋苑期表数据 0否 >0期ID
	ExtraId              int     `gorm:"column:extra_id" db:"extra_id" json:"extra_id"`                                        //屋苑拓展ID，识别屋苑与期，排序和获取图片使用。1.逻辑排序，期在屋苑后面。2.获取图片 期的图片与屋苑图片一致。
	NhId                 int     `gorm:"column:nh_id" db:"nh_id" json:"nh_id"`                                                 //新盘中转过来的，对应new_house id
	Alias                string  `gorm:"column:alias" db:"alias" json:"alias"`                                                 //别名
	AliasEn              string  `gorm:"column:alias_en" db:"alias_en" json:"alias_en"`                                        //英文别名

	EstateImg                []estateimg.EstateImg                               `gorm:"foreignKey:estate_id" json:"estate_img"` //has many 关联 estate_img
	DealSource               dealsourceformat.DealSourceFormat                   `gorm:"foreignKey:estate_id" json:"deal_data"`  //has one 关联 deal_source_format
	Sale                     []sale.Sale                                         `gorm:"foreignKey:cid" json:"sale"`             ////has many 关联 sale
	DealData                 []dealsourcedata.DealSourceData                     `gorm:"foreignKey:estate_id" json:"deal"`       //has many 关联 deal_source_data
	EstateBaseDeveloper      []estateDeveloper.EstateBaseDeveloper               `gorm:"many2many:hk591_data.estate_developer_rel;joinForeignKey:estate_id;joinReferences:developer_id;"`
	EstateBaseInfrastructure []estatebaseinfrastructure.EstateBaseInfrastructure `gorm:"many2many:hk591_data.estate_infrastructure_rel;joinForeignKey:estate_id;joinReferences:infrastructure_id;"`
	Estatebasecompany        []estatebasecompany.EstateBaseCompany               `gorm:"many2many:hk591_data.estate_company_rel;joinForeignKey:estate_id;joinReferences:company_id;"`
}

type APIEstate struct {
	ID         uint
	NameZh     string
	EstateImg  interface{} `gorm:"foreignKey:estate_id" json:"estate_img"`
	DealSource interface{} `gorm:"foreignKey:estate_id" json:"deal_source_format"`
}

//重写表名
func (Estate) TableName() string {
	return "hk591_data.estate"
}
