package shop

import (
	"github.com/gin-gonic/gin"
	"net/http"
	model "server/app/model/system"
	service "server/app/service/shop"
	"server/pkg/code"
	"server/pkg/response"
	"server/pkg/validator"
)

//ShopService
type ShopService interface {
	Test(c *gin.Context)
	Login(c *gin.Context)
}

// ShopApiService 服务层数据处理
type ShopApiService struct {
	Shop service.ShopService
}

// NewShopApi 创建构造函数简单工厂模式
func NewShopApi() ShopService {
	return ShopApiService{Shop: service.NewShopService()}
}

func (es ShopApiService) Test(c *gin.Context) {
	response.Success(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), map[string]interface{}{
		"num":   1,
		"total": 1,
	})
	return
}

func (es ShopApiService) Login(c *gin.Context) {
	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	if user.Username == "admin" && user.Password == "123456" {

	} else {
		response.Error(c, http.StatusBadRequest, code.ServerErr, "账号或者密码不正确", nil)
		return
	}

	//response.Success(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), map[string]interface{}{
	//	"num":   1,
	//	"total": 1,
	//})

	response.Success(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), nil)
	return
}
