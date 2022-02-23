package util

import "reflect"

// TreeNode 树节点
type TreeNode struct {
	Id       uint64     `json:"id"`     // ID
	Label    string     `json:"label"`  // 显示标签
	Status   string     `json:"status"` // 状态
	Children []TreeNode `json:"children,omitempty"`
}

type TreeTag struct {
	Id     string
	Label  string
	Status string
}

type Tree struct {
	Tag TreeTag
}

func valToUint64(data reflect.Value, name string) uint64 {
	return data.FieldByName(name).Uint()
}

func NewTree(id, lable, status string) *Tree {
	t := Tree{Tag: TreeTag{Id: id, Label: lable, Status: status}}
	return &t
}

// 菜单Tree
func (t *Tree) Build(data interface{}) []TreeNode {
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
			t.filterChildren(s, "ParentId", node)
			trees = append(trees, node)
		}
	}
	return trees
}

// newNode 创建数节点
func (t *Tree) newNode(v reflect.Value) TreeNode {
	id := valToUint64(v, t.Tag.Id)
	label := v.FieldByName(t.Tag.Label).String()
	status := v.FieldByName(t.Tag.Status).String()
	return TreeNode{id, label, status, nil}
}

// hasParent 判断权限列表中是否有父节点
func (t *Tree) hasParent(data reflect.Value, parentId uint64) bool {
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
func (t *Tree) filterChildren(data reflect.Value, parentTag string, n TreeNode) {
	for i := 0; i < data.Len(); i++ {
		v := data.Index(i)
		pId := valToUint64(v, parentTag)
		if pId != n.Id {
			continue
		}
		node := t.newNode(v)
		t.filterChildren(data, parentTag, node)
		n.Children = append(n.Children, node)
	}
}
