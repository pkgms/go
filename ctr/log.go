/**
 * Created by zc on 2020/11/20.
 */
package ctr

import "github.com/pkgms/go/logger"

// Log defines the log processing method
var Log logger.Interface = &logger.Empty{Std: &logger.StdEmpty{}}

// InitLog initialization the log processing method
func InitLog(log logger.Interface) {
	Log = log
}

// InitLog initialization the Std log processing method
func InitStdLog(log logger.StdInterface) {
	Log = &logger.Empty{Std: log}
}
