/*
 * @Description:
 * @Version: 1.0
 * @Autor: solid
 * @Date: 2021-12-20 09:58:16
 * @LastEditors: solid
 * @LastEditTime: 2022-07-11 16:37:44
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

//??????????????????
func setLoggerFile(lev_name string, encoder zapcore.Encoder) zapcore.Core {
	//????????????
	priority := zap.LevelEnablerFunc(func(lev2 zapcore.Level) bool { //info???debug??????,debug??????????????????
		return levelMap[lev_name] == lev2
	})

	filename := fmt.Sprintf(`.\log\%s-%s-%s\%s.log`, time.Now().Format("2006"), time.Now().Format("01"), time.Now().Format("02"), lev_name)
	//info??????writeSyncer
	infoFileWriteSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   filename, //??????????????????????????????????????????????????????????????????
		MaxSize:    10,       //??????????????????,??????MB
		MaxBackups: 100,      //??????????????????????????????
		MaxAge:     30,       //????????????????????????
		Compress:   false,    //??????????????????
	})
	return zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(infoFileWriteSyncer, zapcore.AddSync(os.Stdout)), priority)

}

func InitLog(debug bool) {
	var coreArr []zapcore.Core
	Log = &myLog{
		ZapLog: nil,
		debug:  debug,
	}
	//???????????????
	encoderConfig := zap.NewProductionEncoderConfig()            //NewJSONEncoder()??????json?????????NewConsoleEncoder()????????????????????????
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder        //??????????????????
	encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder //???????????????????????????????????????????????????zapcore.CapitalLevelEncoder????????????
	//encoderConfig.EncodeCaller = zapcore.FullCallerEncoder        //????????????????????????
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	for key := range levelMap {
		fileCore := setLoggerFile(key, encoder)
		coreArr = append(coreArr, fileCore)
	}
	Log.ZapLog = zap.New(zapcore.NewTee(coreArr...), zap.AddCaller())
}
