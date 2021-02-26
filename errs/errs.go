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
	"net/http"

	"github.com/pkgms/go/i18n"
)

var defaultHttpCode = http.StatusInternalServerError

func SetDefaultHttpCode(code int) {
	defaultHttpCode = code
}

type Error = i18n.Error

func New(text string) *Error {
	return &Error{
		HttpCode: defaultHttpCode,
		Option: i18n.Option{
			Value: text,
		},
	}
}
