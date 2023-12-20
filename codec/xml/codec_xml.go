package xml

import (
	"encoding/xml"
)

type Xml struct{}

func (Xml) Marshal(v interface{}) ([]byte, error) {
	return xml.Marshal(v)
}

func (Xml) Unmarshal(data []byte, v interface{}) error {
	return xml.Unmarshal(data, v)
}
