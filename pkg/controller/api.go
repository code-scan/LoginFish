package controller

import (
	"LoginFish/pkg/config"
	"LoginFish/pkg/model"
	"github.com/gin-gonic/gin"
	"log"
	"path"
	"path/filepath"
	"strings"
)

type Api struct {
}

func (a *Api) check(c *gin.Context) bool {
	if pwd, err := c.Cookie("pwd"); err != nil {
		log.Println(err)
		return false
	} else if strings.TrimSpace(pwd) != config.GetPassword() {
		log.Println(pwd, config.GetPassword())
		return false
	}
	return true
}
func (a *Api) AddSite(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	var site model.Site
	if err := c.BindJSON(&site); err == nil {
		model.Conn.Save(&site)
		c.String(200, "success")
		return
	}
	c.String(200, "fail")

}
func (a *Api) GetSite(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	var site []model.Site
	model.Conn.Model(&model.Site{}).Find(&site)
	c.JSON(200, &site)
}
func (a *Api) EditSite(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	id := c.GetInt("id")
	var site model.Site
	if err := c.BindJSON(&site); err == nil {
		model.Conn.Model(&model.Site{}).Where("id = ?", id).Updates(&site)
		c.String(200, "success")
		return
	}
}
func (a *Api) DelSite(c *gin.Context) {
	if a.check(c) == false {
		c.Status(404)
		return
	}
	id := c.Query("id")
	model.Conn.Where("id = ?", id).Delete(&model.Site{})
	c.String(200, "success")
}

func (a *Api) GetLog(c *gin.Context) {
	id := c.Query("id")
	log.Println(id)
	var logs []model.Log
	if id == "-1" {
		model.Conn.Model(&model.Log{}).Order("id desc").Limit(200).Find(&logs)
	} else {
		model.Conn.Model(&model.Log{}).Where("site_id = ? ", id).Order("id desc").Limit(200).Find(&logs)
	}
	c.JSON(200, &logs)
}

func (a *Api) GetTemplate(c *gin.Context) {
	templates, _ := filepath.Glob("static/template/front_login*")
	for i := range templates {
		templates[i] = path.Base(templates[i])
	}
	c.JSON(200, templates)
}
