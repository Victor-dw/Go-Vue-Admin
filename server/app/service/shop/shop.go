package shop

import (
	"errors"
	"server/app/model/shop"
	"server/global"
)

// ShopService
type ShopService interface {
	IsAdmin(username string, password string) bool                    //判断是否为管理员
	CreateUserInfo(userInfo *shop.ShopUserInfo, traceId string) error //创建用户信息
	UpdateUserInfo(userInfo *shop.ShopUserInfo, traceId string) error //更新用户信息
	GetUserInfoById(id int) (*shop.ShopUserInfo, error)               //根据Id查询用户信息
	GetUserInfoByPhone(phone string) (*shop.ShopUserInfo, error)      //根据手机号查询用户信息
}

type Shop struct{}

// NewShopService 构造函数
func NewShopService() ShopService {
	return Shop{}
}

func (s Shop) IsAdmin(username string, password string) bool {
	var admin shop.ShopAdmin
	if err := global.DB.Where("username = ?", username).First(&admin).Error; err != nil {
		return false
	}
	if admin.Password == password {
		return true
	}
	return false
}

func (s Shop) CreateUserInfo(userInfo *shop.ShopUserInfo, traceId string) error {
	err := global.DB.Create(&userInfo).Error
	if err != nil {
		// 处理插入失败的情况
		global.Log.Error("CreateUserInfo " + "数据库插入失败" + traceId)
		return err
	}
	return nil
}

// 更新用户信息
func (s Shop) UpdateUserInfo(userInfo *shop.ShopUserInfo, traceId string) error {
	userInfoByPhone, err := s.GetUserInfoByPhone(userInfo.Phone)
	if err != nil {
		global.Log.Error("UpdateUserInfo " + "用户不存在" + traceId)
		return err
	}

	// 检查用户是否存在
	if userInfoByPhone == nil {
		global.Log.Error("UpdateUserInfo " + "用户不存在" + traceId)
		return errors.New("user not found")
	}

	// 更新用户信息
	err = global.DB.Model(&shop.ShopUserInfo{}).Where("id = ?", userInfoByPhone.ID).
		Updates(shop.ShopUserInfo{
			Object:   userInfo.Object,
			Addar:    userInfo.Addar,
			Name:     userInfo.Name,
			Phone:    userInfo.Phone,
			WXNumber: userInfo.WXNumber,
			Code:     userInfo.Code,
		}).Error
	if err != nil {
		global.Log.Error("UpdateUserInfo " + "数据库更新失败" + traceId)
		return err
	}

	return nil
}

// 查询用户信息
func (s Shop) GetUserInfoById(id int) (*shop.ShopUserInfo, error) {
	var userInfo shop.ShopUserInfo
	err := global.DB.Where("id = ?", id).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}

func (s Shop) GetUserInfoByPhone(phone string) (*shop.ShopUserInfo, error) {
	var userInfo shop.ShopUserInfo
	err := global.DB.Where("phone = ?", phone).First(&userInfo).Error
	if err != nil {
		return nil, err
	}
	return &userInfo, nil
}
