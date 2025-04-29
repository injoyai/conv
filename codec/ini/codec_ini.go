package ini

import (
	"bytes"
	"errors"
	"gopkg.in/ini.v1"
	"strings"
)

type Ini struct {
	LoadOptions ini.LoadOptions
}

func (i Ini) Marshal(v any) ([]byte, error) {
	m, ok := v.(map[string]map[string]string)
	if !ok {
		return nil, errors.New("type error: not map[string]map[string]string")
	}
	cfg := ini.Empty(i.LoadOptions)
	for k, v := range m {
		s := cfg.Section(k)
		for vk, vv := range v {
			s.Key(vk).SetValue(vv)
		}
	}
	buf := bytes.NewBuffer(nil)
	_, err := cfg.WriteTo(buf)
	return buf.Bytes(), err
}

func (i Ini) Unmarshal(data []byte, v any) error {

	m, ok := v.(*map[string]any)
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
			if section.Name() == ini.DefaultSection {
				break
			}
			if cm[k] == nil {
				x := map[string]any{}
				cm[k] = x
			}
			cm = cm[k].(map[string]any)
		}
		for _, k := range section.Keys() {
			cm[k.Name()] = k.Value()
		}
	}

	return nil
}
