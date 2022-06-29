package router

import (
	"LoginFish/pkg/config"
	"LoginFish/pkg/controller"
	"LoginFish/static"
	"net/http"

	"github.com/gin-gonic/gin"
)

var Router = gin.Default()

func init() {
	Router.StaticFS("/static", http.FS(static.StatFiles))
	index := new(controller.Index)
	Router.Any("/", index.Index)
	Router.Any("/down", index.Download)
	Router.POST("/login", index.Submit)

	admin := new(controller.Admin)
	adminRouter := Router.Group("/" + config.AdminPath)
	adminRouter.Any("/", admin.Index)
	adminRouter.Any("/login", admin.Login)
	adminRouter.GET("/site", admin.Site)
	adminRouter.POST("/site", admin.SiteAdd)
	adminRouter.Any("/log", admin.Log)

	api := new(controller.Api)
	apiRouter := Router.Group("/v1")
	apiRouter.Any("/add_site", api.AddSite)
	apiRouter.Any("/edit_site", api.EditSite)
	apiRouter.Any("/get_site", api.GetSite)
	apiRouter.Any("/del_site", api.DelSite)
	apiRouter.Any("/get_log", api.GetLog)
	apiRouter.Any("/get_template", api.GetTemplate)

}
