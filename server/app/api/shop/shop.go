package shop

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"server/app/model/shop"
	model "server/app/model/system"
	service "server/app/service/shop"
	"server/global"
	"server/pkg/code"
	"server/pkg/response"
	"server/pkg/validator"
)

//ShopService
type ShopService interface {
	Test(c *gin.Context)
	Login(c *gin.Context)
	CreateUserInfo(c *gin.Context)
	UpdateUserInfo(c *gin.Context)
	GetUserInfoByPhone(c *gin.Context)
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
	traceId := uuid.New().String()

	response.SuccessN(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), map[string]interface{}{
		"num":   1,
		"total": 1,
	}, traceId)
	return
}

func (es ShopApiService) Login(c *gin.Context) {
	traceId := uuid.New().String()

	user := new(model.User)

	if err := c.ShouldBindJSON(&user); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	global.Log.Info("Login " + user.Username + " " + user.Password + " " + traceId)

	isAdmin := es.Shop.IsAdmin(user.Username, user.Password)

	if !isAdmin {
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, "账号或者密码不正确", nil, traceId)
		return
	}

	response.SuccessN(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), nil, traceId)
	return
}

func (es ShopApiService) CreateUserInfo(c *gin.Context) {
	traceId := uuid.New().String()

	userInfo := new(shop.ShopUserInfo)

	if err := c.ShouldBindJSON(&userInfo); err != nil {
		// 参数校验
		validator.HandleValidatorError(c, err)
		return
	}

	// 对 userInfo 进行赋值操作
	userInfoBytes, err := json.Marshal(userInfo)
	if err != nil {
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, "Marshal error", nil, traceId)
		return
	}
	userInfoString := string(userInfoBytes)

	global.Log.Info("CreateUserInfo " + userInfoString + " " + traceId)

	err = es.Shop.CreateUserInfo(userInfo, traceId)
	if err != nil {
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, err.Error(), nil, traceId)
		return
	}

	response.SuccessN(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), nil, traceId)
	return
}

func (es ShopApiService) UpdateUserInfo(c *gin.Context) {
	traceID := uuid.New().String()

	userInfo := new(shop.ShopUserInfo)
	if err := c.ShouldBindJSON(&userInfo); err != nil {
		validator.HandleValidatorError(c, err)
		return
	}

	err := es.Shop.UpdateUserInfo(userInfo, traceID)
	if err != nil {
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, "UpdateUserInfo 更新失败", nil, traceID)
		return
	}

	response.SuccessN(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), nil, traceID)
}

func (es ShopApiService) GetUserInfoByPhone(c *gin.Context) {
	traceID := uuid.New().String()

	phone := c.Query("phone")
	if phone == "" {
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, "GetUserInfoByPhone 参数错误", nil, traceID)
		return
	}

	userInfo, err := es.Shop.GetUserInfoByPhone(phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ErrorN(c, http.StatusNotFound, code.ServerErr, "GetUserInfoByPhone 查询结果不存在", nil, traceID)
			return
		}
		response.ErrorN(c, http.StatusBadRequest, code.ServerErr, "GetUserInfoByPhone 查询失败", nil, traceID)
		return
	}

	response.SuccessN(c, code.SUCCESS, code.GetErrMsg(code.SUCCESS), userInfo, traceID)
}
