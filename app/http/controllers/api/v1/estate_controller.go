package v1

import (
	"hk591_go/app/requests"
	"hk591_go/app/services"
	"hk591_go/pkg/response"

	"github.com/gin-gonic/gin"
)

type EsatteController struct {
	BaseAPIController
}

//屋苑列表
func (ec *EsatteController) Index(c *gin.Context) {
	//获取请求参数，并验证
	request := requests.EstateIndexRequest{}
	if ok := requests.Validate(c, &request, requests.EstateIndex); !ok {
		return
	}

	//获取屋苑列表
	estate, pager := services.GetEstate(c, request)

	response.Data(c, 200, gin.H{
		"lists": estate,
		"pager": pager,
	})
}

//屋苑詳情
func (ec *EsatteController) Show(c *gin.Context) {
	id := c.Query("id")
	estate := services.GetSingleEstate(id)
	if len(estate) == 0 {
		response.Abort404(c, "數據不存在")
	}
	response.Data(c, 200, gin.H{
		"data": estate,
	})
}

//成交记录
func (ec *EsatteController) DealRecord(c *gin.Context) {
	//获取请求参数，并验证
	request := requests.DealRecordRequest{}
	if ok := requests.Validate(c, &request, requests.EstateIndex); !ok {
		return
	}
	//获取屋苑列表
	estate, pager := services.GetDealRecord(c, request)

	response.Data(c, 200, gin.H{
		"lists": estate,
		"pager": pager,
	})
}

//成交图表数据
func (ec *EsatteController) DealChartData(c *gin.Context) {
	id := c.Query("id")
	estate := services.GetDealChartData(id)
	if len(estate) == 0 {
		response.Abort404(c, "數據不存在")
	}
	response.Data(c, 200, gin.H{
		"data": estate,
	})
}

//收藏
func (ec *EsatteController) Collection(c *gin.Context) {
	idcode := c.PostForm("idcode")
	estateId := c.PostForm("estate_id")
	res := services.Collection(idcode, estateId)
	if !res {
		response.Error(c, 500)
	}
	response.Success(c)
}

//取消收藏
func (ec *EsatteController) UnCollection(c *gin.Context) {
	idcode := c.PostForm("idcode")
	estateId := c.PostForm("estate_id")
	res := services.UnCollection(idcode, estateId)
	if !res {
		response.Error(c, 500)
	}
	response.Success(c)
}

//收藏列表
func (ec *EsatteController) CollectionList(c *gin.Context) {
	//从 redis 拿到所有收藏 id
	idcode := c.Query("idcode")
	ids := services.GetCollectionIds(idcode)

	//获取详情
	collectList := services.GetAllEstateOfCollection(ids)
	response.Data(c, 200, gin.H{
		"lists": collectList,
	})
}
