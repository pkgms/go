/*
Copyright © 2021 zc2638 <zc2638@qq.com>.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package logger

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
	"time"
)

const (
	// DefaultPath 默认地址
	DefaultPath = "_log"
)

type FileHook struct {
	ws       map[logrus.Level]io.ReadWriteCloser
	dir      string
	datetime string
}

func NewFileHook(dir string) *FileHook {
	if dir == "" {
		dir = DefaultPath
	}
	return &FileHook{
		dir: dir,
		ws:  make(map[logrus.Level]io.ReadWriteCloser),
	}
}

func (h *FileHook) SetDir(dir string) *FileHook {
	filterDir := strings.TrimRight(dir, "/")
	if filterDir != "" {
		h.dir = filterDir
	}
	return h
}

func (h *FileHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

// 自定义钩子执行（默认协程安全）
func (h *FileHook) Fire(e *logrus.Entry) error {
	// 判断log文件是否变更
	now := time.Now().Format("20060102")
	path := h.dir + "/" + e.Level.String() + "_" + now + ".log"
	if h.datetime != now {
		h.datetime = now
		if h.ws[e.Level] != nil {
			if err := h.ws[e.Level].Close(); err != nil {
				return err
			}
			h.ws[e.Level] = nil
		}
	}
	if _, err := os.Stat(path); err == nil && h.ws[e.Level] != nil {
		return nil
	}

	// 自动创建文件
	var pathArr = strings.Split(path, "/")
	var pathLen = len(pathArr)
	if pathLen > 1 {
		dir := strings.Join(pathArr[:pathLen-1], "/")
		// 自动创建日志文件夹
		_, err := os.Stat(dir)
		if err != nil {
			mkErr := os.MkdirAll(dir, os.ModePerm)
			if mkErr != nil {
				return mkErr
			}
		}
	}

	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		return err
	}

	h.ws[e.Level] = f
	e.Logger.Out = h.ws[e.Level]
	return nil
}
