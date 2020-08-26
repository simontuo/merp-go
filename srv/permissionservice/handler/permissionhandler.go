package handler

import (
	"context"

	"github.com/micro/go-micro/v2/errors"
	"github.com/simontuo/merp-go/srv/permissionservice/model/permission"

	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"
)

type PermissionHandler struct {
}

var ErrId string

func Init(name string) {
	ErrId = name
}

func (h *PermissionHandler) GetPermissionPage(ctx context.Context, req *srv_permission.GetPermissionPageRequest, rsp *srv_permission.GetPermissionPageResponse) error {
	m := permission.Permission{}
	permissions, total, err := m.GetPage(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = permissions
	rsp.Total = total
	return nil
}

func (h *PermissionHandler) GetPermissionInfo(ctx context.Context, req *srv_permission.GetPermissionInfoRequest, rsp *srv_permission.GetPermissionInfoResponse) error {
	m := permission.Permission{}
	info, err := m.Info(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}
	rsp.Data = info

	return nil
}

func (h *PermissionHandler) StorePermission(ctx context.Context, req *srv_permission.StorePermissionRequest, rsp *srv_permission.StorePermissionResponse) error {
	m := permission.Permission{}
	p, err := m.Store(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = p
	return nil
}

func (h *PermissionHandler) UpdatePermission(ctx context.Context, req *srv_permission.UpdatePermissionRequest, rsp *srv_permission.UpdatePermissionResponse) error {
	m := permission.Permission{}
	p, err := m.Update(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = p
	return nil
}

func (h *PermissionHandler) GetPermissionTree(ctx context.Context, req *srv_permission.GetPermissionTreeRequest, rsp *srv_permission.GetPermissionTreeResponse) error {
	m := permission.Permission{}
	permissions, err := m.GetTree(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = permissions
	return nil
}

func (h *PermissionHandler) BatchBanPermission(ctx context.Context, req *srv_permission.BatchBanPermissionRequest, rsp *srv_permission.BatchBanPermissionResponse) error {
	m := permission.Permission{}

	if err := m.BatchBanOrEnable(req); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"
	return nil
}
