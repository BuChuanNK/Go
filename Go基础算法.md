# Go 基础算法

## 目录

## 1. Sorting 排序
Go 的 sort 包实现了内置和用户自定义数据类型的排序功能。 

排序方法是正对内置数据类型的。注意排序是原地更新，所以会改变给定的序列且不返回一个新值。

    strs := []string{"c", "a", "b"}
	sort.Strings(strs)

    ints := []int{7, 2, 4}
	sort.Ints(ints)

同时可以使用 sort 来判断这个切片是否已经排序了。

    s := sort.IntsAreSorted(ints)

**使用函数自定义排序**

有时候我们想使用和集合的自然排序不同的排序方法，因此需要函数自定义排序。
1. 创建一个为内置 []string 类型的别名的 ByLength 类型
2. 实现 sort.Interface 的 Len, Less 和 Swap 方法。其中 Less 将控制实际的自定义排序逻辑. 
3. 通过将原始的切片转化为 ByLength, 最后使用 sort.Sort 来对这个类型进行排序。

    type ByLength []string

    func (s ByLength) Len() int {
        return len(s)
    }

    func (s ByLength) Swap(i, j int) {
        s[i], s[j] = s[j], s[i]
    }

    func (s ByLength) Less(i, j int) bool {
        return len(s[i]) < len(s[j])
    }

## 2. Panic
Panic 意味着有出乎意料的错误发生。 
通常用 Panic 来表示程序正常运行中不应该出现的, 或者没有处理好的错误。

Panic 的一个基本用法就是在一个函数返回了错误值(非零的状态码)。

    panic("a problem")

	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}

## 3. Defer
Defer 被用来确保一个函数调用在程序执行结束前执行。 
同样用来执行一些清理工作。 defer 用在像其他语言中的 ensure 和 finally 用到的地方。

例子: 32_Defer.go 中, 我们想要创建一个文件，向它进行写操作, 然后在结束时关闭它. 
在 closeFile 后得到一个文件对象, 使用 defer 通过 closeFile 来关闭这个文件. 这会在封闭函数 (main) 结束时执行, 就是 writeFile 结束后.

## 4. Collection Function 组合函数
我们经常需要程序在数据集上执行操作, 比如选择满足给定条件的所有项, 或者将所有的项通过一个自定义函数映射到一个新的集合上.

Go 不支持泛型, 在 Go 中, 当程序需要时, 通常是通过组合的方式来提供操作函数. 
