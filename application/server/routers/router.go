package routers

import (
	v1 "application/api/v1"

	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	r := gin.Default()

	apiV1 := r.Group("/api/v1")
	{
		//在花括号{}内，定义了一系列的路由。每个路由都使用了apiV1的实例来添加。例如，apiV1.GET添加了一个HTTP GET路由/api/v1/hello，它将调用v1.Hello函数来处理请求。
		apiV1.GET("/hello", v1.Hello)
		apiV1.POST("/queryAccountList", v1.QueryAccountList)
		apiV1.POST("/createRealEstate", v1.CreateRealEstate)
		apiV1.POST("/queryRealEstateList", v1.QueryRealEstateList)
		apiV1.POST("/createSelling", v1.CreateSelling)
		apiV1.POST("/createSellingByBuy", v1.CreateSellingByBuy)
		apiV1.POST("/querySellingList", v1.QuerySellingList)
		apiV1.POST("/querySellingListByBuyer", v1.QuerySellingListByBuyer)
		apiV1.POST("/updateSelling", v1.UpdateSelling)
		apiV1.POST("/createDonating", v1.CreateDonating)
		apiV1.POST("/queryDonatingList", v1.QueryDonatingList)
		apiV1.POST("/queryDonatingListByGrantee", v1.QueryDonatingListByGrantee)
		apiV1.POST("/updateDonating", v1.UpdateDonating)
	}
	return r
}
