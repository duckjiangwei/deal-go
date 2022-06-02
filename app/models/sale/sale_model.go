package sale

import (
	warepicture "hk591_go/app/models/ware_picture"
	"hk591_go/pkg/config"
	"hk591_go/pkg/database"
	"strings"

	"gorm.io/gorm"
)

// 出售
type Sale struct {
	Id          int     `gorm:"column:id" db:"id" json:"id"`
	UserId      int     `gorm:"column:user_id" db:"user_id" json:"user_id"`                   //用戶id
	Title       string  `gorm:"column:title" db:"title" json:"title"`                         //標題
	Status      int8    `gorm:"column:status" db:"status" json:"status"`                      //狀態: 0 未使用 1 已過期 2 開啟 3 已成交 4 已成交
	Closed      int8    `gorm:"column:closed" db:"closed" json:"closed"`                      //是否關閉: 0未關閉 1自行關閉 2回收站 4客服關閉
	Isvip       int8    `gorm:"column:isvip" db:"isvip" json:"isvip"`                         //是否VIP
	StartTime   int     `gorm:"column:start_time" db:"start_time" json:"start_time"`          //開始時間
	EndTime     int     `gorm:"column:end_time" db:"end_time" json:"end_time"`                //結束時間
	DistrictId  int     `gorm:"column:district_id" db:"district_id" json:"district_id"`       //地區ID
	AreaId      int     `gorm:"column:area_id" db:"area_id" json:"area_id"`                   //區域ID
	Cid         int     `gorm:"column:cid" db:"cid" json:"cid"`                               //屋苑ID
	IsNewEstate uint8   `gorm:"column:is_new_estate" db:"is_new_estate" json:"is_new_estate"` //是否關聯新屋苑數據（1是，0否）
	Longitude   float64 `gorm:"column:longitude" db:"longitude" json:"longitude"`             //經度
	Latitude    float64 `gorm:"column:latitude" db:"latitude" json:"latitude"`                //緯度
	Purpose     int     `gorm:"column:purpose" db:"purpose" json:"purpose"`                   //物業用途
	Kind        int     `gorm:"column:kind" db:"kind" json:"kind"`                            //物業種類
	ShopKind    int     `gorm:"column:shop_kind" db:"shop_kind" json:"shop_kind"`             //物業類型
	Community   string  `gorm:"column:community" db:"community" json:"community"`             //屋苑大廈
	Address     string  `gorm:"column:address" db:"address" json:"address"`                   //租盤地址
	Room        int     `gorm:"column:room" db:"room" json:"room"`                            //幾房
	Hall        int     `gorm:"column:hall" db:"hall" json:"hall"`                            //幾廳
	Toilet      int     `gorm:"column:toilet" db:"toilet" json:"toilet"`                      //幾衛
	UseArea     int     `gorm:"column:use_area" db:"use_area" json:"use_area"`                //可使用面積
	BuildArea   int     `gorm:"column:build_area" db:"build_area" json:"build_area"`          //建築面積
	AreaCheck   uint8   `gorm:"column:area_check" db:"area_check" json:"area_check"`          //使用/建筑面积是否审核（0未审核，1已审核）
	Floor       int     `gorm:"column:floor" db:"floor" json:"floor"`                         //樓層1001:地下,1002:高層,1003:中層,1004:低層,1005:頂樓加建,1006:整栋
	Allfloor    int     `gorm:"column:allfloor" db:"allfloor" json:"allfloor"`                //總樓層
	FloorType   int     `gorm:"column:floor_type" db:"floor_type" json:"floor_type"`          //樓層類型（1地下2高層3中層4低層5楼顶加建）
	Price       int     `gorm:"column:price" db:"price" json:"price"`                         //租金
	Hasimg      int8    `gorm:"column:hasimg" db:"hasimg" json:"hasimg"`                      //是否有圖
	Cover       string  `gorm:"column:cover" db:"cover" json:"cover"`                         //封面

	Img []warepicture.WarePicture `gorm:"foreignKey:post_id" json:"img"` //has many 关联 warepicture
}

//重写表名
func (Sale) TableName() string {
	return "sale"
}

func (s *Sale) AfterFind(tx *gorm.DB) (err error) {
	img := []warepicture.WarePicture{}
	var url string
	database.DB.Model(s).Where("closed=? and type=?", 0, 2).Order("is_cover desc").Association("Img").Find(&img)
	if len(img) > 0 {
		file := strings.Split(img[0].Filename, ".")
		url = config.Get("domain.p1_url") + "/house/active/" + img[0].Path + file[0] + "_212x159.jpg"
	} else {
		url = config.Get("estate.no_photo_detail_small")
	}

	s.Cover = url
	return
}
