package models

import (
	"gorm.io/gorm"
	"service-all/app/init/global"
)

type ModelPrice struct {
	gorm.Model
	// 用户唯一标识
	Openid string `json:"openid"`
	// 用户昵称
	NickName string `json:"nick_name"`
	// 用户头像
	AvatarUrl string `json:"avatar_url"`
	// 1 未开奖 2开奖 3已领奖
	Status int    `json:"status"`
	Uuid   string `json:"uuid"`
}

type RequestPrice struct {
	// 用户唯一标识
	Openid string `json:"openid"`
	// 用户昵称
	NickName string `json:"nick_name"`
	// 用户头像
	AvatarUrl string `json:"avatar_url"`
	// 1 未开奖 2开奖 3已领奖
	Status int `json:"status"`
}

type ActionModelPrice interface {
	QueryWhereOpenid(Openid string) error
	Create() error
	Delete(openid string) error
	Update(openid string, data ModelPrice) error
}

func ActPrice(act ActionModelPrice) ActionModelPrice {
	return act
}

// 根据XX查询
func (art *ModelPrice) QueryWhereOpenid(Openid string) error {
	if err := global.DB.Where("openid = ?", Openid).First(&art).Error; err != nil {
		return err
	}
	return nil
}

// 新增
func (art *ModelPrice) Create() error {
	if err := global.DB.Create(&art).Error; err != nil {
		return err
	}
	return nil
}

// 删除
func (art *ModelPrice) Delete(openid string) error {
	if err := global.DB.Where("openid = ?", openid).Delete(&art).Error; err != nil {
		return err
	}
	return nil
}

// 更新
func (art *ModelPrice) Update(openid string, data ModelPrice) error {
	if err := global.DB.Model(art).Where("openid = ?", openid).Updates(data).Error; err != nil {
		return err
	}
	return nil
}
