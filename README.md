# conv 类型转换工具

- 用于各种类型数据的转换,出现错误(例如不可转),返回默认值



## 如何使用

- 下载安装

      go get github.com/injoyai/conv

---

## 如何使用conv

- 字符串转整数

      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a string = "101"
    
        result := conv.Int(a)

        fmt.Println(result) // 得到结果 101

      }




- 十六进制转十进制

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

- 字节和浮点的转换

      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a float32 = 120.01
    
        // 如果输入类型是float32或者float64,并且期待的类型是uint32或者是uint64 
        // 则根据IEEE二进制浮点数算术标准（IEEE 754）来转换
        result := conv.Uint32(a)

        fmt.Println(result) // 得到结果 1123026207

        var b float32 = 120.01

        // 等同于strconv.Atoi
        fmt.Println(conv.Int(b)) // 得到结果 120

      }

- 二进制字符串

      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a uint16 = 100

        // 根据数字类型的位数 得到对应长度(8的倍数)的二进制(只包含0和1)字符串
        fmt.Println(conv.BinStr(a)) // 得到结果 "0000000001100100"

        fmt.Println(conv.BinStr(int8(b))) // 得到结果 "01100100"

      }

- 布尔类型

      pakeage main

      import "github.com/injoyai/conv"

      func main(){
        
        var a int = 3
      
        // 数字类型 0 为false ,其余为true
        fmt.Println(conv.Bool(a)) // 得到结果 true

        a = 0

        fmt.Println(conv.Bool(a)) // 得到结果 false

        var b string = 

      }


## 如何使用cfg

- cfg包是读取配置,目前支持只json,能够读取到每一层级,默认读取位置(./config/config.json)

      package main

      import "github.com/injoyai/conv"

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

      }


## 技术支持



## 获取更多信息



