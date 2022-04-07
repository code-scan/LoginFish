package model

import (
	"gorm.io/gorm"
	"strings"
)

type Site struct {
	gorm.Model
	UUID          string // siteid - uuid 字符串，需要用来替换马里面的资源
	Domain        string // 域名
	Title         string // 网页标题
	UserName      string
	PassWord      string
	Download      string // 马下载地址
	AccessCount   uint   // 访问次数
	DownloadCount uint   // 马下载次数
	Template      string // 模板
	Mark          string // 备注
}

func GetSite(domain string) Site {
	domain = strings.Split(domain, ":")[0]
	var site Site
	Conn.Model(&site).Where("Domain = ? ", strings.ToLower(domain)).First(&site)
	return site
}

func AddAccessCount(id uint) {
	Conn.Model(&Site{}).Where("id = ?", id).UpdateColumn("access_count", gorm.Expr("access_count + 1"))
}
func AddDownCount(id uint) {

	Conn.Model(&Site{}).Where("id = ?", id).UpdateColumn("download_count", gorm.Expr("download_count + 1"))
}
