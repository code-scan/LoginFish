package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strings"
	"time"
)

type Log struct {
	gorm.Model
	SiteId     uint
	Domain     string
	IpAddress  string
	AccessTime time.Time
	UserAgent  string
	Language   string
	Action     string // 直接记录controller就行了
}

func InsertLog(c *gin.Context, isDownload bool) {
	site := GetSite(c.Request.Host)
	domain := strings.Split(c.Request.Host, ":")[0]

	var log = Log{
		SiteId:     site.ID,
		Domain:     domain,
		IpAddress:  GetRemoteAddr(c),
		AccessTime: time.Now(),
		UserAgent:  c.Request.UserAgent(),
		Language:   c.Request.Header.Get("Accept-Language"),
		Action:     c.Request.URL.Path,
	}
	Conn.Save(&log)
	if isDownload {
		AddDownCount(site.ID)
	} else {
		AddAccessCount(site.ID)
	}
}
