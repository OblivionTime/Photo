/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2021-12-20 09:58:16
 * @LastEditors: solid
 * @LastEditTime: 2021-12-29 09:30:52
 */
/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2021-12-09 09:45:17
 * @LastEditors: solid
 * @LastEditTime: 2021-12-09 13:55:11
 */
package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"
	"time"
	_ "unsafe"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type myLog struct {
	ZapLog *zap.Logger
	debug  bool
}

var Log *myLog

func (log *myLog) Info(v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintln(v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.InfoLevel, message); ce != nil {
		ce.Write(fields...)
	}
}
func (log *myLog) Infof(format string, v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintf(format, v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.InfoLevel, message); ce != nil {
		ce.Write(fields...)
	}
}
func (log *myLog) Error(v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintln(v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.WarnLevel, message); ce != nil {
		ce.Write(fields...)
	}
}
func (log *myLog) Errorf(format string, v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintf(format, v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.WarnLevel, message); ce != nil {
		ce.Write(fields...)
	}
}

func (log *myLog) Debug(v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintln(v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.DebugLevel, message); ce != nil {
		ce.Write(fields...)
	}
}
func (log *myLog) Debugf(format string, v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintf(format, v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.DebugLevel, message); ce != nil {
		ce.Write(fields...)
	}
}
func (log *myLog) Fatal(v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintln(v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.FatalLevel, message); ce != nil {
		ce.Write(fields...)
	}
	os.Exit(1)
}
func (log *myLog) Fatalf(format string, v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintf(format, v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.FatalLevel, message); ce != nil {
		ce.Write(fields...)
	}
	os.Exit(1)
}
func (log *myLog) Panic(v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintln(v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.PanicLevel, message); ce != nil {
		ce.Write(fields...)
	}
	panic(message)
}
func (log *myLog) Patalf(format string, v ...interface{}) {
	if !log.debug {
		return
	}
	message := fmt.Sprintf(format, v...)
	fields := []zap.Field{}
	if ce := log.ZapLog.CheckCe(zap.PanicLevel, message); ce != nil {
		ce.Write(fields...)
	}
	panic(message)
}

var levelMap = map[string]zapcore.Level{

	"debug": zapcore.DebugLevel,

	"info": zapcore.InfoLevel,

	"warn": zapcore.WarnLevel,

	"error": zapcore.ErrorLevel,

	"dpanic": zapcore.DPanicLevel,

	"panic": zapcore.PanicLevel,

	"fatal": zapcore.FatalLevel,
}

//生成日志文件
func setLoggerFile(lev_name string, encoder zapcore.Encoder) zapcore.Core {
	//日志级别
	priority := zap.LevelEnablerFunc(func(lev2 zapcore.Level) bool { //info和debug级别,debug级别是最低的
		return levelMap[lev_name] == lev2
	})
	var abPath string
	_, filename, _, ok := runtime.Caller(0)
	if ok {
		abPath = path.Dir(filename)
	}

	filename = fmt.Sprintf(`%s\log\%s-%s-%s\%s.log`, abPath, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"), lev_name)
	//info文件writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename, //日志文件存放目录，如果文件夹不存在会自动创建
		MaxSize:    10,       //文件大小限制,单位MB
		MaxBackups: 100,      //最大保留日志文件数量
		MaxAge:     30,       //日志文件保留天数
		Compress:   false,    //是否压缩处理
	})
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), priority)

}

func InitLog(debug bool) {
	var coreArr []zapcore.Core
	Log = &myLog{
		ZapLog: nil,
		debug:  debug,
	}
	//获取编码器
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()输出json格式，NewConsoleEncoder()输出普通文本格式
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //指定时间格式
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //按级别显示不同颜色，不需要的话取值zapcore.CapitalLevelEncoder就可以了
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //显示完整文件路径
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	for key := range levelMap {
		fileCore := setLoggerFile(key, encoder)
		coreArr = append(coreArr, fileCore)
	}
	Log.ZapLog = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
}
