package tenantuser

import (
	"time"

	"github.com/jinzhu/gorm"
	"github.com/simontuo/merp-go/helper/helper"
	"github.com/simontuo/merp-go/lib/DB"
	srv_tenantuser "github.com/simontuo/merp-go/srv/tenantservice/proto/tenantuser"
)

type TenantUser struct {
	TenantUser *srv_tenantuser.TenantUser
	Column     []string
}

func Build() *gorm.DB {
	return DB.DB().Model(&srv_tenantuser.TenantUser{})
}

func (m *TenantUser) Store(userId int32, tenantId int32) (tenantUser *srv_tenantuser.TenantUser, err error) {
	now := helper.Time2String(time.Now())
	tu := srv_tenantuser.TenantUser{
		UserId:    userId,
		TenantId:  tenantId,
		UpdatedAt: now,
		CreatedAt: now,
	}

	if err := Build().Create(&tu).Error; err != nil {
		return nil, err
	}

	return &tu, nil
}
