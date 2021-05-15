# Go语言学习
1. 开始创建文件目录

学习教材来自[李文周的博客](https://www.liwenzhou.com/posts/Go/go_menu/)

- GO 语言中的常量


#### keyword 关键字 25个

> `default break func interface struct select`               
> ` case defer go map struct`                    
> `chan else goto package switch `            
> `const fallthrougth if raenge type `                
> `continue for import return var `               

#### 保留字
 
 | ***Constants***| ***Types*** | ***Functions*** |
 | :-: | :-: | :-: |
 | true |  int  int8  int16  int32  int64 | make  len  cap  new  append  copy  close  delete|
 | false |uint  uint8  uint16  uint32  uint64  uintptr|complex  real  imag|
 | iota | float32  float64  complex128  complex64| panic  recover| 
 | nil |bool  byte  rune  string  error| |
 
* 变量： 
  *计算机中数据是存在内存当中的，当我们想改变数据时往往需要去操作数据的内存地址，但在实际代码过程中，频繁地操作内存地址既不美观并且也容易出错，所以就用变量把内存地址保存起来。*
* 变量类型：
  *一般来讲，编程语言的变量类型可以分为 **整形**， **浮点型**， **布尔型***
> #### 变量申明
 * 标准声明 ：
  > `var`/ `const` 变量名 变量类型
   `var name string`            
   ` var age int`     
   `var isOk bool`        
  * 批量声明 ：

  > var(       
  >  a string   
  >  b int      
  >  c bool     
  >
  >)

  >__nil__      代表一个控制 字符串就是 "" int 就是 0 


  `iota`在 **`const`**<u>关键字出现时将被重置为0<u>。
  **const中每新增一行常量声明将使iota计数一次(iota可理解为const语句块中的行索引)。 使用iota能简化定义，在定义枚举时很有用**