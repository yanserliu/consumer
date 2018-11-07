package log

import (
	"os"

	"fmt"
	//"uauth/infrastructure/helper"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

//日志变量
var Logger = logs.NewLogger()

//日志初始化
func InitLogs() {
	//创建日志目录
	if _, err := os.Stat("logs"); err != nil {
		os.Mkdir("logs", os.ModePerm)
	}
	level, _ := beego.AppConfig.Int("logs::level")

	maxLines, _ := beego.AppConfig.Int64("logs::max_lines")

	maxDays, _ := beego.AppConfig.Int("logs::max_days")

	//初始化日志各种配置
	LogsConf := fmt.Sprintf(`{"filename":"logs/uauth.log","level":%v,"maxlines":%v,"maxsize":0,"daily":true,"maxdays":%v}`, level, maxLines, maxDays)
	Logger.SetLogger(logs.AdapterFile, LogsConf)
	if level == 7 {
		Logger.SetLogger("console")
		beego.Info("日志配置信息：" + LogsConf)
	} else {
		//是否异步输出日志
		Logger.Async(1e3)
	}
	Logger.EnableFuncCallDepth(true) //是否显示文件和行号
}
