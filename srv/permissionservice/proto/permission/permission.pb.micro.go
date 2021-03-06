// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/permission/permission.proto

package srv_permission

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for PermissionService service

func NewPermissionServiceEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for PermissionService service

type PermissionService interface {
	GetPermissionPage(ctx context.Context, in *GetPermissionPageRequest, opts ...client.CallOption) (*GetPermissionPageResponse, error)
	GetPermissionInfo(ctx context.Context, in *GetPermissionInfoRequest, opts ...client.CallOption) (*GetPermissionInfoResponse, error)
	StorePermission(ctx context.Context, in *StorePermissionRequest, opts ...client.CallOption) (*StorePermissionResponse, error)
	UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...client.CallOption) (*UpdatePermissionResponse, error)
	GetPermissionTree(ctx context.Context, in *GetPermissionTreeRequest, opts ...client.CallOption) (*GetPermissionTreeResponse, error)
	BatchBanPermission(ctx context.Context, in *BatchBanPermissionRequest, opts ...client.CallOption) (*BatchBanPermissionResponse, error)
}

type permissionService struct {
	c    client.Client
	name string
}

func NewPermissionService(name string, c client.Client) PermissionService {
	return &permissionService{
		c:    c,
		name: name,
	}
}

func (c *permissionService) GetPermissionPage(ctx context.Context, in *GetPermissionPageRequest, opts ...client.CallOption) (*GetPermissionPageResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GetPermissionPage", in)
	out := new(GetPermissionPageResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) GetPermissionInfo(ctx context.Context, in *GetPermissionInfoRequest, opts ...client.CallOption) (*GetPermissionInfoResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GetPermissionInfo", in)
	out := new(GetPermissionInfoResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) StorePermission(ctx context.Context, in *StorePermissionRequest, opts ...client.CallOption) (*StorePermissionResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.StorePermission", in)
	out := new(StorePermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, opts ...client.CallOption) (*UpdatePermissionResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.UpdatePermission", in)
	out := new(UpdatePermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) GetPermissionTree(ctx context.Context, in *GetPermissionTreeRequest, opts ...client.CallOption) (*GetPermissionTreeResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.GetPermissionTree", in)
	out := new(GetPermissionTreeResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *permissionService) BatchBanPermission(ctx context.Context, in *BatchBanPermissionRequest, opts ...client.CallOption) (*BatchBanPermissionResponse, error) {
	req := c.c.NewRequest(c.name, "PermissionService.BatchBanPermission", in)
	out := new(BatchBanPermissionResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for PermissionService service

type PermissionServiceHandler interface {
	GetPermissionPage(context.Context, *GetPermissionPageRequest, *GetPermissionPageResponse) error
	GetPermissionInfo(context.Context, *GetPermissionInfoRequest, *GetPermissionInfoResponse) error
	StorePermission(context.Context, *StorePermissionRequest, *StorePermissionResponse) error
	UpdatePermission(context.Context, *UpdatePermissionRequest, *UpdatePermissionResponse) error
	GetPermissionTree(context.Context, *GetPermissionTreeRequest, *GetPermissionTreeResponse) error
	BatchBanPermission(context.Context, *BatchBanPermissionRequest, *BatchBanPermissionResponse) error
}

func RegisterPermissionServiceHandler(s server.Server, hdlr PermissionServiceHandler, opts ...server.HandlerOption) error {
	type permissionService interface {
		GetPermissionPage(ctx context.Context, in *GetPermissionPageRequest, out *GetPermissionPageResponse) error
		GetPermissionInfo(ctx context.Context, in *GetPermissionInfoRequest, out *GetPermissionInfoResponse) error
		StorePermission(ctx context.Context, in *StorePermissionRequest, out *StorePermissionResponse) error
		UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, out *UpdatePermissionResponse) error
		GetPermissionTree(ctx context.Context, in *GetPermissionTreeRequest, out *GetPermissionTreeResponse) error
		BatchBanPermission(ctx context.Context, in *BatchBanPermissionRequest, out *BatchBanPermissionResponse) error
	}
	type PermissionService struct {
		permissionService
	}
	h := &permissionServiceHandler{hdlr}
	return s.Handle(s.NewHandler(&PermissionService{h}, opts...))
}

type permissionServiceHandler struct {
	PermissionServiceHandler
}

func (h *permissionServiceHandler) GetPermissionPage(ctx context.Context, in *GetPermissionPageRequest, out *GetPermissionPageResponse) error {
	return h.PermissionServiceHandler.GetPermissionPage(ctx, in, out)
}

func (h *permissionServiceHandler) GetPermissionInfo(ctx context.Context, in *GetPermissionInfoRequest, out *GetPermissionInfoResponse) error {
	return h.PermissionServiceHandler.GetPermissionInfo(ctx, in, out)
}

func (h *permissionServiceHandler) StorePermission(ctx context.Context, in *StorePermissionRequest, out *StorePermissionResponse) error {
	return h.PermissionServiceHandler.StorePermission(ctx, in, out)
}

func (h *permissionServiceHandler) UpdatePermission(ctx context.Context, in *UpdatePermissionRequest, out *UpdatePermissionResponse) error {
	return h.PermissionServiceHandler.UpdatePermission(ctx, in, out)
}

func (h *permissionServiceHandler) GetPermissionTree(ctx context.Context, in *GetPermissionTreeRequest, out *GetPermissionTreeResponse) error {
	return h.PermissionServiceHandler.GetPermissionTree(ctx, in, out)
}

func (h *permissionServiceHandler) BatchBanPermission(ctx context.Context, in *BatchBanPermissionRequest, out *BatchBanPermissionResponse) error {
	return h.PermissionServiceHandler.BatchBanPermission(ctx, in, out)
}
