package ini

import (
	"gopkg.in/ini.v1"
	"strings"
)

type Ini struct {
	LoadOptions ini.LoadOptions
}

func (i Ini) Marshal(v interface{}) ([]byte, error) {
	//todo
	return nil, nil
}

func (i Ini) Unmarshal(data []byte, v interface{}) error {

	m, ok := v.(*map[string]interface{})
	if !ok {
		return nil
	}

	cfg := ini.Empty(i.LoadOptions)

	err := cfg.Append(data)
	if err != nil {
		return err
	}

	for _, section := range cfg.Sections() {
		cm := *m
		for _, k := range strings.Split(section.Name(), ".") {
			if section.Name() == "DEFAULT" {
				break
			}
			if cm[k] == nil {
				x := map[string]interface{}{}
				cm[k] = x
			}
			cm = cm[k].(map[string]interface{})
		}
		for _, k := range section.Keys() {
			cm[k.Name()] = k.Value()
		}
	}

	return nil
}
