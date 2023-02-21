package http

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/renmcc/go-cmdb/apps/host"
)

// 用于暴露Host service接口
func (h *Handler) createHost(c *gin.Context) {
	ins := host.NewHost()
	// 将HTTP协议里面 解析出来用户的请求参数
	// read c.Request.Body
	// json unmarshal

	// 用户传递过来的参数进行解析, 实现了一个json 的unmarshal
	if err := c.ShouldBindJSON(ins); err != nil {
		h.log.Named("createHost").Error(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{"code": http.StatusBadRequest, "message": "require data error"})
		return
	}

	// 进行接口调用, 写入数据库
	ins, err := h.svc.CreateHost(c.Request.Context(), ins)
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

func (h *Handler) test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "message": "hello test"})
}
