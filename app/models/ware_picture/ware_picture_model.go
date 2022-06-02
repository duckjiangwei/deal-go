package warepicture

// 物件图片表
type WarePicture struct {
	Id          uint   `gorm:"column:id" db:"id" json:"id"`
	UserId      int    `gorm:"column:user_id" db:"user_id" json:"user_id"`                //会员id
	Type        uint8  `gorm:"column:type" db:"type" json:"type"`                         //類型 1 出租 2 出售 5 6 7 建案
	PostId      int    `gorm:"column:post_id" db:"post_id" json:"post_id"`                //刊登id
	Path        string `gorm:"column:path" db:"path" json:"path"`                         //目录
	Filename    string `gorm:"column:filename" db:"filename" json:"filename"`             //圖片名稱
	Posttime    int    `gorm:"column:posttime" db:"posttime" json:"posttime"`             //上傳時間
	CateId      uint8  `gorm:"column:cate_id" db:"cate_id" json:"cate_id"`                //0 物件图 1 房屋格局图
	PictureType uint8  `gorm:"column:picture_type" db:"picture_type" json:"picture_type"` //图片类型（1单位相片、2格局图、3屋苑相片）
	KindId      int8   `gorm:"column:kind_id" db:"kind_id" json:"kind_id"`                //房屋類型類型ID
	Closed      uint8  `gorm:"column:closed" db:"closed" json:"closed"`                   //關閉狀態：0 正常 1 回收站 2 物理刪除 3 异常 4隐藏(会员看不见)
	IsCover     uint8  `gorm:"column:is_cover" db:"is_cover" json:"is_cover"`             //是否为封面
	AddVal      int8   `gorm:"column:add_val" db:"add_val" json:"add_val"`                //附加值
	LayoutVal   uint8  `gorm:"column:layout_val" db:"layout_val" json:"layout_val"`       //圖片格局值
	Note        string `gorm:"column:note" db:"note" json:"note"`                         //圖片說明
	Device      int8   `gorm:"column:device" db:"device" json:"device"`                   //1PC 2ios 4android 6wap
	Qrcode      int8   `gorm:"column:qrCode" db:"qrCode" json:"qrCode"`                   //0非扫码上传1扫码上传
}

//重写表名
func (WarePicture) TableName() string {
	return "ware_picture"
}
