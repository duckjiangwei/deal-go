package config

import "hk591_go/pkg/config"

var Age map[int]string
var Price map[int]string

func init() {
	//屋苑图片配置
	config.Add("estate", func() map[string]interface{} {
		return map[string]interface{}{
			// 列表默认图
			"no_photo_list": "https://www.591.com.hk/Public/Static/images/no_list.png",

			//详情页默认图
			"no_photo_detail_large":  "https://www.591.com.hk/Public/Static/images/nophoto_730x564.png",
			"no_photo_detail_middle": "https://www.591.com.hk/Public/Static/images/no_photo378X269.png",
			"no_photo_detail_small":  "https://www.591.com.hk/Public/Static/images/nophoto108x81.png",

			//详情页显示关联的物件数
			"relation_count": 3,
		}
	})

	//楼龄筛选
	Age = map[int]string{
		0:   "不限",
		10:  "10年以內",
		15:  "15年以內",
		20:  "20年以內",
		30:  "30年以內",
		200: "30年以上",
	}

	//尺价筛选
	Price = map[int]string{
		0: "不限",
		1: "10000元以下",
		2: "10000-13000元",
		3: "13000-15000元",
		4: "15000-18000元",
		5: "18000-20000元",
		6: "20000元以上",
	}
}
