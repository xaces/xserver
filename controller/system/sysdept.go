package system

import (
	"errors"
	"xserver/model"
	"xserver/service"

	"github.com/xaces/xutils/ctx"
	"github.com/xaces/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Dept 部门
type Dept struct {
}

// PageHandler 列表
func (o *Dept) PageHandler(c *gin.Context) {
	var p Where
	if err := c.ShouldBind(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var data []model.SysDept
	total, _ := p.Where().Model(&model.SysDept{}).Find(&data)
	ctx.JSONWrite(gin.H{"total": total, "data": data}, c)
}

// ListExcludeHandler 列表（排除节点）
func (o *Dept) ListExcludeHandler(c *gin.Context) {
	// id, err := ctxQueryInt(c, "id")
	// if err != nil {
	// 	JSONP(StatusError).WriteTo(c)
	// }
	// where := fmt.Sprintf("id != %d", id)
	// var depts []model.Dept
	// orm.DbFindAll(where, depts, "order_num asc")
	ctx.JSONOk(c)
}

// GetHandler 查询详细
func (o *Dept) GetHandler(c *gin.Context) {
	service.QueryByID(&model.SysDept{}, c)
}

// TreeselectHandler 查询下拉树结构
func (o *Dept) TreeselectHandler(c *gin.Context) {
	// trees, err := buildDeptTree(c)
	// if err != nil {
	// 	ctx.JSONWriteError(err, c)
	// }
	// ctx.JSONOk().WriteData(trees, c)
}

// RoleDeptTreeselectHandler 根据角色ID查询树结构
func (o *Dept) RoleDeptTreeselectHandler(c *gin.Context) {
	ctx.JSONOk(c)
}

func checkAddDept(req *model.SysDept) error {
	var data model.SysDept
	if err := orm.DbFirstBy(&data, "dept_name like ?", req.DeptName); err != nil {
		return err
	}
	return nil
}

// AddHandler 新增
func (o *Dept) AddHandler(c *gin.Context) {
	var p model.SysDept
	//获取参数
	if err := c.ShouldBind(&p.SysDeptOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := checkAddDept(&p); err == nil {
		ctx.JSONWriteError(errors.New("dept already exists"), c)
		return
	}
	if err := orm.DbCreate(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// UpdateHandler 修改
func (o *Dept) UpdateHandler(c *gin.Context) {
	var p model.SysDept
	if err := c.ShouldBind(&p.SysDeptOpt); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(&p); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk(c)
}

// DeleteHandler 删除
func (o *Dept) DeleteHandler(c *gin.Context) {
	service.Deletes(&model.SysDept{}, c)
}

func (o Dept) Routers(r *gin.RouterGroup) {
	r.GET("/list", o.PageHandler)
	// r.GET("/list/exclude/:id", o.ListExcludeHandler)
	r.GET("/:id", o.GetHandler)
	r.GET("/treeselect", o.TreeselectHandler)
	r.GET("/roleDeptTreeselect/:id", o.RoleDeptTreeselectHandler)
	r.POST("", o.AddHandler)
	r.PUT("", o.UpdateHandler)
	r.DELETE("/:id", o.DeleteHandler)
}
