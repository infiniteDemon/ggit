package models

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"service-all/app/init/global"
)

type ModelAdminUser struct {
	gorm.Model
	// 用户名
	UserName string `json:"username" gorm:"unique"`
	// 密码
	Password string `json:"password"`
	// 权限 -1 未激活/已禁用 0 开启并标准化用户 1 管理员用户 2 超级用户
	Status int `json:"status"`
	// 店铺名称
	Name string `json:"name"`
	// 店铺地址
	Site string `json:"site"`
	// 店铺描述
	Desc string `json:"desc"`
	// 是否开业 1 开业 -1 不开业 0 歇业
	InOperation int `json:"in_operation"`
}

type RequestAdminLoginData struct {
	// 姓名
	UserName string `json:"username" validate:"required" example:"huanghua"`
	// 密码
	Password string `json:"password" validate:"required" example:"test123"`
	// 验证码
	Captcha string `json:"captcha" validate:"required" example:"9782"`
}

type RequestAdminAddUserData struct {
	// 姓名
	UserName string `json:"username" validate:"required" example:"huanghua"`
	// 初始密码
	Password string `json:"password" validate:"required" example:"test123"`
	// 权限 -1 未激活/已禁用 0 开启并标准化用户 1 管理员用户 2 超级用户
	Status int `json:"status"`
	// 店铺名称
	Name string `json:"name" validate:"required" example:"副食店"`
	// 店铺地址
	Site string `json:"site" validate:"required" example:"白云小区"`
	// 店铺描述
	Desc string `json:"desc" validate:"required" example:"专属副食"`
	// 是否开业 1 开业 -1 不开业 0 歇业
	InOperation int `json:"in_operation" validate:"required" example:1`
}

type ActionModelAdmin interface {
	InitAdminUser() error
	QueryWhereUserName(userName string) error
	Create() error
	//QueryAllOffsetLimit() ([]ModelAdminUser, error)
	//Init() error
	//SelectAllAndInsert() (int, error)
}

func ActAdmin(act ActionModelAdmin) ActionModelAdmin {
	return act
}

// 初始化超级管理员
func (art *ModelAdminUser) InitAdminUser() error {
	if err := global.DB.Clauses(clause.OnConflict{UpdateAll: true}).Create(&art).Error; err != nil {
		return err
	}
	return nil
}

// 根据标题和cid mid查询 可以用来去重
func (art *ModelAdminUser) QueryWhereUserName(userName string) error {
	if err := global.DB.Where("user_name = ?", userName).First(&art).Error; err != nil {
		return err
	}
	return nil
}

// 新增文章
func (art *ModelAdminUser) Create() error {
	if err := global.DB.Create(&art).Error; err != nil {
		return err
	}
	return nil
}

// 新增文章
//func (art *ModelAdminUser) Init() error {
//	if err := global.DB.Where("id > 0").Delete(&art).Error; err != nil {
//		return err
//	}
//	return nil
//}
//
//// 根据mid查询所有分类下的文章
//func (art *ModelAdminUser) QueryAllOffsetLimit() ([]ModelAdminUser, error) {
//	records := []ModelAdminUser{}
//	if err := global.DB.Order("updated_at DESC").Find(&records).Error; err != nil {
//		return nil, err
//	}
//	return records, nil
//}
//
//func (art *ModelAdminUser) SelectAllAndInsert() (int, error) {
//	records := []ModelAdminUser{}
//	var success_num int
//	if err := global.DB.Find(&records).Error; err != nil {
//		return success_num, err
//	}
//
//	global.LOG.Info("备份数据查询成功，开始写入远程数据库")
//
//	for _, v := range records {
//		if err1 := global.BackupDB.Create(&v).Error; err1 != nil {
//			if err2 := global.BackupDB.Where("id = ?", v.ID).Updates(&v).Error; err2 != nil {
//				global.LOG.Error("备份更新出错", zap.Error(err2))
//				return success_num, err2
//			}
//		}
//		success_num++
//	}
//	return success_num, nil
//}
