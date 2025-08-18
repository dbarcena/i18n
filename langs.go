package i18n

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"github.com/hjson/hjson-go/v4"
)

type Translation struct {
	Translation nodes_t `json:"translation"`
}

func (f *I18n) Langs(path string) error {

	if _, err := os.Stat(path); os.IsNotExist(err) {

		return err
	}

	files, err := os.ReadDir(path)
	if err != nil {

		return err
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		fileName := path + "/" + file.Name()

		if lang, err := load(fileName); err == nil {

			name := strings.TrimSuffix(file.Name(), filepath.Ext(file.Name()))

			f.langs[name] = lang
		} else {
			return err
		}
	}

	return nil
}

func load(fileName string) (*lang_t, error) {

	l := new(lang_t)

	file, e := os.ReadFile(fileName)
	if e != nil {
		return nil, e
	}

	var translation Translation
	if err := hjson.Unmarshal(file, &translation); err != nil {

		return nil, err
	}

	l.textos = make(textos_t)
	rebuild(l, "", translation.Translation)

	return l, nil
}

/*
func Default() *lang_t {
	return &lang_t{make(textos_t)}
}
*/

func rebuild(l *lang_t, key string, v interface{}) {

	tipo := reflect.TypeOf(v).Kind()

	if tipo == reflect.String {
		l.textos[key] = v.(string)
	} else if tipo == reflect.Map {
		iter := reflect.ValueOf(v).MapRange()
		for iter.Next() {

			k := iter.Key().Interface().(string)
			k_tmp := key
			if len(key) > 0 {
				k_tmp = key + "." + k
			} else {
				k_tmp = k
			}

			rebuild(
				l,
				k_tmp,
				iter.Value().Interface(),
			)
		}
	}
}

func (f *I18n) T(find string) string {

	if lang, ok := f.langs[f.lang]; ok {

		if v, ok := lang.textos[find]; ok {
			return v
		}
	} else {
		return f.lang + "." + find
	}
	return find
}
