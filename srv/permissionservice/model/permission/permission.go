package permission

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
)

type Permission struct {
	Permission srv_permission.Permission
	Column     []string
}

func Build() *gorm.DB {
	return DB.DB().Model(&srv_permission.Permission{})
}

func (m *Permission) GetPage(req *srv_permission.GetPermissionPageRequest) (permissions []*srv_permission.Permission, count int32, err error) {
	sql := Build()
	if req.Name != "" {
		sql = sql.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := sql.Offset(req.PageSize * (req.CurrentPage - 1)).Limit(req.PageSize).Find(&permissions).Error; err != nil {
		return nil, 0, err
	}
	var total int32
	sql.Count(&total)

	return permissions, total, nil
}

func (m *Permission) Info(req *srv_permission.GetPermissionInfoRequest) (info *srv_permission.Permission, err error) {
	var permission srv_permission.Permission
	if err := Build().First(&permission, req.Id).Error; err != nil {
		return nil, err
	}

	return &permission, nil
}

func (m *Permission) GetTree(req *srv_permission.GetPermissionTreeRequest) (permissions []*srv_permission.Permission, err error) {
	sql := Build()
	if req.Pid != 0 {
		sql = sql.Where("pid = ?", req.Pid)
	}

	if err := sql.Find(&permissions).Error; err != nil {
		return nil, err
	}

	return permissions, nil
}

func (m *Permission) BatchBanOrEnable(req *srv_permission.BatchBanPermissionRequest) (err error) {
	enable := make(map[string]interface{})
	if req.Type == "ban" {
		enable["enable"] = false
	} else {
		enable["enable"] = true
	}

	if err := Build().Where("id in (?)", req.Ids).Updates(enable).Error; err != nil {
		return err
	}

	return nil
}

func (m *Permission) Update(req *srv_permission.UpdatePermissionRequest) (permission *srv_permission.Permission, err error) {
	permission = req.Form
	permission.UpdatedAt = helper.Time2String(time.Now())
	if err := Build().Where("id = ?", permission.Id).Update(permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}

func (m *Permission) Store(req *srv_permission.StorePermissionRequest) (permission *srv_permission.Permission, err error) {
	permission = req.Form
	now := helper.Time2String(time.Now())
	permission.UpdatedAt = now
	permission.CreatedAt = now
	permission.Enable = true
	if err := Build().Create(permission).Error; err != nil {
		return nil, err
	}

	return permission, nil
}
