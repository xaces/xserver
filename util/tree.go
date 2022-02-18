package util

import "reflect"

// TreeNode 树节点
type TreeNode struct {
	Id       uint64     `json:"id"`     // ID
	Label    string     `json:"label"`  // 显示标签
	Status   string     `json:"status"` // 状态
	Children []TreeNode `json:"children,omitempty"`
}

type treeTag struct {
	Id     string
	Label  string
	Status string
}

type tree struct {
	Tag treeTag
}

func valToUint64(data reflect.Value, name string) uint64 {
	return data.FieldByName(name).Uint()
}

func NewTree(id, lable, status string, data interface{}) []TreeNode {
	tag := treeTag{Id: id, Label: lable, Status: status}
	t := tree{Tag: tag}
	return t.build(data)
}

// 菜单Tree
func (t *tree) build(data interface{}) []TreeNode {
	var trees []TreeNode
	s := reflect.ValueOf(data)
	switch s.Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < s.Len(); i++ {
			v := s.Index(i)
			pId := valToUint64(v, "ParentId")
			if t.hasParent(s, pId) {
				continue
			}
			node := t.newNode(v)
			t.filterChildren(s, node)
			trees = append(trees, node)
		}
	}
	return trees
}

// newNode 创建数节点
func (t *tree) newNode(v reflect.Value) TreeNode {
	id := valToUint64(v, t.Tag.Id)
	label := v.FieldByName(t.Tag.Label).String()
	status := v.FieldByName(t.Tag.Status).String()
	return TreeNode{id, label, status, nil}
}

// hasParent 判断权限列表中是否有父节点
func (t *tree) hasParent(data reflect.Value, parentId uint64) bool {
	for i := 0; i < data.Len(); i++ {
		id := valToUint64(data.Index(i), t.Tag.Id)
		if id != parentId {
			continue
		}
		return true
	}
	return false
}

// filterChildren 获取子节点
func (t *tree) filterChildren(data reflect.Value, n TreeNode) {
	for i := 0; i < data.Len(); i++ {
		v := data.Index(i)
		pId := valToUint64(v, "ParentId")
		if pId != n.Id {
			continue
		}
		node := t.newNode(v)
		t.filterChildren(data, node)
		n.Children = append(n.Children, node)
	}
}
