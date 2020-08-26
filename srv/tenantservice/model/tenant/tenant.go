package tenant

import (
	"time"

	srv_tenant2 "github.com/simontuo/merp-go/srv/tenantservice/proto/department"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/tenant"
)

type Tenant struct {
	tenant *srv_tenant.Tenant
	Column []string
}

func Build() *gorm.DB {
	return DB.DB().Model(&srv_tenant.Tenant{})
}

func (m *Tenant) GetPage(req *srv_tenant.GetTenantPageRequest) (tenants []*srv_tenant.Tenant, count int32, err error) {
	sql := Build()
	if req.Name != "" {
		sql = sql.Where("name LIKE ?", "%"+req.Name+"%")
	}

	if err := sql.Offset(req.PageSize * (req.CurrentPage - 1)).Limit(req.PageSize).Find(&tenants).Error; err != nil {
		return nil, 0, err
	}
	var total int32
	sql.Count(&total)

	return tenants, total, nil
}

func (m *Tenant) Store(req *srv_tenant.StoreTenantRequest) (tenant *srv_tenant.Tenant, err error) {
	tenant = req.Form
	tenant.MnemonicCode = helper.SliceCapitalize(helper.GetPinYin(tenant.Name))
	now := helper.Time2String(time.Now())
	tenant.UpdatedAt = now
	tenant.CreatedAt = now
	tenant.Enable = true

	if err := Build().Create(tenant).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (m *Tenant) Info(req *srv_tenant.GetTenantInfoRequest) (tenant *srv_tenant.Tenant, err error) {
	var t srv_tenant.Tenant
	if err := Build().First(&t, req.Id).Error; err != nil {
		return nil, err
	}

	return &t, nil
}

func (m *Tenant) Update(req *srv_tenant.UpdateTenantRequest) (tenant *srv_tenant.Tenant, err error) {
	// todo 时间转换
	tenant = req.Form
	now := helper.Time2String(time.Now())
	tenant.UpdatedAt = now
	enable := make(map[string]bool)
	enable["enable"] = req.Form.Enable

	if err := Build().Where("id = ?", tenant.Id).Update(tenant).Updates(enable).Error; err != nil {
		return nil, err
	}

	return tenant, nil
}

func (m *Tenant) StoreDepartment(t *srv_tenant.Tenant) (department *srv_tenant2.Department, err error) {
	now := helper.Time2String(time.Now())
	d := srv_tenant2.Department{
		Name:         t.Name,
		MnemonicCode: t.MnemonicCode,
		TenantId:     t.Id,
		UpdatedAt:    now,
		CreatedAt:    now,
		Enable:       true,
	}

	if err := DB.DB().Model(srv_tenant2.Department{}).Create(&d).Error; err != nil {
		return nil, err
	}

	return &d, nil
}
