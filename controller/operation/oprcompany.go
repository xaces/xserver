package operation

import (
	"xserver/middleware"
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Company
type Company struct {
}

// ListHandler 列表
func (o *Company) ListHandler(c *gin.Context) {
	var p orm.DbPage
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	where := p.DbWhere()
	where.Append("parent_id = ?", 0) // 上级节点为0，表示公司
	var data []model.OprOrganization
	toatl, _ := orm.DbByWhere(&model.OprOrganization{}, where).Preload("SysStation").Find(&data)
	ctx.JSONOk().Write(gin.H{"total": toatl, "data": data}, c)
}

// GetHandler 详细
func (o *Company) GetHandler(c *gin.Context) {
	var data model.OprOrganization
	service.QueryById(&data, c)
}

// AddHandler 新增
func (o *Company) AddHandler(c *gin.Context) {
	var data model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&data.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	data.OrganizeGuid = util.NUID()
	tok := middleware.GetUserToken(c)
	u := &model.SysUser{}
	u.UserName = data.UserName
	u.CreatedBy = tok.UserName
	u.OrganizeName = data.Name
	u.UserType = model.SysUserTypeAdmin
	u.OrganizeGuid = data.OrganizeGuid
	if err := service.SysUserCreate(u); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbCreate(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

type companyConfig struct {
	CompanyId uint64 `json:"companyId" binding:"required"`
	StationId uint64 `json:"stationId"`
}

// AddHandler 新增
func (o *Company) ConfigHandler(c *gin.Context) {
	var param companyConfig
	//获取参数
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	orm.DbUpdateColById(&model.OprOrganization{}, param.CompanyId, "station_id", param.StationId)
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Company) UpdateHandler(c *gin.Context) {
	var data model.OprOrganization
	//获取参数
	if err := c.ShouldBind(&data.OprOrganizationOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&data); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Company) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.OprOrganization{}, c)
}

func CompanyRouters(r *gin.RouterGroup) {
	o := Company{}
	r.GET("/company/list", o.ListHandler)
	r.GET("/company/:id", o.GetHandler)
	r.POST("/company", o.AddHandler)
	r.PUT("/company/config", o.ConfigHandler)
	r.PUT("/company", o.UpdateHandler)
	r.DELETE("/company/:id", o.DeleteHandler)
}
