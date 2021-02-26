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
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Option struct {
	Language string
	Prefix   string
	Key      string
	Value    string
	Args     []interface{}
}

type Locales struct {
	path string
	data map[string]Resources
}

func NewLocale(path string, langs ...string) (*Locales, error) {
	locale := &Locales{
		path: path,
		data: make(map[string]Resources),
	}
	err := locale.Add(langs...)
	return locale, err
}

func (l *Locales) Get(opt Option) string {
	resources := l.data[opt.Language]
	return resources.Get(opt)
}

func (i *Locales) Add(langs ...string) error {
	for _, lang := range langs {
		dirPath := filepath.Join(i.path, lang)
		dir, err := ioutil.ReadDir(dirPath)
		if err != nil {
			return err
		}
		langResources := make(Resources)
		for _, file := range dir {
			if file.IsDir() {
				continue
			}
			filePath := filepath.Join(dirPath, file.Name())
			resource, err := NewResources(filePath)
			if err != nil {
				return err
			}
			for kk, vv := range resource {
				langResources[kk] = vv
			}
		}
		i.data[lang] = langResources
	}
	return nil
}

type Resources map[string]Resource

func (r Resources) Get(opt Option) string {
	resource := r[opt.Prefix]
	return resource.Parse(opt.Key)
}

func NewResources(path string) (Resources, error) {
	out, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var data map[string]interface{}
	if err := yaml.Unmarshal(out, &data); err != nil {
		return nil, err
	}
	res := make(Resources)
	for k, v := range data {
		switch vv := v.(type) {
		case string:
			res[""] = map[string]string{k: vv}
		case interface{}:
			res[k] = newResource(vv)
		}
	}
	return res, err
}

type Resource map[string]string

func (r Resource) Parse(key string) string {
	return r[key]
}

func newResource(data interface{}) Resource {
	res := make(Resource)
	dataMap, ok := data.(map[interface{}]interface{})
	if !ok {
		return res
	}
	for k, v := range dataMap {
		var ck, cv string
		if ck, ok = k.(string); !ok {
			continue
		}
		if cv, ok = v.(string); !ok {
			continue
		}
		res[ck] = cv
	}
	return res
}
