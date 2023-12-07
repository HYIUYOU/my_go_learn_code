# Go学习文档

## ch1：入门

### 1.1 Hello, Word

[source](https://golang-china.github.io/gopl-zh/ch1/ch1-01.html)

```go
package main
//不需要分号
import "fmt"

func main() {
	fmt.Println("Hello, World!")
    //Go 语言原生支持 Unicode，它可以处理全世界任何语言的文本。
}
```

#### a.运行：go run 命令

```sh
go run helloword.go

#输出
(base) PS D:\GoProject\my_go_learn_code\ch1> go run helloword.go
Hello, World!
```

#### b.编译： go build 命令 

```shell
go build hellword.go #这种做法不可取，因为build是编译整个项目，而不是一个文件

go build main.go #编译整个文件
```

#### c.初始化：go init mod 命令

> Go在1.11版本之后引入了模块（modules）的概念，它需要一个`go.mod`文件来管理项目的依赖。如果你在没有`go.mod`文件的目录下运行Go命令

```shell
go init mod github.com/HYIUYOU/my_go_learn_code

#生成的go.mod文件
module github.com/HYIUYOU/my_go_learn_code
//module helloworld

go 1.21.5

```



### 1.2命令行参数

[soucre](https://golang-china.github.io/gopl-zh/ch1/ch1-02.html)

#### a.获取Args变量的方式 

- **os包中的Args变量**
- **os包外：os.Args**

> `os` 包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
>
> 程序的命令行参数可从 `os` 包的 `Args` 变量获取；`os` 包外部使用 `os.Args` 访问该变量。

#### b. 细节

- **`os.Args` 变量是一个字符串（string）的 *切片*（slice）**
- **用 `s[i]` 访问单个元素**
- **用 `s[m:n]` 获取子序列**(左闭右开形式，即，区间包括第一个索引元素，不包括最后一个) 
- **序列的元素数目为 `len(s)`**
- **`os.Args` 的第一个元素：`os.Args[0]`，是命令本身的名字**

> `os.Args` 变量是一个字符串（string）的 ==*切片*（slice）==（译注：slice 和 Python 语言中的切片类似，是一个简版的动态数组），切片是 Go 语言的基础概念，稍后详细介绍。现在先把切片 `s` 当作数组元素序列，序列的长度动态变化，用 `s[i]` 访问单个元素，用 `s[m:n]` 获取子序列（译注：和 Python 里的语法差不多）。序列的元素数目为 `len(s)`。和大多数编程语言类似，区间索引时，Go 语言里也采用左闭右开形式，即，区间包括第一个索引元素，不包括最后一个，因为这样可以简化逻辑。（译注：比如 `a=[1,2,3,4,5]`, `a[0:3]=[1,2,3]`，不包含最后一个元素）。比如 `s[m:n]` 这个切片，`0≤m≤n≤len(s)`，包含 `n-m` 个元素。
>
> `os.Args` 的第一个元素：`os.Args[0]`，是命令本身的名字；其它的元素则是程序启动时传给它的参数。
>
> `s[m:n]` 形式的切片表达式，产生从第 `m` 个元素到第 `n-1` 个元素的切片，下个例子用到的元素包含在 `os.Args[1:len(os.Args)]` 切片中。如果省略切片表达式的 `m` 或 `n`，会默认传入 `0` 或 `len(s)`，因此前面的切片可以简写成 `os.Args[1:]`。