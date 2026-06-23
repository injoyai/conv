# conv

`conv` 是一个偏宽松风格的 Go 类型转换工具库，提供基础类型转换、任意数据解析、层级数据访问和配置读取能力。

它的核心特点是：
- 常见类型转换直接可用
- 解析失败时通常返回零值或你提供的默认值
- 支持把任意数据继续按对象 / 数组路径读取

## 安装

```bash
go get github.com/injoyai/conv
```

## 基础转换

### 字符串转整数

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	var a string = "101"
	fmt.Println(conv.Int(a)) // 101
}
```

### 进制字符串转数字

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	fmt.Println(conv.Uint8("0xaa"))    // 170
	fmt.Println(conv.Uint8("0xff"))    // 255
	fmt.Println(conv.Int("0b01011"))   // 11
	fmt.Println(conv.Int("20"))        // 20
}
```

### 二进制字符串

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	fmt.Println(conv.BINStr(uint16(100))) // 0000000001100100
	fmt.Println(conv.BINStr(int8(100)))   // 01100100
}
```

### 布尔转换

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	fmt.Println(conv.Bool(3))   // true
	fmt.Println(conv.Bool(0))   // false
	fmt.Println(conv.Bool("开")) // true
}
```

## Unmarshal

`conv.Unmarshal` 可以把任意数据尽量解析到目标结构中，适合做 map、struct、slice 之间的宽松转换。

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	m := map[string]any{
		"a": 1,
		"b": "2",
		"c": true,
		"d": 1.02,
		"f": 30.6,
		"g": map[string]any{
			"h": "10",
		},
	}

	type Example struct {
		A int     `json:"a"`
		B string  `json:"b"`
		C bool    `json:"c"`
		D float64 `json:"d"`
		F string  `json:"f"`
		G struct {
			H int `json:"h"`
		} `json:"g"`
	}

	var x Example
	if err := conv.Unmarshal(m, &x); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", x)
}
```

## Map 深度读取

`conv.NewMap` 可以把字符串、对象、数组等数据包装成可继续深入读取的 `Map`。

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
)

func main() {
	data := `{"a":{"b":[0,1,2,3],"c":{"d":"d"}}}`

	m := conv.NewMap(data)

	fmt.Println(m.GetString("a.b[3]")) // 3

	m.Append("a.b", "4")
	fmt.Println(m.String()) // {"a":{"b":[0,1,2,3,"4"],"c":{"d":"d"}}}

	m.Set("a.c.d", []int{1, 2, 3})
	fmt.Println(m.String()) // {"a":{"b":[0,1,2,3,"4"],"c":{"d":[1,2,3]}}}

	m.Del("a.c.d[0]")
	m.Del("a.c.d[-1]")
	fmt.Println(m.String()) // {"a":{"b":[0,1,2,3,"4"],"c":{"d":[2]}}}
}
```

支持的默认路径格式：
- `a.b.c`
- `a.list[0]`
- `a.list[-1]`

默认使用 JSON 解析，也可以显式传入 codec，例如 YAML、TOML、INI。

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv"
	"github.com/injoyai/conv/codec"
)

func main() {
	yamlData := "name: yaml\nage: 18"
	m := conv.NewMap(yamlData, codec.Yaml)

	fmt.Println(m.GetString("name"))
	fmt.Println(m.GetInt("age"))
}
```

## cfg 配置读取

`cfg` 包用于按层级读取配置，默认会组合默认文件源和环境变量源。

```go
package main

import (
	"fmt"

	"github.com/injoyai/conv/cfg"
)

func main() {
	fmt.Println(cfg.GetInt("http.port"))
	fmt.Println(cfg.GetString("app.name"))
	fmt.Println(cfg.GetString("test[0].name"))
}
```

如果你需要自定义配置源，可以基于 `cfg.New(...)` / `cfg.Append(...)` 组合自己的读取方式。

## 常用 API

- 基础转换：`Int`、`Int64`、`Uint64`、`Float64`、`Bool`、`String`、`Bytes`
- 集合转换：`Strings`、`Ints`、`Int64s`、`Interfaces`
- 编码辅助：`BINStr`、`OCTStr`、`HEXStr`
- 数据解析：`Unmarshal`、`NewMap`、`DMap`

## 注意事项

- 这是偏宽松的转换库，很多场景下失败不会返回错误，而是返回零值。
- `Map` 的路径分隔符使用 `.` 和 `[]`，因此不适合直接访问包含这些字符的 key。
- `Uint32` / `Uint64` 对浮点数默认不是 IEEE754 bit pattern 转换；如果你需要位级转换，请显式使用标准库 `math.Float32bits` / `math.Float64bits`。

## License

MIT
