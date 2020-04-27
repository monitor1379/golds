package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:24:48
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-27 23:25:36
 */

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	var err error
	config := zap.NewProductionConfig()

	logLevel := zapcore.InfoLevel
	if isDebug, ok := os.LookupEnv("GOLDS_LOG_DEBUG"); ok && isDebug == "on" {
		logLevel = zapcore.DebugLevel
	}
	config.Level.SetLevel(logLevel)

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	logger, err = config.Build()
	if err != nil {
		panic(err)
	}

	logger.Info("logger created", zap.String("level", logLevel.String()))
}
