package operation

import "github.com/wlgd/xutils/orm"

// Where 分页
type Where struct {
	orm.DbPage
	DeviceNo     string `form:"deviceNo"`
	OrganizeId   *int   `form:"organizeId"` // 每页数
	OrganizeGUID string `form:"-"`
}

// Where 初始化
func (o *Where) Vehicle() *orm.DbWhere {
	where := o.DbWhere()
	where.Equal("device_no", o.DeviceNo)
	where.Equal("organized_id", o.OrganizeId)
	where.Equal("organized_guid", o.OrganizeGUID)
	return where
}
