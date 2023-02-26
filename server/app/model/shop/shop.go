package shop

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

type ShopAdmin struct {
	ID         uint      `gorm:"primaryKey;autoIncrement"`
	Username   string    `gorm:"not null"`
	Password   string    `gorm:"not null"`
	CreateTime time.Time `gorm:"autoCreateTime"`
	UpdateTime time.Time `gorm:"autoUpdateTime"`
}

type ShopUserInfo struct {
	ID        int       `gorm:"primarykey"`
	Object    string    `gorm:"not null;comment:'商品名称'"`
	Addar     string    `gorm:"not null;comment:'用户地址'"`
	Name      string    `gorm:"not null;comment:'用户姓名'"`
	Phone     string    `gorm:"not null;comment:'用户电话'uniqueIndex"`
	WXNumber  string    `gorm:"not null;comment:'用户微信'"`
	Code      string    `gorm:"not null;comment:'快递单号';index:code_index"`
	CreatedAt time.Time `gorm:"autoCreateTime;comment:'创建时间'"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;comment:'更新时间'"`
}
