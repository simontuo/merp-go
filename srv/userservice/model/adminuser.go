package model

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"
)

type AdminUser struct {
	AdminUser srv_user.AdminUser
	Column    []string
}

func BuildAdminUser() *gorm.DB {
	return DB.DB().Model(&srv_user.AdminUser{})
}

func (m *AdminUser) GetAdminUserList(req *srv_user.GetAdminUserListRequest) (adminUsers []*srv_user.AdminUser, err error) {
	sql := BuildAdminUser()
	if len(req.Ids) > 0 {
		sql = sql.Where("user_id in (?)", req.Ids)
	}

	if err := sql.Find(&adminUsers).Error; err != nil {
		return nil, err
	}

	return adminUsers, nil
}

func (m *AdminUser) DelAllAdminUsers() (err error) {
	if err := BuildAdminUser().Delete(&m).Error; err != nil {
		return err
	}

	return nil
}

func (m *AdminUser) AdminUserStore(userId int32) (ad *srv_user.AdminUser, err error) {
	var admin srv_user.AdminUser
	admin.UserId = userId
	now := helper.Time2String(time.Now())
	admin.CreatedAt = now
	admin.UpdatedAt = now

	if err := BuildAdminUser().Create(&admin).Error; err != nil {
		return nil, err
	}

	return &admin, nil
}
