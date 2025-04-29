package xml

import (
	"encoding/xml"
)

type Xml struct{}

func (Xml) Marshal(v any) ([]byte, error) {
	return xml.Marshal(v)
}

func (Xml) Unmarshal(data []byte, v any) error {
	return xml.Unmarshal(data, v)
}
