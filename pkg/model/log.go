package model

import (
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Log struct {
	gorm.Model
	SiteId     uint
	Domain     string
	IpAddress  string
	AccessTime string
	UserAgent  string
	Language   string
	Action     string // 直接记录controller就行了
	Body       string // post 的数据
}

func InsertLog(c *gin.Context, isDownload bool) {
	site := GetSite(c.Request.Host)
	domain := strings.Split(c.Request.Host, ":")[0]
	value, _ := ioutil.ReadAll(c.Request.Body)
	log.Println("[*] Body : ", string(value))
	var log = Log{
		SiteId:     site.ID,
		Domain:     domain,
		IpAddress:  GetRemoteAddr(c),
		AccessTime: time.Now().Format("2006-01-02 15-04-05"),
		UserAgent:  c.Request.UserAgent(),
		Language:   c.Request.Header.Get("Accept-Language"),
		Action:     c.Request.URL.Path,
		Body:       string(value),
	}
	Conn.Save(&log)
	if isDownload {
		AddDownCount(site.ID)
	} else {
		AddAccessCount(site.ID)
	}
}
