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

import (
	"fmt"
)

type MultiData struct {
	lang   string
	prefix string
}

func NewMultiData(lang string) *MultiData {
	return &MultiData{lang: lang}
}

func (d *MultiData) Prefix(prefix string) *MultiData {
	return &MultiData{
		lang:   d.lang,
		prefix: prefix,
	}
}

func (d *MultiData) Key(key string) *Data {
	return NewData().Lang(d.lang).Prefix(d.prefix).Key(key)
}

func (d *MultiData) Value(value string) *Data {
	return NewData().Lang(d.lang).Prefix(d.prefix).Text(value)
}

type Data struct {
	httpCode int
	lang     string
	prefix   string
	key      string
	text     string
	args     []interface{}
}

func NewData() *Data {
	return &Data{}
}

func (d *Data) clone() *Data {
	return &Data{
		httpCode: d.httpCode,
		lang:     d.lang,
		prefix:   d.prefix,
		key:      d.key,
		text:     d.text,
	}
}

func (d *Data) Lang(lang string) *Data {
	d.lang = lang
	return d
}

func (d *Data) Prefix(prefix string) *Data {
	d.prefix = prefix
	return d
}

func (d *Data) Key(key string) *Data {
	d.key = key
	return d
}

func (d *Data) Text(value string) *Data {
	d.text = value
	return d
}

func (d *Data) WithArgs(args ...interface{}) *Data {
	data := d.clone()
	data.args = args
	return data
}

func (d *Data) WithHttpCode(code int) *Data {
	d.httpCode = code
	return d
}

func (d *Data) HttpCode() int {
	return d.httpCode
}

func (d *Data) String() string {
	return fmt.Sprintf(d.text, d.args...)
}

func (d *Data) Option() Option {
	return Option{
		Language: d.lang,
		Prefix:   d.prefix,
		Key:      d.key,
		Value:    d.text,
		Args:     d.args,
	}
}

func (d *Data) Error() *Error {
	return &Error{
		Option:   d.Option(),
		HttpCode: d.httpCode,
	}
}

func Verify(locales *Locales, data *Data) string {
	format := locales.Get(data.Option())
	if format == "" {
		format = data.text
	}
	return fmt.Sprintf(format, data.args...)
}
