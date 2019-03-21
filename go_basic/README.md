# Go-Basic

##1.变量的声明

    var power1 int = 1000
    power2 := 1000
    power3, power4 := 2000, 3000
    var power5, power6 int  = 4000, 5000
    	
##2.函数机制

    func log(message string)  {
    	
    }
    
    func add(a, b int) int {
    	return a + b
    }
    
    func add1(a, b int) (c int)  {
    	c = a + b
    	return
    }
    
    func add2(a, b int) (c int, compare bool)  {
    	c = a + b
    	compare = a > b
    	return
    }
    
    func add3(a, b int) (int, bool)  {
    	c := a + b
    	compare := a > b
    	return c, compare
    }
    调用
    c, _ := add3(100, 200) // _ 丢弃值
    fmt.Println(c)
    
##3.结构使用：声明和初始化

    type Person struct {
    	Name string
    	Age int
    	Sex int
    }
    
    person := Person {
        Name: "name",
        Age:  10,
        Sex:  1,
    }
    
    person.Sex = 0
    
    如果想修改函数中传入参数的值 请使用 *Person 类型
    
    
    func (p *Person) show() string {
        return p.Name
    }
    
##4.new使用
    
	person1 := new(Person)
	// or
	person2 := &Person{}
	
##5.访问Struct里的隐式属性的方法
    
    type Person struct {
    	Name string
    	Age int
    	Sex int
    }
        
    func (p *Person) showSex() {
    	fmt.Println(p.Sex)
    }
    
    type Girl struct {
    	*Person
    }

    girl := &Girl {
    	Person: &Person {
    		Sex: 1,
    	},
    }
    
    girl.showSex()
    // or
    girl.Person.showSex()
    
##6.使用值还是指针

    1.局部变量赋值
    2.结构体指针
    3.函数返回值
    4.函数参数
    5.方法接收器
    
##7.数组

    var array [10]int //最大可以装10个数字
    array[0] = 100
    ...
    
    array := [10]int{100, 100, 100, 200}
    // for循环
    for index,value := range array {
        // index 索引 
        // value 值
    }
    
##8.切片 动态数组
    
    score := []int{100, 200, 300, 400} //没有标注数字多少 就是可以动态添加
    score := make([]int, length)
    score := make([]int, length, cap) // 预估的容量
    
##9.map

    map[string]int // key 是string value 是int
    
    mapvalue := map[string]int {
        "key1": value1,
        "key2": value2,
    }
    
    mapvalue := make(map[string]int)
    
##10.指针类型和值类型

    使用指针或者值类型 决定在于如何使用这个值 只读 值类型 操作使用 指针类型
    
##11.接口 interface 
    
    提供方法
    使用struct 实现接口的方法
    type DemoInterface interface {
        func show() string
    }
   
    type Demo struct {
   
    }
   
    func (ds *DemoStruct) show() string {
        // 
    }
    
##12.错误处理
    
    创建错误
    errors.New("error description")
    
##13.Defer使用

    处理一些需要时候处理的事情
    
##14.空接口

    interface{}
    
    空接口转换
    
    a interface{}
    a.(int) // int
    
    
    a.(type) 可以获取类型
    
##15.字符串和字节数组

    获取字符串长度 需要将字符串转变成字节数组再去获取 直接获取得到的是
    Unicode码点runes构成的字符串
    
    stra := "this is a string"
    bytes := []byte(stra)
    strb := string(bytes)
    
##16函数类型

    ***函数是一种类型***

##17 协程

    go func() {
    
    }()
    
    func process() {
    }
    
    go process()
    
##18 同步
    
    
    lock sync.Mutex
    
    lock.Lock()
    
    lock.Unlock()
    
    rw sync.RWMutex // 读写锁
    
##19 chan
   ***难***
   
    比较难 看不大明白什么时候需要用到这个东西
        
    