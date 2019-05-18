package main

import "fmt"
import "regexp"

func main() {

	match, _ := regexp.MatchString("p[a-z]+)ch", "peach")
	fmt.Println(match)

	r, _ := regexp.Compile("p[a-z]+)ch")
	// 查找匹配字符串的
	fmt.Println(r.MatchString("peach"))
	// 查找匹配字符串的
	fmt.Println(r.FindString("peach punch"))
	// 查找第一个匹配的字符串，但返回的匹配开始和结束位置索引
	fmt.Println(r.FindStringIndex("peach punch"))
	// Submatch 返回完全匹配和局部匹配的字符串
	fmt.Println(r.FindStringSubmatch("peach punch"))
	// 返回完全匹配和局部匹配的索引位置
	fmt.Println(r.FindStringSubmatchIndex("peach punch"))
	// 带 All 的函数返回所有匹配项, 不仅仅是首次匹配项.
	fmt.Println(r.FindAllString("peach punch pinch", -1))
	fmt.Println(r.FindAllStringSubmatchIndex("peach punch pinch", -1))
	// 使用字符串作为参数, 使用MatchString这样的方法. 也提供[]byte参数并将String从函数命中去掉
	fmt.Println(r.Match([]byte("peach")))

}
