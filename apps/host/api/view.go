package http

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/imdario/mergo"
	"github.com/renmcc/go-cmdb/apps/host"
)

// 用于暴露Host service接口
func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()

	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
	if err := c.ShouldBindJSON(ins); err != nil {
		h.log.Named("createHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 数据校验
	if err := ins.Validate(); err != nil {
		h.log.Named("createHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 进行接口调用, 写入数据库
	err := h.svc.CreateHost(c.Request.Context(), ins)
	if err != nil {
		h.log.Named("createHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}

	// 成功, 把对象实例返回给HTTP API调用方
	c.JSON(http.StatusCreated, gin.H{"code": http.StatusCreated, "message": ins})
}

func (h *Handler) listHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	qs := c.Request.URL.Query()
	// 请求参数
	req := host.NewListHostRequest()

	var err error
	if qs.Get("page_size") != "" {
		req.PageSize, err = strconv.Atoi(qs.Get("page_size"))
		if err != nil {
			h.log.Named("listHost").Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
			return
		}
	}

	if qs.Get("page_number") != "" {
		req.PageNumber, err = strconv.Atoi(qs.Get("page_number"))
		if err != nil {
			h.log.Named("listHost").Error(err.Error())
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
			return
		}
	}

	if qs.Get("serial_number") != "" {
		req.SerialNumber = qs.Get("serial_number") + "%"
	}

	if qs.Get("privateip") != "" {
		req.PrivateIp = qs.Get("privateip") + "%"
	}

	// 数据校验
	if err := req.Validate(); err != nil {
		h.log.Named("listHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
		return
	}

	// 进行数据库查询
	set, err := h.svc.ListHost(c.Request.Context(), req)
	if err != nil {
		h.log.Named("listHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": set})
}

func (h *Handler) queryHost(c *gin.Context) {

	// 从http请求的query string 中获取参数
	req := host.NewQueryHostRequest(c.Param("id"))

	// 数据校验
	if err := req.Validate(); err != nil {
		h.log.Named("queryHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
		return
	}

	// 进行数据库查询
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		h.log.Named("queryHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}

	// 空查询处理
	if set == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "no rows in result set"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": set})
}

func (h *Handler) putHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewQueryHostRequest(c.Param("id"))

	// 数据校验
	if err := req.Validate(); err != nil {
		h.log.Named("putHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
		return
	}

	// 进行数据库查询
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		h.log.Named("putHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "query Host error"})
		return
	}

	// 空查询处理
	if set == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "no rows in result set"})
		return
	}

	// 解析Body里面的数据
	ins := host.NewHost()
	if err := c.ShouldBindJSON(ins); err != nil {
		h.log.Named("putHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 确保ID一致
	ins.Xid = req.Xid

	// 数据校验
	if err := ins.Validate(); err != nil {
		h.log.Named("putHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 进行数据更新
	ret, err := h.svc.UpdateHost(c.Request.Context(), ins)
	if err != nil {
		h.log.Named("putHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": ret})
}

func (h *Handler) patchHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewQueryHostRequest(c.Param("id"))

	// 数据校验
	if err := req.Validate(); err != nil {
		h.log.Named("patchHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
		return
	}

	// 进行数据库查询
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		h.log.Named("patchHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "query Host error"})
		return
	}

	// 空查询处理
	if set == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "no rows in result set"})
		return
	}

	// 解析Body里面的数据
	ins := host.NewHost()
	if err := c.ShouldBindJSON(ins); err != nil {
		h.log.Named("patchHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 合并数据
	if err := mergo.MergeWithOverwrite(set, ins); err != nil {
		h.log.Named("patchHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 数据校验
	if err := set.Validate(); err != nil {
		h.log.Named("patchHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 确保ID一致
	set.Xid = req.Xid

	// 进行数据更新
	ret, err := h.svc.UpdateHost(c.Request.Context(), set)
	if err != nil {
		h.log.Named("patchHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": ret})
}

func (h *Handler) deleteHost(c *gin.Context) {
	// 从http请求的query string 中获取参数
	req := host.NewQueryHostRequest(c.Param("id"))
	// 数据校验
	if err := req.Validate(); err != nil {
		h.log.Named("deleteHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "args error"})
		return
	}
	// 进行数据库查询
	set, err := h.svc.QueryHost(c.Request.Context(), req)
	if err != nil {
		h.log.Named("deleteHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "query Host error"})
		return
	}

	// 空查询处理
	if set == nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "no rows in result set"})
		return
	}

	// 删除数据
	deleteReq := host.NewDeleteHostRequest(set.Xid)
	err = h.svc.DeleteHost(c.Request.Context(), deleteReq)
	if err != nil {
		h.log.Named("deleteHost").Error(err.Error())
		if strings.Contains(err.Error(), "No connection could be made because the target machine actively refused it.") {
			c.JSON(http.StatusServiceUnavailable, gin.H{"code": http.StatusServiceUnavailable, "message": "ServiceUnavailable"})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		}
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "delete success"})
}
