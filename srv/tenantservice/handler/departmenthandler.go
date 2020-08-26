package handler

import (
	"context"

	"github.com/micro/go-micro/errors"

	"github.com/simontuo/merp-go/srv/tenantservice/model/department"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/department"
)

type DepartmentHandler struct {
}

var ErrId string

func Init(name string) {
	ErrId = name
}

func (h *DepartmentHandler) GetDepartmentTree(ctx context.Context, req *srv_tenant.GetDepartmentTreeRequest, rsp *srv_tenant.GetDepartmentTreeResponse) error {
	m := department.Department{}

	departments, err := m.GetTree(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = departments

	return nil
}

func (h *DepartmentHandler) GetDepartmentInfo(ctx context.Context, req *srv_tenant.GetDepartmentInfoRequest, rsp *srv_tenant.GetDepartmentInfoResponse) error {
	m := department.Department{}
	d, err := m.Info(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = d

	return nil
}

func (h *DepartmentHandler) StoreDepartment(ctx context.Context, req *srv_tenant.StoreDepartmentRequest, rsp *srv_tenant.StoreDepartmentResponse) error {
	m := department.Department{}
	d, err := m.Store(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = d

	return nil
}

func (h *DepartmentHandler) UpdateDepartment(ctx context.Context, req *srv_tenant.UpdateDepartmentRequest, rsp *srv_tenant.UpdateDepartmentResponse) error {
	m := department.Department{}
	d, err := m.Update(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = d

	return nil
}

func (h *DepartmentHandler) BatchBanDepartment(ctx context.Context, req *srv_tenant.BatchBanDepartmentRequest, rsp *srv_tenant.BatchBanDepartmentResponse) error {
	m := department.Department{}
	if err := m.BatchBanDepartment(req); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"
	return nil
}
