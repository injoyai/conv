# conv 类型转换工具

- 用于各种类型数据的转换,出现错误(例如不可转),返回默认值



## 如何使用

- 下载安装

      go get github.com/injoyai/conv

---

## 如何使用conv

- 字符串转整数
  ```go
  
      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a string = "101"
    
        result := conv.Int(a)

        fmt.Println(result) // 得到结果 101

      }
  
  ```



- 十六进制转十进制
  ```go
  
      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a string = "x0aa"

        var b string = "x0ff"

        var c string = "0b01011"

        fmt.Println(conv.Uint8(a)) // 得到结果 170

        fmt.Println(conv.Uint8(b)) // 得到结果 255

        fmt.Println(conv.Int(c)) // 得到结果 11

      }
  
  ```


- 二进制字符串
  ```go

      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a uint16 = 100

        // 根据数字类型的位数 得到对应长度(8的倍数)的二进制(只包含0和1)字符串
        fmt.Println(conv.BinStr(a)) // 得到结果 "0000000001100100"

        fmt.Println(conv.BinStr(int8(b))) // 得到结果 "01100100"

      }
  ```

- 布尔类型
  ```go
  
        pakeage main
  
        import "github.com/injoyai/conv"
  
        func main(){
          
          var a int = 3
        
          // 数字类型 0 为false ,其余为true
          fmt.Println(conv.Bool(a)) // 得到结果 true
  
          a = 0
  
          fmt.Println(conv.Bool(a)) // 得到结果 false
  
          var b string = "开"
  
          fmt.Println(conv.Bool(b)) // 得到结果 true
  
        }
        
  ```

## 如何解析任意类型

- conv.Unmarshal(a,b)能解析任意数据a到b

  ```go
  
  package main
  
  import "github.com/injoyai/conv"
  
  func main(){
    m := map[string]interface{}{
          "a": 1,
          "b": "2",
          "c": true,
          "d": 1.02,
          "e": 20.1,
          "f": 30.6,
          "G": map[string]interface{}{
              "h": "10",
          },
      }
      type _struct struct {
          A int     `json:"a"`
          B string  `json:"b"`
          C bool    `json:"c"`
          D float64 `json:"d"`
          E float64
          F string `json:"f"`
          G struct {
              H int `json:"h"`
          }
      }
	  
      x := new(_struct)
      if err := conv.Unmarshal(m, x); err != nil {
             fmt.Println(err)
             return
      }
  
      //得到 _struct{A:1, B:"2", C:true, D:1.02, E:0, F:"30.6", G:struct { H int "json:\"h\"" }{H:10}} 
      fmt.Println("%#v", *x) 
	  
  }

  ```

## 如何使用cfg

- cfg包是读取配置,能够读取到每一层级,默认解析JSON,读取位置(./config/config.json)

  ```go

      package main

      import (
          "github.com/injoyai/conv"
          "github.com/injoyai/conv/codec"
	     )

      func main(){

        /*
          默认读取配置路径 ./config/config.json 假设内容如下:
          {
          "http":{
            "port":8000
           },
          "tcp":{
            "port":9000
           },
          "test":[
            {
             "name":"injoy"
            }
           ]
          }


        */
       
        fmt.Println(cfg.GetInt("http.port")) // 得到结果 8000

        fmt.Println(cfg.GetInt("tcp.port")) // 得到结果 9000

        fmt.Println(cfg.GetInt("test[0].name")) // 得到结果 "injoy"
  
        //自定义yaml配置文件读取
        c :=cfg.WithPath("./xxx.yaml", codec.Yaml)
		
        c.GetString("http.port")

      }
	  
  ```

## 技术支持



## 获取更多信息



