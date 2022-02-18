package system

import (
	"errors"
	"xserver/model"
	"xserver/service"
	"xserver/util"

	"github.com/wlgd/xutils/ctx"
	"github.com/wlgd/xutils/orm"

	"github.com/gin-gonic/gin"
)

// Dept 部门
type Dept struct {
}

// PageHandler 列表
func (o *Dept) PageHandler(c *gin.Context) {
	var param service.DeptPage
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var depts []model.SysDept
	orm.DbPage(&model.SysDept{}, param.Where()).Find(param.PageNum, param.PageSize, &depts)
	ctx.JSONOk().WriteData(depts, c)
}

// ListExcludeHandler 部门列表（排除节点）
func (o *Dept) ListExcludeHandler(c *gin.Context) {
	// id, err := ctxQueryInt(c, "id")
	// if err != nil {
	// 	JSONP(StatusError).WriteTo(c)
	// }
	// where := fmt.Sprintf("id != %d", id)
	// var depts []model.Dept
	// orm.DbFindAll(where, depts, "order_num asc")
	ctx.JSONOk().WriteTo(c)
}

// GetHandler 查询部门详细
func (o *Dept) GetHandler(c *gin.Context) {
	deptId, err := ctx.ParamInt(c, "id")
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	var dept model.SysDept
	err = orm.DbFirstById(&dept, deptId)
	if err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteData(dept, c)
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
	ctx.JSONOk().WriteTo(c)
}

func checkAddDept(req *model.SysDept) error {
	var dept model.SysDept
	if err := orm.DbFirstBy(&dept, "dept_name like ?", req.DeptName); err != nil {
		return err
	}
	return nil
}

// AddHandler 新增
func (o *Dept) AddHandler(c *gin.Context) {
	var param model.SysDept
	//获取参数
	if err := c.ShouldBind(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := checkAddDept(&param); err == nil {
		ctx.JSONWriteError(errors.New("dept already exists"), c)
		return
	}
	if err := orm.DbCreate(&param); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// UpdateHandler 修改
func (o *Dept) UpdateHandler(c *gin.Context) {
	var dept model.SysDept
	if err := c.ShouldBind(&dept); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	if err := orm.DbUpdateModel(dept); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

// DeleteHandler 删除
func (o *Dept) DeleteHandler(c *gin.Context) {
	idstr := ctx.ParamString(c, "id")
	if idstr == "" {
		ctx.JSONError().WriteTo(c)
		return
	}
	ids := util.StringToIntSlice(idstr, ",")
	if err := orm.DbDeleteByIds(model.SysDept{}, ids); err != nil {
		ctx.JSONWriteError(err, c)
		return
	}
	ctx.JSONOk().WriteTo(c)
}

func DeptRouters(r *gin.RouterGroup) {
	sysDept := Dept{}
	r.GET("/dept/list", sysDept.PageHandler)
	// r.GET("/dept/list/exclude/:id", sysDept.ListExcludeHandler)
	r.GET("/dept/:id", sysDept.GetHandler)
	r.GET("/dept/treeselect", sysDept.TreeselectHandler)
	r.GET("/dept/roleDeptTreeselect/:id", sysDept.RoleDeptTreeselectHandler)
	r.POST("/dept", sysDept.AddHandler)
	r.PUT("/dept", sysDept.UpdateHandler)
	r.DELETE("/dept/:id", sysDept.DeleteHandler)
}
