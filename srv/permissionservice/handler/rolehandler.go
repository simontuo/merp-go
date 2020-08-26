package handler

import (
	"context"

	"github.com/micro/go-micro/errors"
	"github.com/simontuo/merp-go/srv/permissionservice/model/role"

	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/role"
)

type RoleHandler struct {
}

func (h *RoleHandler) BatchBanRole(ctx context.Context, req *srv_permission.BatchBanRoleRequest, rsp *srv_permission.BatchBanRoleResponse) error {
	m := role.Role{}
	if err := m.BatchBan(req); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	return nil
}

func (h *RoleHandler) GetRolePage(ctx context.Context, req *srv_permission.GetRolePageRequest, rsp *srv_permission.GetRolePageResponse) error {
	m := role.Role{}
	roles, total, err := m.GetPage(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = roles
	rsp.Total = total

	return nil
}

func (h *RoleHandler) StoreRole(ctx context.Context, req *srv_permission.StoreRoleRequest, rsp *srv_permission.StoreRoleResponse) error {
	m := role.Role{}

	r, err := m.Store(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = r
	return nil
}

func (h *RoleHandler) GetRoleInfo(ctx context.Context, req *srv_permission.GetRoleInfoRequest, rsp *srv_permission.GetRoleInfoResponse) error {
	m := role.Role{}

	info, err := m.GetInfo(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = info
	return nil
}

func (h *RoleHandler) UpdateRole(ctx context.Context, req *srv_permission.UpdateRoleRequest, rsp *srv_permission.UpdateRoleResponse) error {
	m := role.Role{}
	r, err := m.Update(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}
	roleInfo := srv_permission.RoleInfo{
		Id:            r.Id,
		Name:          r.Name,
		Label:         r.Label,
		Desc:          r.Desc,
		PermissionIds: req.PermissionIds,
	}
	rsp.Data = &roleInfo
	return nil
}

func (h *RoleHandler) BindUser(ctx context.Context, req *srv_permission.BindUserRequest, rsp *srv_permission.BindUserResponse) error {
	m := role.Role{}

	if err := m.BindUser(req); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"

	return nil
}
