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
package i18n

import "fmt"

type Error struct {
	Option
	HttpCode int
}

func NewError(opt Option) *Error {
	return &Error{
		Option: opt,
	}
}

func NewErrorWithData(data *Data) *Error {
	return NewError(data.Option()).WithHttpCode(data.HttpCode())
}

func (e *Error) WithArgs(args ...interface{}) *Error {
	opt := e.Option
	opt.Args = args
	return &Error{Option: opt}
}

func (e *Error) WithHttpCode(code int) *Error {
	e.HttpCode = code
	return e
}

func (e *Error) Error() string {
	return fmt.Sprintf(e.Option.Value, e.Option.Args...)
}

func VerifyError(locales *Locales, err *Error) *Error {
	opt := err.Option
	format := locales.Get(err.Option)
	if format != "" {
		opt.Value = format
	}
	return &Error{Option: opt}
}
