package department

import (
	"fmt"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/department"
)

type Department struct {
	Department srv_tenant.Department
	Column     []string
}

func Build() *gorm.DB {
	return DB.DB().Model(&srv_tenant.Department{})
}

func (m *Department) GetTree(req *srv_tenant.GetDepartmentTreeRequest) (departments []*srv_tenant.Department, err error) {
	sql := Build()

	if req.TenantId != 0 {
		sql = sql.Where("tenant_id = ?", req.TenantId)
	}

	if err := sql.Find(&departments).Error; err != nil {
		return nil, err
	}

	return departments, nil
}

func (m *Department) Store(req *srv_tenant.StoreDepartmentRequest) (department *srv_tenant.Department, err error) {
	department = req.Form
	department.MnemonicCode = helper.SliceCapitalize(helper.GetPinYin(department.Name))
	now := helper.Time2String(time.Now())
	department.Enable = true
	department.CreatedAt = now
	department.UpdatedAt = now

	if err := Build().Create(department).Error; err != nil {
		return nil, err
	}

	return department, nil
}

func (m *Department) Info(req *srv_tenant.GetDepartmentInfoRequest) (info *srv_tenant.Department, err error) {
	var department = srv_tenant.Department{}
	if err := Build().First(&department, req.Id).Error; err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	if department.Pid < 1 {
		department.Pid = -1
	}
	return &department, nil
}

func (m *Department) Update(req *srv_tenant.UpdateDepartmentRequest) (department *srv_tenant.Department, err error) {
	department = req.Form
	department.UpdatedAt = helper.Time2String(time.Now())
	if err := Build().Where("id = ?", req.Form.Id).Update(department).Error; err != nil {
		return nil, err
	}

	return department, nil
}

func (m *Department) BatchBanDepartment(req *srv_tenant.BatchBanDepartmentRequest) (err error) {
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
