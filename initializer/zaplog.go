package initializer

import (
	"SH-admin/global"
	"fmt"
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
	"time"
)

func InitLogger() {
	//zapcore.Core需要3個配置——Encoder，WriteSyncer，LogLevel。
	//
	//1.Encoder:編碼器(如何寫入日誌)。
	encoder := GetEncoder()
	//2.WriterSyncer：指定日誌寫到哪裡去。
	writeSyncer := GetLogWriter()
	//3.Log Level：哪種級別的日誌會被寫入。
	core := zapcore.NewCore(encoder, writeSyncer, TransportLevel())

	logger := zap.New(core, zap.AddCaller()) //zap.AddCaller() 為一個option，添加調用的信息至日誌
	global.Log = logger.Sugar()
}

// TransportLevel 將配置文件中的level轉為 zapcore.Level
func TransportLevel() zapcore.Level {
	switch strings.ToLower(global.Config.Zap.Level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.WarnLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}

func GetEncoder() zapcore.Encoder {
	ec := zap.NewProductionEncoderConfig()
	ec.EncodeTime = FormatTimeEncoder            //自定義日期格式，原本是zapcore.ISO8601TimeEncoder
	ec.EncodeLevel = zapcore.CapitalLevelEncoder //level變成大寫
	ec.MessageKey = "message"
	ec.LevelKey = "level"
	ec.TimeKey = "time"
	ec.NameKey = "name"
	ec.CallerKey = "caller"
	ec.EncodeCaller = zapcore.FullCallerEncoder //caller的長短

	//根據配置文件決定輸出格式
	switch strings.ToLower(global.Config.Zap.Format) {
	case "json":
		return zapcore.NewJSONEncoder(ec) //輸出json格式
	case "console":
		return zapcore.NewConsoleEncoder(ec) //輸出控制台格式
	default:
		return zapcore.NewConsoleEncoder(ec)
	}
}

// FormatTimeEncoder 指定日誌輸出的日期格式
func FormatTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006/01/02 - 15:04:05"))
}

func GetLogWriter() zapcore.WriteSyncer {
	//Lumberjack Logger采用以下属性作为输入:
	//
	//Filename: 檔案位置
	//MaxSize：在進行切割之前，檔案的最大大小（以MB為單位）
	//MaxBackups：保留舊檔案的最大個數
	//MaxAges：保留舊文件的最大天数
	//Compress：是否壓縮/歸檔舊文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.log", global.Config.Zap.Director, "log"),
		MaxSize:    100,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	if global.Config.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}
