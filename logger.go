package golds

/*
 * @Author: ZhenpengDeng(monitor1379)
 * @Date: 2020-04-26 16:24:48
 * @Last Modified by: ZhenpengDeng(monitor1379)
 * @Last Modified time: 2020-04-26 16:29:59
 */

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	var err error
	config := zap.NewProductionConfig()
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}
