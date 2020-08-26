# merp-go
### 架构
![avatar](http://photo.maguas.com/%E6%9E%B6%E6%9E%84.jpg)
### 目录结构
```
.
├── config // 配置文件
│   └── file
├── gateway // 自定义网关
│   └── build
├── helper // 帮助函数
│   ├── err
│   ├── helper
│   ├── password
│   ├── response
│   ├── rsp
│   ├── token
│   └── tree
├── lib // lib
│   ├── DB
│   ├── jwt
│   ├── redis
│   ├── registry
│   ├── sms
│   └── tracer
├── plugin // 插件
│   ├── auth
│   ├── tracer
│   │   ├── gateway
│   │   └── web
│   └── util
│       ├── http
│       └── response
├── srv // 基础服务
│   ├── logservice
│   │   ├── bin
│   │   ├── handler
│   │   ├── proto
│   │   └── subscriber
│   ├── permissionservice
│   │   ├── bin
│   │   ├── build
│   │   ├── handler
│   │   ├── model
│   │   └── proto
│   ├── smsservice
│   │   ├── build
│   │   ├── handler
│   │   ├── proto
│   │   └── subscriber
│   ├── tenantservice
│   │   ├── build
│   │   ├── handler
│   │   ├── model
│   │   └── proto
│   └── userservice
│       ├── build
│       ├── handler
│       ├── model
│       ├── proto
│       └── subscriber
└── web // web聚合服务
    └── admin
        ├── build
        ├── handler
        ├── middleware
        ├── request
        └── router
        
```
### 目标功能
#### 自定义网关
 - [x] JWT认证
 - [x] 限流熔断
 - [ ] 链路追踪
 - [ ] 流量染色
#### web服务聚合
- [x] 后台admin
  - [ ] 数据仪表
  - [x] 用户管理
  - [x] 租户管理
  - [x] 权限管理
  - [ ] 服务管理
  - [ ] 事务管理
  - [ ] 日志管理
- [ ] 前台app
  - [ ] 数据仪表
  - [ ] 客户管理
  - [ ] 订单管理
  - [ ] 商务管理
  - [ ] 财务管理
  - [ ] 数据统计
  - [ ] 系统管理
#### RPC聚合服务
#### 基础服务
 - [x] 用户服务
   - [x] 新增用户
   - [x] 分页搜索
   - [x] 更新用户
   - [x] 用户详情
   - [x] 批量禁用/启用
 - [x] 租户服务
   - [x] 新增租户
   - [x] 分页搜索
   - [x] 更新租户
   - [x] 租户详情
   - [x] 租户批量禁用/启用
   - [x] 新增部门
   - [x] 部门树列表
   - [x] 更新部门
   - [x] 部门详情
   - [x] 部门批量禁用/启用
 - [x] 权限服务
   - [x] 新增权限
   - [x] 权限分页搜索
   - [x] 更新权限
   - [x] 权限详情
   - [x] 权限批量禁用/启用
   - [x] 新增角色
   - [x] 角色分页搜索
   - [x] 更新角色
   - [x] 角色详情
   - [x] 角色批量禁用/启用
 - [x] 短信服务
   - [x] 登录验证码
   - [x] 重置密码验证码
 - [ ] 日志服务
 


  
