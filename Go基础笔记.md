# Go 语言基础

## Go 简介
> Go 由 Google 开发，用于建立简单，快速，可靠软件的程序设计语言
> 
> 参考书籍：《Go by Example》

## 目录
[1. Hellor World](#1-%E4%BD%A0%E5%A5%BD%E4%B8%96%E7%95%8C)   
[2. 值](#2-%E5%80%BC)   
[3. 变量](#3-%E5%8F%98%E9%87%8F)    
[4. 常量](#4-%E5%B8%B8%E9%87%8F)    
[5. For 循环](#5-for-%E5%BE%AA%E7%8E%AF)    
[6. if/else 分支](#6-ifelse-%E5%88%86%E6%94%AF)     
[7. Switch 分支结构](#7-switch-%E5%88%86%E6%94%AF%E7%BB%93%E6%9E%84)    
[8. Array 数组](#8-array-%E6%95%B0%E7%BB%84)    
[9. Slice 切片](#9-slice-%E5%88%87%E7%89%87)    
[10. Map 关联数组](#10-map-%E5%85%B3%E8%81%94%E6%95%B0%E7%BB%84)    
[11. Range 遍历](#11-range-%E9%81%8D%E5%8E%86)  
[12. Func 函数](#12-function-%E5%87%BD%E6%95%B0)    
[13. Pointer 指针](#13-pointer-%E6%8C%87%E9%92%88)  
[14. Struct 结构体](#14-struct-%E7%BB%93%E6%9E%84%E4%BD%93)     
[15. Interface 接口](#15-interface-%E6%8E%A5%E5%8F%A3)  
[16. Error 错误处理](#16-error-%E9%94%99%E8%AF%AF%E5%A4%84%E7%90%86)    


## 1. 你好，世界
HelloWorld测试代码：

    package main
    
    import "fmt"
    func main(){
        fmt.Println("hello world!")
    }

如何运行HelloWorld.go文件：

    直接运行
    $ go run HelloWorld.go
    编译为二进制文件
    $ go build HelloWorld.go
    $ ./HelloWorld

## 2. 值
Go 拥有的值类型，包括String ,int float, boolean

Var.go测试代码：

    package main

    import "fmt"
    func main(){
        fmt.Println("go"+"lang")

        fmt.Println("1+1=", 1+1)
        fmt.Println("7.0/3.0 = ", 7.0/3.0)

        fmt.Println(true && false)
        fmt.Println(true || false)
        fmt.Println(!true)
    }


## 3. 变量
使用var来显式声明一个或多个变量。
Go将自动推断已经初始化的变量类型。
:=语句是声明并初始化变量的简写。

    var a string = "initial"
	fmt.Println(a)

	var b,c int = 1,2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "short"
	fmt.Println(f)

## 4. 常量
Const用于声明一个常量。
当上下文需要时，一个常熟可以被给定一个类型。

    package main

    import "fmt"
    import "math"

    const s string  = "constant"

    func main()  {
        fmt.Println(s)

        const n = 500000000

        const d = 3e20/n
        fmt.Println(d)

        fmt.Println(int64(d))

        fmt.Println(math.Sin(n))
	}

## 5. For 循环
和Java写法差不多, 
但不带条件的for循环将会一直执行，知道在循环体内使用了break或者return跳出循环. 

    i:= 1
	for i<=3{
	    fmt.Println(i)
	    i = i+1
	}

	for j:=7; j<=9; j++{
	    fmt.Println(j)
	}

	for{
	    fmt.Println("loop")
	    break
	}

## 6. if/else 分支
和Java写法差不多,
在条件语句之前可以有一个语句：仍和在这里声明的遍历啊那个都可以在所有条件分支中使用. 

    if 7%2==0{
		fmt.Println("7 is even")
	}else{
		fmt.Println("7 is odd")
	}

	if 8%4==0{
		fmt.Println("8 is divisible by 4")
	}

	if num := 9; num<0{
		fmt.Println(num, "is negative")
	} else if num < 10{
		fmt.Println(num, "has 1 digit")
	} else{
		fmt.Println(num, "has multiple digits")
	}

## 7. Switch 分支结构
依旧和Java写法差不多

    switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("it's the Weekend")
	default:
		fmt.Println("it's a Weekday")
	}

	t:=time.Now()
	switch{
	case t.Hour()<12:
		fmt.Println("it's before noon")
	default:
		fmt.Println("it's after noon")
	}

## 8. Array 数组
在Go中，数组是一个固定长度的数列。

使用 array[index] = value 语法来设置数组制定位置的值。
使用内置函数len返回数组的长度。

    var a [5]int
    a[4] = 100
    b:=[5]int{1,2,3,4,5}

    var twoD [2][3]int
	for i:=0; i<0; i++{
		for j:=0; j<3; j++{
			twoD[i][j] = i+j
		}
	}

但是在Go程序中，相对于Array而言，Slice使用的更多。

## 9. Slice 切片
非常重要！！！
Slice 是 Go 中一个关键的数据类型，是一个比数组更加强大的序列接口。
使用内建方法make来创建slice。

    s := make([]string, 3)

通过内置方法append，可以返回一个包含来一个或者多个新值的slice

    s = append(s, "e", "f")

同时 Slice 可以被 Copy.

    c := make([]string, len(s))
	copy(c, s)

Slice 也可以构建多维数据结构. 内部的Slice长度可以不同.

    # 得到[[0] [1 2] [2 3 4]]
    twoD := make([][]int, 3)
	for i:=0; i<3; i++{
		innerLen:=i+1;
		twoD[i] = make([]int, innerLen)
		for j:=0; j<innerLen; j++{
			twoD[i][j] = i+j
		}
	}

## 10. Map 关联数组
map 是 Go 内置关联数据类型。
要创建一个空 map, 需要使用内建的 make:make(map[key-type]val-type).
使用典型的 make[key]=val 语法来设置键值对.
使用 name[key] 来获取一个键的值.

    m := make(map[string]int)
    m["k1"] = 7
    v1 := m["k1"]

内建的 delete 可以从一个 map 中移除键值对.

    delete(m, "k2")

当从一个 map 中取值时, 可选的第二个返回值指示这个键是在这个 map 中.
这可以用来消除键不存在和键有零值, 像0或者""而产生的歧义.

    _, prs := m["k2"]

可以在同一行申明和初始化一个新的 map.

    n := map[string]int{"foo": 1, "bar": 2}

## 11. Range 遍历
range 迭代各种各样的数据结构. 
可以使用 range 来统计一个 slice 的元素个数. 
range 在 Array 和 Slice 中都同样提供每个项的索引和值. 

    for _, num := range nums{
	    sum += num
    }

range 在 map 中迭代键值对. 

    kvs := map[string]string{"a":"apple", "b":"banana"}
	for k, v := range kvs {
	    fmt.Printf("%s -> %s\n", k, v)
	}

range 在字符串中迭代 unicode 编码. 第一个返回值是 string 的起始字节位置, 第二个是 string 自己.

    for i, c := range "go"{
	    fmt.Println(i, c)
	}

## 12. Function 函数
函数是 Go 的中心. 
Go函数需要声明确的返回值, 括号中是输入值数据类型.

    func plus(a int, b int) int{
	    return a + b
    }

然后通过 name(args) 来调用一个函数.

    res := plus(1, 2)

**多返回值**
Go 内建*多返回值*支持。

    func vals() (int, int){
	    return 3, 7
    }

**变参函数**
可变参数函数. 可以用任意数量的参数调用.

    // 该函数可以使用任何数目的 int 作为参数
    func sum(nums ...int){
        fmt.Println(nums, " ")
        total := 0
        for _, num := range nums{
            total += num
        }
        fmt.Println(total)
    }

**闭包**
Go 支持通过*闭包*来使用*匿名函数*。
匿名函数: 定义一个不需要命名的内联函数。

    // intSeq()函数，返回值为一个内部定义的匿名函数。
    // 该返回的函数使用闭包的方式隐藏变量i。
    func intSeq() func() int{
        i := 0
        return func() int {
            i += 1
            return i
        }
    }

在调用 intSeq 函数时，将返回值赋给nextInt。
这个函数的值包含了自己的值 i, 这样在每次调用nextInt时都会更新 i 的值。
    
    nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())

**递归**
Go支持递归. 

    func fact(n int) int{
        if n == 0{
            return 1
        }
        return n*fact(n-1)
    }

## 13. Pointer 指针
Go 支持*指针*, 允许在程序中通过引用传递至或者数据结构。
声明一个指针，同时需要声明该指针的数据类型：

    // 指定一个int类型的指针
    iptr *int

对于一个指针*解引用*：

    *iptr = 0

利用 *&i* 来获得i的内存地址。

    fmt.Println("pointer: ", &i)

## 14. Struct 结构体
Go 的结构体是各个字段的类型的集合。可以有多重声明方式。

    person{"Bob", 20}
    person{name: "Alice", age: 30}
    person{name: "Fred"}

也可以对结构体声明指针

    sp := &s
	fmt.Println(sp.age)

**结构体方法**
同时可以在结构体中定义*方法*
函数的输入类型可以是 值, 也可以是 指针。

    func (r *rect) area() int {
        return r.width * r.height
    }

    func (r rect) perim() int {
        return 2*r.width +2*r.height
    }

Go 自动处理方法调用时的值和指针之间的转化。
可以使用指针来调用方法，避免在方法调用时产生一个拷贝, 或让方法能够改变接受的数据. 

## 15. Interface 接口
在 Go 语言中，接口为方法特征的命名集合. 

    type geometry interface {
        area() float64
        perim() float64
    }

要实现一个接口，只需要实现接口中的所有方法。

    type rect struct {
        width, height float64
    }

    func (r rect) area() float64{
        return r.width * r.height
    }

    func (r rect) perim() float64{
        return 2*r.width + 2*r.height
    }

如果一个变量是接口类型，那么可以调用这个接口中的方法。
比如下面这个通用的 measure 函数，利用这个特性，可以用到任何 geometry 上。

    func measure(g geometry){
        fmt.Println(g)
        fmt.Println(g.area())
        fmt.Println(g.perim())
    }

凡是实现了 geometry 接口的结构体类型，都可以作为 measure 的参数。

    r := rect{width: 3, height: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)

## 16. Error 错误处理
Go 语言使用一个独立的，明确的返回值来传递错误信息的.
Go 语言的处理方式能清楚地知道哪个函数返回了错误, 并能像调用那些没有出错的函数一样调用. 

按照惯例, 错误通常是最后一个返回值并且是 error 类型, 一个内建的接口. 
errors.New 构造一个使用给定的错误信息的基本 error 值. 
返回错误值为 nil 代表没有错误. 

    func f1(arg int) (int, error){
        if arg == 42{
            return -1, errors.New("can't work with 42")
        }

        return arg+3, nil
    }

通过实现 Error 方法来自定义 error 类型是可以的. 这里使用自定义错误类型来表示上面的参数错误. 

    func (e *argError) Error() string{
        return fmt.Sprintf("%d - %s", e.arg, e.prob)
    }

可以使用 &argError 语法来简历一个新的结构体, 并提供来 arg 和 prob 这两个字段的值. 

    func f2(arg int) (int, error){
        if arg == 42{
            return -1, &argError{arg, "can't work with it"}
        }
        return arg+3, nil
    }

