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
