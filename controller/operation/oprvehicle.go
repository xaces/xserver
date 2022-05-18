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
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	p.OrganizeGUID = middleware.GetUserToken(c).OrganizeGUID
	var data []model.OprVehicle
	total, _ := orm.DbByWhere(&data, p.Vehicle()).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// GetHandler 获取指定id
func (o *Vehicle) GetHandler(c *gin.Context) {
	service.QueryById(&model.OprVehicle{}, c)
}

// AddHandler 新增
func (o *Vehicle) AddHandler(c *gin.Context) {
	var p model.OprVehicle
	//获取参数
	if err := c.ShouldBind(&p.OprVehicleOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	p.GUID = util.UUID()
	p.OrganizeGUID = middleware.GetUserToken(c).OrganizeGUID
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
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
		v.GUID = util.UUID()
		v.OrganizeGUID = t.OrganizeGUID
		data = append(data, v)
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
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
	ctx.JSONOk(c)
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
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *Vehicle) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.OprVehicle{}, c)
}

func VehicleRouter(r *gin.RouterGroup) {
	v := Vehicle{}
	r.GET("/list", v.ListHandler)
	r.GET("/:id", v.GetHandler)
	r.POST("", v.AddHandler)
	r.POST("/batchAdd", v.BatchAddHandler)
	r.PUT("", v.UpdateHandler)
	r.PUT("/resetOrganize", v.ResetOrganizeHandler)
	r.DELETE("/:id", v.DeleteHandler)
}
