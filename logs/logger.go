package logs

import (
	"encoding/json"
	"fmt"

	"myproject/conf"

	"github.com/astaxie/beego/logs"
)

var Logger *LoggerProvider

type LoggerProvider struct {
	*logs.BeeLogger
}

var LogLevelMap = map[string]int{
	"debug":   logs.LevelDebug,
	"info":    logs.LevelInfo,
	"warning": logs.LevelWarning,
	"error":   logs.LevelError,
}

func InitLogger() {
	Logger = new(LoggerProvider)
	Logger.BeeLogger = logs.GetBeeLogger()

	logConf := conf.Config.Logger
	_ = Logger.SetLogger(logs.AdapterConsole)

	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 直接调用的层级, 默认是 2
	logs.SetLogFuncCallDepth(3)
	// 异步输出日志 允许设置缓冲 chan 的大小
	logs.Async(1e3)
	// 设置日志等级
	logLevel := LogLevelMap[logConf.LogLevel]
	if logLevel == 0 {
		logLevel = logs.LevelInfo
	}
	logs.SetLevel(logLevel)

	var configMap = make(map[string]interface{})
	configMap["filename"] = fmt.Sprintf("%v/%v.%v", logConf.LogDirPath, logConf.LogFileName, "log")
	configMap["color"] = true
	configMap["daily"] = true
	if logConf.LogStorageDay > 0 {
		configMap["maxdays"] = logConf.LogStorageDay
	}
	configMap["maxsize"] = logConf.LogMaxSize

	confByte, err := json.Marshal(configMap)
	if err != nil {
		panic(err)
	}

	err = logs.SetLogger(logs.AdapterFile, string(confByte))
	if err != nil {
		panic(err)
	}
}