package service

import (
	"xserver/model"

	"github.com/wlgd/xutils/orm"
)

// 主组织
func OprPrimaryOrganization(guid string) (model.OprOrganization, error) {
	var data model.OprOrganization
	err := orm.DB().Model(&data).Where("organize_guid = ? AND parent_id = 0", guid).Preload("SysStation").First(&data).Error
	return data, err
}

type OprVehicle struct {
	Id         int    `json:"deviceId"`
	DeviceNo   string `json:"deviceNo"`
	DeviceName string `json:"deviceName"`
	ChlCount   int    `json:"chlCount"`
	ChlNames   string `json:"chlNames"`
	OrganizeId int    `json:"organizeId"` // 分组Id
}

type OprOrgainze struct {
	Id       int           `json:"id"`   //
	Name     string        `json:"name"` // 名称
	ParentId int           `json:"parentId"`
	Vehicles []OprVehicle  `json:"vehis,omitempty" gorm:"-"`
	Children []OprOrgainze `json:"children,omitempty" gorm:"-"`
}

func (o *OprOrgainze) filterChildren(data []OprOrgainze, vehis []OprVehicle) {
	nid := o.Id
	if o.ParentId == 0 {
		nid = 0
	}
	for _, v := range vehis {
		if nid == v.OrganizeId {
			o.Vehicles = append(o.Vehicles, v)
		}
	}
	for _, v := range data {
		if v.ParentId != o.Id {
			continue
		}
		v.filterChildren(data, vehis)
		o.Children = append(o.Children, v)
	}
}

func OprOrganizeTree(guid string, vehis []OprVehicle) (tree []OprOrgainze) {
	var data []OprOrgainze
	orm.DB().Model(&model.OprOrganization{}).Find(&data, "organize_guid = ?", guid)
	for _, v := range data {
		if v.ParentId != 0 {
			continue
		}
		v.filterChildren(data, vehis)
		tree = append(tree, v)
	}
	return
}
