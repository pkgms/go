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
package errs

import (
	"fmt"
	"net/http"
)

var defaultHttpCode = http.StatusInternalServerError

func SetDefaultHttpCode(code int) {
	defaultHttpCode = code
}

type Error struct {
	code int
	s    string
}

func New(text string) *Error {
	return &Error{
		s:    text,
		code: defaultHttpCode,
	}
}

func (e *Error) Error() string {
	return e.s
}

func (e *Error) WithArgs(vs ...interface{}) *Error {
	str := fmt.Sprintf(e.s, vs...)
	return &Error{
		code: e.code,
		s:    str,
	}
}

func (e *Error) WithHttpCode(code int) *Error {
	return &Error{
		code: code,
		s:    e.s,
	}
}

func (e *Error) HttpCode() int {
	return e.code
}
