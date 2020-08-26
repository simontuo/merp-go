package role

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
	srv_role "github.com/simontuo/merp-go/srv/permissionservice/proto/role"
)

type Role struct {
	Role   srv_role.Role
	Column []string
}

func Build() *gorm.DB {
	return DB.DB().Model(&srv_role.Role{})
}

func (m *Role) GetPage(req *srv_role.GetRolePageRequest) (roles []*srv_role.Role, count int32, err error) {
	sql := Build()
	if req.Name != "" {
		sql = sql.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := sql.Offset(req.PageSize * (req.CurrentPage - 1)).Limit(req.PageSize).Find(&roles).Error; err != nil {
		return nil, 0, err
	}
	var total int32
	sql.Count(&total)

	return roles, total, nil
}

func (m *Role) GetInfo(req *srv_role.GetRoleInfoRequest) (info *srv_role.RoleInfo, err error) {
	var role srv_role.Role
	if err := Build().First(&role, req.Id).Error; err != nil {
		return nil, err
	}

	var roleUsers []*srv_role.RoleUser
	if err := DB.DB().Model(srv_role.RoleUser{}).Where("role_id = ?", role.Id).Find(&roleUsers).Error; err != nil {
		return nil, err
	}

	var userIds []int32
	for _, v := range roleUsers {
		userIds = append(userIds, v.UserId)
	}

	var permissionRoles []*srv_permission.PermissionRole
	if err := DB.DB().Model(srv_permission.PermissionRole{}).Where("role_id = ?", role.Id).Find(&permissionRoles).Error; err != nil {
		return nil, err
	}
	var permissionId []int32
	for _, v := range permissionRoles {
		permissionId = append(permissionId, v.PermissionId)
	}

	i := srv_role.RoleInfo{
		Id:            role.Id,
		Name:          role.Name,
		Label:         role.Label,
		Desc:          role.Desc,
		Enable:        role.Enable,
		PermissionIds: permissionId,
		UserIds:       userIds,
	}

	return &i, nil
}

func (m *Role) BatchBan(req *srv_role.BatchBanRoleRequest) (err error) {
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

func (m *Role) Store(req *srv_role.StoreRoleRequest) (r *srv_role.Role, err error) {
	var role srv_role.Role
	role.Name = req.Name
	role.Label = req.Label
	role.Desc = req.Desc
	role.Enable = true
	now := helper.Time2String(time.Now())
	role.CreatedAt = now
	role.UpdatedAt = now

	if err := Build().Create(&role).Error; err != nil {
		return nil, err
	}

	for _, v := range req.PermissionIds {
		permissionRole := srv_permission.PermissionRole{
			RoleId:       role.Id,
			PermissionId: v,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		if err := DB.DB().Model(&srv_permission.PermissionRole{}).Create(&permissionRole).Error; err != nil {
			break
		}
	}

	return &role, nil
}

func (m *Role) BindUser(req *srv_role.BindUserRequest) (err error) {
	if err := DB.DB().Model(&srv_role.RoleUser{}).Where("role_id = ?", req.Id).Delete(&srv_role.RoleUser{}).Error; err != nil {
		return err
	}

	now := helper.Time2String(time.Now())
	for _, v := range req.UserIds {
		roleUser := srv_role.RoleUser{
			UserId:    v,
			RoleId:    req.Id,
			CreatedAt: now,
			UpdatedAt: now,
		}
		if err := DB.DB().Model(&srv_role.RoleUser{}).Create(&roleUser).Error; err != nil {
			break
		}
	}

	return nil
}

func (m *Role) Update(req *srv_role.UpdateRoleRequest) (r *srv_role.Role, err error) {
	var role srv_role.Role
	role.Name = req.Name
	role.Label = req.Label
	role.Desc = req.Desc
	now := helper.Time2String(time.Now())
	role.UpdatedAt = now
	if err := Build().Where("id = ?", req.Id).Update(&role).Error; err != nil {
		return nil, err
	}

	// 删除权限关系
	if err := DB.DB().Model(&srv_permission.PermissionRole{}).Where("role_id = ?", req.Id).Delete(&srv_permission.PermissionRole{}).Error; err != nil {
		return nil, err
	}

	// 插入权限关系
	for _, v := range req.PermissionIds {
		permissionRole := srv_permission.PermissionRole{
			RoleId:       role.Id,
			PermissionId: v,
			CreatedAt:    now,
			UpdatedAt:    now,
		}

		if err := DB.DB().Model(&srv_permission.PermissionRole{}).Create(&permissionRole).Error; err != nil {
			break
		}
	}

	return &role, nil
}
