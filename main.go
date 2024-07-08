package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sampleGo/models"
	"sampleGo/pkg/logging"
	"sampleGo/pkg/setting"
	"sampleGo/routers"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

// @title Telepathy API
// @version 0.1.1
// @description An example of gin
// @termsOfService https://github.com/hallago/telepathy-server
// @license.name MIT
// @license.url https://github.com/hallago/telepathy-server/blob/main/LICENSE
// @schemes http https
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-api-key
func main() {
	gin.SetMode(setting.ServerSetting.RunMode)

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}
	log.Printf("[info] start http server listening %s", endPoint)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
