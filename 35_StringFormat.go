package main

import "fmt"

type point struct {
	x, y int
}

func main() {

	p := point{1, 2}
	fmt.Printf("%v\n", p)
	// 如果值是一个结构体, %+v 的格式化输出内容将包括结构体的字段名
	fmt.Printf("%+v\n", p)
	// %#v 形式则输出这个值的Go语法表示
	fmt.Printf("%#v\n", p)
	// 需要打印值的类型, 使用%T
	fmt.Printf("%T\n", p)
	// 格式化布尔值
	fmt.Printf("%t\n", true)
	// 格式化整形数, 使用%d进行标准的十进制格式化
	fmt.Printf("%d\n", 123)
	// 输出二进制表达形式
	fmt.Printf("%b\n", 14)
	// 输出给定整数的对应字符
	fmt.Printf("%x\n", 456)
	// 使用%f进行最基本的十进制格式化
	fmt.Printf("%f\n", 78.9)
	// %e 和 %E 将浮点型格式化为科学计数法表示形式
	fmt.Printf("%e\n", 123400000.0)
	fmt.Printf("%E\n", 123400000.0)
	// 使用%s进行基本的字符串输出
	fmt.Printf("%s\n", "\"string\"")
	// 使用%q, 像Go源代码中那样带有双引号的输出
	fmt.Printf("%q\n", "\"string\"")
	// %x输出使用base-16编码的字符串, 每个字节使用两个字符表示
	fmt.Printf("%x\n", "hex this")
	// %p 输出一个指针
	fmt.Printf("%p\n", &p)

}
