package handler

import (
	"context"

	"github.com/micro/go-micro/errors"
	"github.com/simontuo/merp-go/srv/tenantservice/model/tenantuser"

	"github.com/simontuo/merp-go/srv/tenantservice/model/tenant"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/tenant"
)

type TenantHandler struct {
}

func (h *TenantHandler) GetTenantInfo(ctx context.Context, req *srv_tenant.GetTenantInfoRequest, rsp *srv_tenant.GetTenantInfoResponse) error {
	m := tenant.Tenant{}
	t, err := m.Info(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = t

	return nil
}

func (h *TenantHandler) GetTenantPage(ctx context.Context, req *srv_tenant.GetTenantPageRequest, rsp *srv_tenant.GetTenantPageResponse) error {
	m := tenant.Tenant{}

	tenants, total, err := m.GetPage(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = tenants
	rsp.Total = total

	return nil
}

func (h *TenantHandler) StoreTenant(ctx context.Context, req *srv_tenant.StoreTenantRequest, rsp *srv_tenant.StoreTenantResponse) error {
	// todo 事务
	m := tenant.Tenant{}
	t, err := m.Store(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}
	_, err = m.StoreDepartment(t)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	tu := tenantuser.TenantUser{}
	if _, err := tu.Store(t.AdministratorId, t.Id); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = t

	return nil
}

func (h *TenantHandler) UpdateTenant(ctx context.Context, req *srv_tenant.UpdateTenantRequest, rsp *srv_tenant.UpdateTenantResponse) error {
	m := tenant.Tenant{}

	t, err := m.Update(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = t
	return nil
}
