package controller

import (
	"LoginFish/pkg/config"
	"LoginFish/pkg/model"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"path"
	"strings"
)

type Admin struct {
}

func (a *Admin) check(c *gin.Context) bool {
	if pwd, err := c.Cookie("pwd"); err != nil {
		log.Println(err)
		return false
	} else if strings.TrimSpace(pwd) != config.GetPassword() {
		log.Println(pwd, config.GetPassword())
		return false
	}
	return true
}
func (a *Admin) Login(c *gin.Context) {

	c.HTML(200, "admin_login.html", nil)
}

// Index 导航页面
func (a *Admin) Index(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	c.HTML(200, "admin_index.html", nil)
}

// Site  站点列表 展示站点列表和访问次数
func (a *Admin) Site(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	c.HTML(200, "admin_site.html", nil)
}

// EditSite TODO 编辑站点
func (a *Admin) EditSite(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
}

// SiteAdd  增加站点
func (a *Admin) SiteAdd(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	err := c.Request.ParseMultipartForm(200000)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%#v", c.Request.MultipartForm)

	var site = model.Site{
		Domain:   c.Request.MultipartForm.Value["Domain"][0],
		Title:    c.Request.MultipartForm.Value["Title"][0],
		Mark:     c.Request.MultipartForm.Value["Mark"][0],
		Template: c.Request.MultipartForm.Value["Template"][0],
	}

	if files, ok := c.Request.MultipartForm.File["File"]; ok {
		files := files[0]
		download := "static/download/" + path.Base(site.Domain) + "_" + model.RandStringRunes(20) + "__" + files.Filename
		site.Download = download
		f, _ := files.Open()
		out, _ := os.Create(download)
		io.Copy(out, f)
	}

	if id, ok := c.Request.MultipartForm.Value["ID"]; ok && len(id) > 0 && id[0] != "" {
		model.Conn.Model(&model.Site{}).Where("id = ?", id[0]).Updates(&site)
	} else {
		model.Conn.Model(&model.Site{}).Save(&site)

	}
	c.Redirect(302, "/"+config.AdminPath+"/site")
}

// Log 访问日志
func (a *Admin) Log(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	id := c.Query("id")
	if id == "" {
		id = "-1"
	}
	c.HTML(200, "admin_log.html", gin.H{
		"ID": id,
	})

}
