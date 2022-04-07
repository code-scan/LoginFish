package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"math/rand"
	"time"
)

var Conn *gorm.DB

func init() {
	var err error
	Conn, err = gorm.Open(sqlite.Open("data/fish.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	Conn.AutoMigrate(&Log{})
	Conn.AutoMigrate(&Site{})
	Conn.Logger = logger.Default.LogMode(logger.Info)
}
func GetRemoteAddr(c *gin.Context) string {
	var ip string
	if ip = c.Request.Header.Get("CF-Connecting-IP"); ip != "" {
		return ip
	}
	if ip = c.Request.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	if ip = c.Request.Header.Get("Client-IP"); ip != "" {
		return ip
	}
	return c.Request.RemoteAddr
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RandStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
