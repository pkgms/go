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

type Interface interface {
	InfoInterface
	WarnInterface
	ErrorInterface
	FatalInterface
	PanicInterface
}
type StdInterface interface {
	InfoInterface
	FatalInterface
	PanicInterface
}

type InfoInterface interface {
	Print(v ...interface{})
	Printf(format string, v ...interface{})
	Println(v ...interface{})
}

type WarnInterface interface {
	Warning(v ...interface{})
	Warningf(format string, v ...interface{})
	Warningln(v ...interface{})
}

type ErrorInterface interface {
	Error(v ...interface{})
	Errorf(format string, v ...interface{})
	Errorln(v ...interface{})
}

type FatalInterface interface {
	Fatal(v ...interface{})
	Fatalf(format string, v ...interface{})
	Fatalln(v ...interface{})
}

type PanicInterface interface {
	Panic(v ...interface{})
	Panicf(format string, v ...interface{})
	Panicln(v ...interface{})
}

type Empty struct {
	Std StdInterface
}

func (l *Empty) Print(v ...interface{}) {
	l.Std.Print(v...)
}

func (l *Empty) Printf(format string, v ...interface{}) {
	l.Std.Printf(format, v...)
}

func (l *Empty) Println(v ...interface{}) {
	l.Std.Println(v...)
}

func (l *Empty) Warning(v ...interface{})                 {}
func (l *Empty) Warningf(format string, v ...interface{}) {}
func (l *Empty) Warningln(v ...interface{})               {}

func (l *Empty) Error(v ...interface{})                 {}
func (l *Empty) Errorf(format string, v ...interface{}) {}
func (l *Empty) Errorln(v ...interface{})               {}

func (l *Empty) Fatal(v ...interface{}) {
	l.Std.Fatal(v...)
}

func (l *Empty) Fatalf(format string, v ...interface{}) {
	l.Std.Fatalf(format, v...)
}

func (l *Empty) Fatalln(v ...interface{}) {
	l.Std.Fatalln(v...)
}

func (l *Empty) Panic(v ...interface{}) {
	l.Std.Panic(v...)
}

func (l *Empty) Panicf(format string, v ...interface{}) {
	l.Std.Panicf(format, v...)
}

func (l *Empty) Panicln(v ...interface{}) {
	l.Std.Panicln(v...)
}

type StdEmpty struct{}

func (l *StdEmpty) Print(v ...interface{})                 {}
func (l *StdEmpty) Printf(format string, v ...interface{}) {}
func (l *StdEmpty) Println(v ...interface{})               {}

func (l *StdEmpty) Fatal(v ...interface{})                 {}
func (l *StdEmpty) Fatalf(format string, v ...interface{}) {}
func (l *StdEmpty) Fatalln(v ...interface{})               {}

func (l *StdEmpty) Panic(v ...interface{})                 {}
func (l *StdEmpty) Panicf(format string, v ...interface{}) {}
func (l *StdEmpty) Panicln(v ...interface{})               {}
