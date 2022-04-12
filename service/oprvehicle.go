package service

import (
	"github.com/wlgd/xutils/orm"
)

// VehiclePage 分页
type VehiclePage struct {
	orm.DbPage
	DeviceNo   string `form:"deviceNo"`
	DeviceName string `form:"deviceName"`
	OrganizeId *int   `form:"organizeId"` // 每页数
}

// Where 初始化
func (s *VehiclePage) Where() *orm.DbWhere {
	where := s.DbWhere()
	where.String("device_no like ?", s.DeviceNo)
	where.String("device_name like ?", s.DeviceName)
	if s.OrganizeId != nil {
		where.Append("organize_id = ?", *s.OrganizeId)
	}
	return where
}
