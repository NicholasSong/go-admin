package router

import (
	"github.com/gin-gonic/gin"
	"go-admin/app/admin/apis"
)

func init() {
	routerNoCheckRole = append(routerNoCheckRole, registerTimeDemoRouter)
}

// registerSysApiRouter
func registerTimeDemoRouter(v1 *gin.RouterGroup) {
	v1.GET("/demo/time", apis.GetCurrentTime)
}
