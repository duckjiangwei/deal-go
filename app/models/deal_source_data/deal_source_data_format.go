package dealsourcedata

import "time"

func DealDataFormat(deal []DealSourceData) (results []map[string]interface{}) {
	for _, d := range deal {
		result := make(map[string]interface{})
		//基础信息
		result["id"] = d.Id
		result["user_price"] = d.UserPrice
		result["unit_address"] = d.UnitAddress
		result["user_area"] = d.UserArea
		result["deal_price"] = d.DealPrice
		t1, _ := time.Parse("20060102", d.DealDate)
		result["deal_date"] = t1.Format("2006-01-02")
		results = append(results, result)
	}
	return
}
