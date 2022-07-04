package service

import (
	"xserver/model"

	"github.com/xaces/xutils/orm"
)

// 主组织
func OprPrimaryOrganization(guid string) (model.OprOrganization, error) {
	var data model.OprOrganization
	err := orm.DB().Model(&data).Where("guid = ? AND parent_id = 0", guid).First(&data).Error
	return data, err
}

type OprOrgainze struct {
	Id       int           `json:"id"`   //
	Name     string        `json:"name"` // 名称
	ParentId int           `json:"parentId"`
	Count    int           `json:"count" gorm:"-"`
	Vehicles []Vehicle     `json:"vehis,omitempty" gorm:"-"`
	Children []OprOrgainze `json:"children,omitempty" gorm:"-"`
}

func (o *OprOrgainze) filterChildren(data []OprOrgainze, vehis []Vehicle) {
	nid := o.Id
	if o.ParentId == 0 {
		nid = 0
	}
	for _, v := range vehis {
		if nid != v.OrganizeId {
			continue
		}
		o.Vehicles = append(o.Vehicles, v)
		o.Count++
	}
	for _, v := range data {
		if v.ParentId != o.Id {
			continue
		}
		v.filterChildren(data, vehis)
		o.Children = append(o.Children, v)
		o.Count += v.Count
	}
}

func OprOrganizeTree(guid string, vehis []Vehicle) (tree []OprOrgainze) {
	var data []OprOrgainze
	orm.DB().Model(&model.OprOrganization{}).Find(&data, "guid = ?", guid)
	for _, v := range data {
		if v.ParentId != 0 {
			continue
		}
		v.filterChildren(data, vehis)
		tree = append(tree, v)
	}
	return
}
