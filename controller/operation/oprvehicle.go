package operation

import (
	"fmt"
	"strconv"
	"xserver/middleware"
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

type Vehicle struct {
}

func (o *Vehicle) ListHandler(c *gin.Context) {
	var p service.VehiclePage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	w := p.Where()
	w.Append("organize_guid = ?", middleware.GetUserToken(c).OrganizeGuid)
	var data []model.OprVehicle
	total, _ := orm.DbByWhere(&data, w).Find(&data)
	ctx.JSONOk().Write(gin.H{"total": total, "data": data}, c)
}

// GetHandler 获取指定id
func (o *Vehicle) GetHandler(c *gin.Context) {
	var data model.OprVehicle
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Vehicle) AddHandler(c *gin.Context) {
	var p model.OprVehicle
	//获取参数
	if err := c.ShouldBind(&p.OprVehicleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	p.Guid = util.UUID()
	p.OrganizeGuid = middleware.GetUserToken(c).OrganizeGuid
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

type batchAdd struct {
	Prefix      string `json:"prefix"`
	StartNumber int    `json:"startNumber"`
	Count       int    `json:"count"`
	model.OprVehicleOpt
}

// BatchAddHandler 新增
func (o *Vehicle) BatchAddHandler(c *gin.Context) {
	var p batchAdd
	//获取参数
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	t := middleware.GetUserToken(c)
	lzero := len(strconv.Itoa(p.StartNumber + p.Count - 1))
	var data []model.OprVehicle
	for i := 0; i < p.Count; i++ {
		v := model.OprVehicle{}
		v.OprVehicleOpt = p.OprVehicleOpt
		v.DeviceNo = fmt.Sprintf("%s%0*d", p.Prefix, lzero, p.StartNumber+i)
		v.DeviceName = v.DeviceNo
		v.Guid = util.UUID()
		v.OrganizeGuid = t.OrganizeGuid
		data = append(data, v)
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Vehicle) UpdateHandler(c *gin.Context) {
	var p model.OprVehicle
	//获取参数
	if err := c.ShouldBind(&p.OprVehicleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// resetOrganize 更新
type resetOrganize struct {
	DeviceIds  string `json:"deviceIds"`
	OrganizeId int    `json:"organizeId"`
}

func (o *Vehicle) ResetOrganizeHandler(c *gin.Context) {
	var p resetOrganize
	//获取参数
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ids := util.StringToIntSlice(p.DeviceIds, ",")
	if err := orm.DbUpdateByIds(&model.OprVehicle{}, ids, orm.H{"organize_id": p.OrganizeId}); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Vehicle) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	var devs []model.OprVehicle
	if _, err := orm.DbFindBy(&devs, "id in (?)", ids); err != nil {
		ctx.JSONError().WriteTo(c)
		return
	}
	if err := orm.DbDeletes(&devs); err != nil {
		ctx.JSONError().WriteTo(c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func VehicleRouter(r *gin.RouterGroup) {
	v := Vehicle{}
	r.GET("/vehicle/list", v.ListHandler)
	r.GET("/vehicle/:id", v.GetHandler)
	r.POST("/vehicle", v.AddHandler)
	r.POST("/vehicle/batchAdd", v.BatchAddHandler)
	r.PUT("/vehicle", v.UpdateHandler)
	r.PUT("/vehicle/resetOrganize", v.ResetOrganizeHandler)
	r.DELETE("/vehicle/:id", v.DeleteHandler)
}
