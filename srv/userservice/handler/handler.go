package handler

import (
	"context"

	"github.com/simontuo/merp-go/lib/DB"

	"github.com/simontuo/merp-go/helper/password"

	"github.com/micro/go-micro/v2/errors"
	"github.com/simontuo/merp-go/srv/userservice/model"

	srv_user "github.com/simontuo/merp-go/srv/userservice/proto/user"
)

type Handler struct {
}

func (h *Handler) ResetPassword(ctx context.Context, req *srv_user.ResetPasswordRequest, rsp *srv_user.ResetPasswordResponse) error {
	m := model.User{}
	m.User.Phone = req.Phone
	user, err := m.GetInfo()
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	restUser := srv_user.User{
		Password: password.Hash(req.Password),
	}
	if err := DB.DB().Model(&srv_user.User{}).Where("id = ?", user.Id).Update(&restUser).Error; err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"

	return nil
}

func (h *Handler) GetAdminUserList(ctx context.Context, req *srv_user.GetAdminUserListRequest, rsp *srv_user.GetAdminUserListResponse) error {
	m := model.AdminUser{}
	adminUsers, err := m.GetAdminUserList(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = adminUsers

	return nil
}

var ErrId string

func Init(name string) {
	ErrId = name
}

func (h *Handler) GetUserInfo(ctx context.Context, req *srv_user.GetUserInfoRequest, rsp *srv_user.GetUserResponse) error {
	m := model.User{}
	m.User.Id = req.Id
	m.User.Phone = req.Phone
	user, err := m.GetInfo()
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = user
	return nil
}

func (h *Handler) VerifyUser(ctx context.Context, req *srv_user.VerifyUserRequest, rsp *srv_user.VerifyUserResponse) error {
	var user srv_user.User
	if err := DB.DB().Model(&srv_user.User{}).Where("phone = ?", req.Phone).First(&user).Error; err != nil {
		return errors.Unauthorized(ErrId, "不存在的用户")
	}
	m := model.User{}
	m.User.Id = user.Id

	isAdmin, err := m.IsAdmin()
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}
	if !isAdmin {
		return errors.Unauthorized(ErrId, "该用户不是后台管理员")
	}

	rsp.Pass = false
	if password.Verify(req.Password, user.Password) {
		rsp.Pass = true
		rsp.Id = user.Id
	}

	return nil
}

func (h *Handler) GetUserPage(ctx context.Context, req *srv_user.GetUserPageRequest, rsp *srv_user.GetUserPageResponse) error {
	m := model.User{}

	users, total, err := m.GetPage(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = users
	rsp.Total = total
	return nil
}

func (h *Handler) StoreUser(ctx context.Context, req *srv_user.StoreUserRequest, rsp *srv_user.StoreUserResponse) error {
	m := model.User{}

	user, err := m.Store(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = user

	return nil
}

func (h *Handler) BatchBanUser(ctx context.Context, req *srv_user.BatchBanUserRequest, rsp *srv_user.BatchBanUserResponse) error {
	m := model.User{}

	if err := m.BatchBan(req); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Status = "success"

	return nil
}

func (h *Handler) UpdateUser(ctx context.Context, req *srv_user.UpdateUserRequest, rsp *srv_user.UpdateUserResponse) error {
	panic("implement me")
}

func (h *Handler) GetUserList(ctx context.Context, req *srv_user.GetUserListRequest, rsp *srv_user.GetUserListResponse) error {
	m := model.User{}

	users, err := m.GetList(req)
	if err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	rsp.Data = users

	return nil
}

func (h *Handler) BindAdmin(ctx context.Context, req *srv_user.BindAdminRequest, rsp *srv_user.BindAdminResponse) error {
	m := model.AdminUser{}
	if err := m.DelAllAdminUsers(); err != nil {
		return errors.InternalServerError(ErrId, err.Error())
	}

	for _, v := range req.Ids {
		m := model.AdminUser{}
		if _, err := m.AdminUserStore(v); err != nil {
			break
		}
	}

	rsp.Status = "success"
	return nil
}
