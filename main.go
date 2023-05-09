package main

import (
	"embed"
	"encoding/json"
	"github.com/zgwit/iot-master/v3/model"
	"github.com/zgwit/iot-master/v3/pkg/banner"
	"github.com/zgwit/iot-master/v3/pkg/build"
	"github.com/zgwit/iot-master/v3/pkg/db"
	"github.com/zgwit/iot-master/v3/pkg/log"
	"github.com/zgwit/iot-master/v3/pkg/mqtt"
	"github.com/zgwit/iot-master/v3/pkg/web"
	"net/http"
	"sms/api"
	"sms/config"
	_ "sms/docs"
	"sms/types"
)

//go:embed all:app/sms
var wwwFiles embed.FS

// @title 历史数据接口文档
// @version 1.0 版本
// @description API文档
// @BasePath /app/sms/api/
// @query.collection.format multi
func main() {
	banner.Print("iot-master-plugin:sms")
	build.Print()

	config.Load()

	err := log.Open()
	if err != nil {
		log.Fatal(err)
	}

	//加载数据库
	err = db.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//同步表结构
	err = db.Engine.Sync2(
		new(types.Sms),
	)
	if err != nil {
		log.Fatal(err)
	}

	//MQTT总线
	err = mqtt.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer mqtt.Close()

	//注册应用
	//for _, v := range config.Config.Apps {
	payload, _ := json.Marshal(model.App{
		Id:   "sms",
		Name: "短消息推送",
		Entries: []model.AppEntry{{
			Path: "app/sms/sms",
			Name: "消息历史",
		}, {
			Path: "app/sms/setting",
			Name: "配置",
		}},
		Type:    "tcp",
		Address: "http://localhost" + web.GetOptions().Addr,
	})
	_ = mqtt.Publish("master/register", payload, false, 0)
	//}

	app := web.CreateEngine()

	//注册前端接口
	api.RegisterRoutes(app.Group("/app/sms/api"))

	//注册接口文档
	web.RegisterSwaggerDocs(app.Group("/app/sms"))

	//前端静态文件
	app.RegisterFS(http.FS(wwwFiles), "", "app/sms/index.html")

	//监听HTTP
	app.Serve()
}