package model

import (
	"time"

	"github.com/simontuo/merp-go/helper/helper"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/password"
	"github.com/simontuo/merp-go/lib/DB"
	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"
)

type User struct {
	User   srv_user.User
	Column []string
}

func BuildUser() *gorm.DB {
	return DB.DB().Model(&srv_user.User{})
}

func BuildSelectUser() *gorm.DB {
	model := &srv_user.User{}
	user := User{}
	user.Column = []string{"id", "name", "phone", "email", "enable", "created_at", "updated_at", "deleted_at"}
	return DB.DB().Model(model).Select(user.Column)
}

func (m *User) GetPage(req *srv_user.GetUserPageRequest) (users []*srv_user.User, count int32, err error) {
	sql := BuildSelectUser()
	if req.Name != "" {
		sql = sql.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := sql.Offset(req.PageSize * (req.CurrentPage - 1)).Limit(req.PageSize).Find(&users).Error; err != nil {
		return nil, 0, err
	}
	var total int32
	sql.Count(&total)

	return users, total, nil
}

func (m *User) GetInfo() (info *srv_user.User, err error) {
	sql := BuildSelectUser()
	if m.User.Id != 0 {
		sql = sql.Where("id = ?", m.User.Id)
	}
	if m.User.Phone != "" {
		sql = sql.Where("phone = ?", m.User.Phone)
	}

	var user srv_user.User

	if err := sql.First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (m *User) IsAdmin() (bool, error) {
	sql := BuildAdminUser()
	var count int
	if err := sql.Where("user_id = ?", m.User.Id).Count(&count).Error; err != nil {
		return false, err
	}

	if count < 1 {
		return false, nil
	} else {
		return true, nil
	}
}

func (m *User) Store(req *srv_user.StoreUserRequest) (user *srv_user.User, err error) {
	sql := BuildUser()
	user = req.Form
	user.Password = password.Hash(user.Password)
	user.Enable = true
	now := helper.Time2String(time.Now())
	user.CreatedAt = now
	user.UpdatedAt = now

	if err := sql.Create(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func (m *User) BatchBan(req *srv_user.BatchBanUserRequest) error {
	enable := make(map[string]interface{})
	if req.Type == "ban" {
		enable["enable"] = false
	} else {
		enable["enable"] = true
	}

	if err := BuildUser().Where("id in (?)", req.Ids).Updates(enable).Error; err != nil {
		return err
	}

	return nil
}

func (m *User) GetList(req *srv_user.GetUserListRequest) (users []*srv_user.User, err error) {
	sql := BuildUser()

	if req.Name != "" {
		sql = sql.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := sql.Find(&users).Error; err != nil {
		return nil, err
	}

	return users, nil
}
