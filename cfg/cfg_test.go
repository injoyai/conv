package cfg

import (
	"github.com/injoyai/conv/codec"
	"testing"
)

func TestPrint(t *testing.T) {
	if Default.Err() != nil {
		t.Error(Default.Err())
	}
	//t.Log(Default.Map.String())
	t.Log(GetString("name"))
	t.Log(GetInt("age"))
	t.Log(GetString("child"))
	t.Log(GetString("child[0].name"))
	t.Log(GetInt("child[0].age"))
	t.Log(GetString("child[1].name"))
	t.Log(GetInt("child[1].age"))
	t.Log(GetString("child[2].name", "null"))
	t.Log(GetInt("child[2].age", -1))
}

func TestJson(t *testing.T) {
	Default = New("../testdata/config.json")
	TestPrint(t)
}

func TestYaml(t *testing.T) {
	Default = New("../testdata/config.yaml", codec.Yaml)
	TestPrint(t)
}

func TestToml(t *testing.T) {
	Default = New("../testdata/config.toml", codec.Toml)
	TestPrint(t)
}

// 测试未通过,待完成
func TestXml(t *testing.T) {
	Default = New("../testdata/config.xml", codec.Xml)
	TestPrint(t)
	t.Log(GetString("user.name"))
	t.Log(GetString("name"))
	t.Log(GetString("project.component"))
	t.Log(GetString("component"))
}
