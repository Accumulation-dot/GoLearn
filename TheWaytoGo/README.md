# 入门


### 记录一些常用的内容

#### 字符串
#####1.前缀和后缀

```
strings.HasPrefix(s, prefix string) bool
strings.HasSuffix(s, suffix string) bool

```

#####2.字符串包含关系
```
strings.Contains(s, substr string) bool
```

#####3.判断字符串或者字符在父字符串中出现的位置（索引）
```$xslt
// 如果返回-1表示不包含
strings.Index(s, str string) int
// 最后出现的索引 -1表示不包含
strings.LastIndex(s, str string) int
// 非ASCII编码使用这个方法 返回结果同上
strings.IndexRune(s, str string) int 
```

#####4.字符串替换
```$xslt
// 替换str中出现的old 为new (前n个)
// 如果n = -1 全部替换 
strings.Replace(str, old, new, n) string
```

#####5.统计字符串出现次数
```$xslt
// str 在s中出现的次数
strings.Count(s, str string) int
```

#####6.重复字符串
````$xslt
// 用来重复count个s字符串拼接到一起
strings.Repeat(s, count int) string
````
#####7.修改字符串大小写
```$xslt
// 转小写
stirngs.ToLower(s) string
// 转小写
strings.ToUpper(s) string
```

#####8.删除制定的字符串
```$xslt
// 删除前后的空白符号
strings.TrimSpace(s)
// 将前后出现的 trim 删除
strings.Trim(s, trim string)
// 移除左侧出现的trim
strings.TrimLeft(s, trim string)
// 移除出现在右侧的trim
strings.TrimRight(s, trim string)
```

#####9.分割字符串
```$xslt
// 利用一个或者多个空白字符串 进行分割
strings.Fields(s) []string
// 用自定义的字符串分割
strings.Split(s, sep) []string
```

#####10.拼接
```$xslt
// 拼接
strings.Join(sl []string, sep string) string
```

#####11.从字符串读取内容
```$xslt
strings.NewReader(str) //创建一个Reader

```

#####12.字符串转换

    任何类型转换为字符串总是成功的
    但是字符串转其他类型可能会抛出错误
    
    // 转int类型
    strconv.Atoi(s string) (i int, err error)
    
    // 转float64
    strconv.ParseFloat(s string, bitSize int) (f float64, err error)

#### time使用

#####1.常用函数
    
    // 当前时间
    time.Now()
    
    // UTC 格式
    time.Now().UTC()
    
    t.Day() 
    t.Month()
    t.Year()
    
    t.Hour()
    t.Minute()
    t.Second()
    
    t.Format("02 Jan 2006 15:04")
    
    
    t.Format(time.RFC822()) utc
    
#### 内置函数

#####1.常用函数
    
    close
    len
    cap
    new
    make
    copy
    append
    panic
    recover
    print
    println
    complex
    real
    imag   
    
#####2.sorter 排序

    实现Sort接口
    3个函数分别为
    Len()
    Less(i, j int) bool
    Swap(i, j int) //i, j 为index
    
    
    type 
    
    type IntArray []int // 重命名类型
         
#####3接口继承
    
    声明一个要继承的类型的一个成员 *type就可以调用 type的方法
             