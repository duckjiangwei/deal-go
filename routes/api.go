// Package routes 注册路由
package routes

import (
	"hk591_go/app/http/controllers/api/v1/auth"

	controllers "hk591_go/app/http/controllers/api/v1"

	"github.com/gin-gonic/gin"
)

// RegisterAPIRoutes 注册网页相关路由
func RegisterAPIRoutes(r *gin.Engine) {

	//v1 的路由组
	v1 := r.Group("/v1")
	{

		/* ------------------ 测试 ------------------ */
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}

		/* ------------------ 用户 ------------------ */
		usersGroup := v1.Group("/users")
		uc := new(controllers.UsersController)
		usersGroup.GET("", uc.Index)

		/* ------------------ 屋苑 ------------------ */
		estateGroup := v1.Group("/estate")
		ec := new(controllers.EsatteController)
		estateGroup.GET("", ec.Index)                          //屋苑列表頁
		estateGroup.GET("/detail", ec.Show)                    //屋苑詳情
		estateGroup.GET("/deal-record", ec.DealRecord)         //成交记录
		estateGroup.GET("/deal-chart", ec.DealChartData)       //成交图表
		estateGroup.POST("/collection", ec.Collection)         //收藏
		estateGroup.POST("/un-collection", ec.UnCollection)    //取消收藏
		estateGroup.GET("/collection-list", ec.CollectionList) //收藏列表

		/* ------------------ 配置项 ------------------ */
		configGroup := v1.Group("/config")
		cc := new(controllers.ConfigController)
		configGroup.GET("", cc.GetConfig) //获取区域
	}
}
