package shop

import (
	"github.com/gin-gonic/gin"
	"server/app/api/shop"
)

// InitShopRouter 商户模块
func InitShopRouter(r *gin.RouterGroup) gin.IRouter {
	shopRouter := shop.NewShopApi()
	router := r.Group("/shop")
	{
		//router.POST("/create", exampleRouter.PostExample)        // 创建
		//router.GET("/id", exampleRouter.GetExample)              // 单条数据
		//router.GET("/list", exampleRouter.GetExampleList)        // 列表
		//router.PUT("/put", exampleRouter.PutExample)             // 更新
		//router.DELETE("/delete", exampleRouter.DeleteExample)    // 删除
		//router.DELETE("/remove", exampleRouter.DeleteExampleAll) // 全部删除
		//router.GET("/rank", exampleRouter.GetExampleRank)        // 排行榜
		//router.POST("/vote", exampleRouter.GetExampleVote)       // 投票
		//router.POST("/test", exampleRouter.Test)                 // 测试
		router.POST("/test", shopRouter.Test)                               //
		router.POST("/login", shopRouter.Login)                             //
		router.POST("/createUserInfo", shopRouter.CreateUserInfo)           //
		router.POST("/updateUserInfo", shopRouter.UpdateUserInfo)           //
		router.GET("/searchUserInfoByPhone", shopRouter.GetUserInfoByPhone) //
	}

	return r
}
