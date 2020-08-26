package tree

import (
	"sort"

	srv_permission "github.com/simontuo/merp-go/srv/permissionservice/proto/permission"

	srv_tenant "github.com/simontuo/merp-go/srv/tenantservice/proto/department"
)

type DepartmentTreeNode struct {
	ID       int32                 `json:"id"`
	Pid      int32                 `json:"pid"`
	Label    string                `json:"label"`
	TenantId int32                 `json:"tenant_id"`
	Desc     string                `json:"desc"`
	Disabled bool                  `json:"disabled"`
	Children []*DepartmentTreeNode `json:"children"`
}

func BuildDepartmentTree(data []*srv_tenant.Department) map[int32]map[int32]*DepartmentTreeNode {

	var treeNodes []*DepartmentTreeNode

	for _, v := range data {
		if v.Pid == 0 {
			v.Pid = -1
		}
		node := &DepartmentTreeNode{
			ID:       v.Id,
			Pid:      v.Pid,
			Label:    v.Name,
			TenantId: v.TenantId,
			Desc:     v.Desc,
			Disabled: !v.Enable,
			Children: []*DepartmentTreeNode{},
		}
		treeNodes = append(treeNodes, node)
	}
	tree := make(map[int32]map[int32]*DepartmentTreeNode)
	for _, v := range treeNodes {
		id := v.ID
		pid := v.Pid
		if _, ok := tree[pid]; !ok { // 如果不存在则创建一个新节点
			tree[pid] = make(map[int32]*DepartmentTreeNode)
		}
		tree[pid][id] = v
	}
	return tree
}

// 递归构造结构
func BuildTree(pid int32, data map[int32]map[int32]*DepartmentTreeNode) []*DepartmentTreeNode {
	node := make([]*DepartmentTreeNode, 0)

	for id, item := range data[pid] {
		if data[id] != nil {
			item.Children = BuildTree(id, data)
		}
		node = append(node, item)
	}

	// 升序排序
	sort.Slice(node, func(i, j int) bool {
		return node[i].ID < node[j].ID
	})

	return node
}

type PermissionTreeNode struct {
	ID       int32                 `json:"id"`
	Pid      int32                 `json:"pid"`
	Label    string                `json:"label"`
	TenantId int32                 `json:"tenant_id"`
	Desc     string                `json:"desc"`
	Disabled bool                  `json:"disabled"`
	Children []*PermissionTreeNode `json:"children"`
}

func BuildPermissionTree(data []*srv_permission.Permission) map[int32]map[int32]*PermissionTreeNode {

	var treeNodes []*PermissionTreeNode

	for _, v := range data {
		if v.Pid == 0 {
			v.Pid = -1
		}
		node := &PermissionTreeNode{
			ID:       v.Id,
			Pid:      v.Pid,
			Label:    v.Name,
			Desc:     v.Desc,
			Disabled: !v.Enable,
			Children: []*PermissionTreeNode{},
		}
		treeNodes = append(treeNodes, node)
	}
	tree := make(map[int32]map[int32]*PermissionTreeNode)
	for _, v := range treeNodes {
		id := v.ID
		pid := v.Pid
		if _, ok := tree[pid]; !ok { // 如果不存在则创建一个新节点
			tree[pid] = make(map[int32]*PermissionTreeNode)
		}
		tree[pid][id] = v
	}
	return tree
}

// 递归构造结构
func BuildPTree(pid int32, data map[int32]map[int32]*PermissionTreeNode) []*PermissionTreeNode {
	node := make([]*PermissionTreeNode, 0)

	for id, item := range data[pid] {
		if data[id] != nil {
			item.Children = BuildPTree(id, data)
		}
		node = append(node, item)
	}

	// 升序排序
	sort.Slice(node, func(i, j int) bool {
		return node[i].ID < node[j].ID
	})

	return node
}
