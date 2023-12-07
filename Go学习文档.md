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

#### a. 运行：go run 命令

```sh
go run helloword.go

#输出
(base) PS D:\GoProject\my_go_learn_code\ch1> go run helloword.go
Hello, World!
```

#### b. 编译： go build 命令 

```shell
go build hellword.go #这种做法不可取，因为build是编译整个项目，而不是一个文件

go build main.go #编译整个文件
```

#### c. 初始化：go init mod 命令

> Go在1.11版本之后引入了模块（modules）的概念，它需要一个`go.mod`文件来管理项目的依赖。如果你在没有`go.mod`文件的目录下运行Go命令

```shell
go init mod github.com/HYIUYOU/my_go_learn_code

#生成的go.mod文件
module github.com/HYIUYOU/my_go_learn_code
//module helloworld

go 1.21.5

```



### 1.2 命令行参数

[soucre](https://golang-china.github.io/gopl-zh/ch1/ch1-02.html)

#### a.获取Args变量的方式 

- **os包中的Args变量**
- **os包外：os.Args**

> `os` 包以跨平台的方式，提供了一些与操作系统交互的函数和变量。
>
> 程序的命令行参数可从 `os` 包的 `Args` 变量获取；`os` 包外部使用 `os.Args` 访问该变量。

#### b. code

```go
// Echo1 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    var s, sep string
    for i := 1; i < len(os.Args); i++ {
        s += sep + os.Args[i]
        sep = " "
    }
    fmt.Println(s)
}

```

```shell
$ go run .\main.go s asdfasf
$ s asdfasf
```

> 其中：
>
> os.Args[0] : 命令本身的名字													 “C:\Users\cloud\AppData\Local\Temp\go-build1381150014\b001\exe\main.exe” 
>
> os.Args[1] :  “s”
>
> os.Args[2] : “asdfasf”

#### c. 细节

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



>自增语句 `i++` 给 `i` 加 `1`；这和 `i+=1` 以及 `i=i+1` 都是等价的。对应的还有 `i--` 给 `i` 减 `1`。
>
>==它们是语句，而不像 C 系的其它语言那样是表达式。==
>
>所以 `j=i++` 非法，而且 `++` 和 `--` 都只能放在变量名后面，因此 `--i` 也非法。



#### d. for循环

> **Go 语言只有 `for` 循环这一种循环语句。**

**形式1： for loop**

```go
for initialization; condition; post {
    // zero or more statements
}
```

> `for` 循环三个部分不需括号包围。大括号强制要求，左大括号必须和 *`post`* 语句在同一行。
>
> *`initialization`* 语句是可选的，在循环开始前执行。*`initalization`* 如果存在，必须是一条 *简单语句*（simple statement），即，短变量声明、自增语句、赋值语句或函数调用。
>
> `condition` 是一个布尔表达式（boolean expression），其值在每次循环迭代开始时计算。如果为 `true` 则执行循环体语句。
>
> `post` 语句在**循环体执行结束后**，之后再次对 `condition` 求值。`condition` 值为 `false` 时，循环结束。
>
> for 循环的这三个部分每个都可以省略，如果省略 `initialization` 和 `post`，分号也可以省略：变成while循环



**形式2： while loop**

```go
// a traditional "while" loop
for condition {
    // ...
}
```

> 如果连 `condition` 也省略了，像下面这样：

**形式3：无限循环 infinite loop**

```go
// a traditional infinite loop
for {
    // ...
}
```

> 这就变成一个无限循环，尽管如此，还可以用其他方式终止循环，**如一条 `break` 或 `return` 语句**。

#### e.range 用法

- 每次循环迭代，`range` 产生一对值 : **(索引，值)** 

  > 注意：这里的索引永远从0开始，也就是重新加的索引，而不是原本的索引
  >
  > ```go
  > func main() {
  > 	for index, sep := range os.Args[1:] {
  > 		fmt.Printf("%s, %d\n", sep, index)
  > 	}
  > }
  > ```
  >
  > 这里的输出
  >
  > ```shell
  > $ go run hw12.go s asdfasf
  > $ s, 0
  > $ asdfasf, 1
  > ```
  >
  > index从0开始，但是os.Args[1:] 从1开始

- `range` 的语法要求，要处理元素，**必须处理索引**

  > - Go 语言中这种情况的解决方法是用 *空标识符*（blank identifier），即 `_`（也就是下划线）
  >
  > - **空标识符**可用于在任何语法需要变量名但程序逻辑不需要的时候（如：在循环里）丢弃不需要的循环索引，并保留元素值。

> `for` 循环的另一种形式，在某种数据类型的区间（range）上遍历，如字符串或切片。

```go
// Echo2 prints its command-line arguments.
package main

import (
    "fmt"
    "os"
)

func main() {
    s, sep := "", ""
    for _, arg := range os.Args[1:] {
        s += sep + arg
        sep = " "
    }
    fmt.Println(s)
}
```

> 这个例子不需要索引，但 `range` 的语法要求，要处理元素，必须处理索引。一种思路是把索引赋值给一个临时变量（如 `temp`）然后忽略它的值，但 Go 语言不允许使用无用的局部变量（local variables），因为这会导致编译错误。
>
> Go 语言中这种情况的解决方法是用 *空标识符*（blank identifier），即 `_`（也就是下划线）。
>
> **空标识符**可用于在任何语法需要变量名但程序逻辑不需要的时候（如：在循环里）丢弃不需要的循环索引，并保留元素值。大多数的 Go 程序员都会像上面这样使用 `range` 和 `_` 写 `echo` 程序，因为隐式地而非显式地索引 `os.Args`，容易写对。

#### f. 变量的声明

> 下面这些都等价：
>
> 实践中一般使用前两种形式中的某个，初始值重要的话就显式地指定变量的值，否则指定类型使用隐式初始化。

```go
s := ""  //是一条短变量声明，最简洁，但只能用在函数内部，而不能用于包变量。

var s string // 依赖于字符串的默认初始化零值机制，被初始化为 ""

var s = "" // 用得很少，除非同时声明多个变量

var s string = "" // 显式地标明变量的类型，当变量类型与初值类型相同时，类型冗余，但如果两者类型不同，变量类型就必须了。
```

#### g. Join 的用法

> `+=` 连接原字符串、空格和下个参数，产生新字符串，并把它赋值给 `s`。`s` 原来的内容已经不再使用，将在适当时机对它进行垃圾回收。
>
> 如果连接涉及的数据量很大，这种方式代价高昂。一种**简单且高效**的解决方案是使用 `strings` 包的 `Join` 函数：
>
> `strings.Join` 是 Go 语言中的一个字符串操作函数，用于将字符串切片（`[]string`）中的元素连接成一个单独的字符串

```go
func Join(elems []string, sep string) string
```

- `elems`: 一个字符串切片，包含需要连接的字符串元素。
- `sep`: 作为分隔符的字符串，用于将元素连接起来。

`Join` 函数会返回一个包含所有元素的字符串，元素之间用分隔符 `sep` 分隔。

```go
package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(strings.Join(os.Args[1:], "+"))
	//fmt.Println(os.Args[1:]) 
    //fmt.Print(reflect.TypeOf(os.Args[1:])) //返回os.Args[1:]的类型
}
```

**输出**

```shell
$ go run .\main.go s asdfasf
$ s+asdfasf
```

```go
fmt.Println(os.Args[1:]) 
```

> 这条语句的输出结果跟 `strings.Join` 得到的结果很像，只是被放到了一对方括号里。切片都会被打印成这种格式。



### 1.3 查找重复的行

```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() { //只要有输入就不会停
		counts[input.Text()]++
	}
	// NOTE: ignoring potential errors from input.Err()
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
```



#### a. if 循环

```go
if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
}
```

> `if` 语句条件两边也不加括号，但是主体部分需要加。`if` 语句的 `else` 部分是可选的，在 `if` 的条件为 `false` 时执行。

#### b. map

```go
counts := make(map[string]int)
//key : string
//value : int
```

> **map** 存储了键/值（key/value）的集合，对集合元素，提供**常数时间**的存、取或测试操作。(hash)
>
> **键可以是任意类型**，只要其值能用 `==` 运算符比较，最常见的例子是字符串；
>
> **值则可以是任意类型。**
>
> 这个例子中的键是字符串，值是整数。
>
> **内置函数 `make` 创建空 `map`，**此外，它还有别的作用。
>
> 从功能和实现上说，`Go` 的 `map` 类似于 `Java` 语言中的 `HashMap`，Python 语言中的 `dict`，`Lua` 语言中的 `table`，通常使用 `hash` 实现。
>
> `map` 中不含某个键时不用担心，首次读到新行时，等号右边的表达式 `counts[line]` 的值将被计算为其类型的零值，对于 `int` 即 `0`。

```go
input := bufio.NewScanner(os.Stdin)
counts[input.Text()]++ 

//等价于
line := input.Text()
counts[line] = counts[line] + 1
```

> 每次 `dup` 读取一行输入，该行被当做键存入 `map`，其对应的值递增。
>
> `intput.Text()` 是输入的值作为key

#### c. os.Args & os.Stdin

> `os.Args` 和 `os.Stdin` 都是 Go 语言标准库 `os` 包中的功能，但它们解决的问题和使用场景是不同的。
>
> **`os.Args`：**
>
> - **作用：** 用于获取**命令行参数**。
>
> **`os.Stdin`：**
>
> - **作用：** 用于**从标准输入读取数据**。

> 1. **`os.Args`：**
>
>    - **作用：** 用于获取命令行参数。
>
>    - **使用场景：** 当你运行一个可执行程序并在命令行中传递参数时，`os.Args` 用于访问这些参数。`os.Args` 是一个字符串切片，其中包含程序名称和从命令行传递的参数。
>
>    - **示例：**
>
>      ```go
>      package main
>      
>      import (
>          "fmt"
>          "os"
>      )
>      
>      func main() {
>          // 获取命令行参数
>          args := os.Args
>      
>          // 输出所有命令行参数
>          fmt.Println("Command-line arguments:", args)
>      }
>      ```
>
>    - **运行示例：**
>
>      ```bash
>      go run main.go arg1 arg2 arg3
>      # 输出: Command-line arguments: [main arg1 arg2 arg3]
>      ```
>
> 2. **`os.Stdin`：**
>
>    - **作用：** 用于从标准输入读取数据。
>
>    - **使用场景：** 当你希望从终端（命令行）接收用户输入时，`os.Stdin` 用于创建一个标准输入的 `*os.File` 对象，然后你可以使用其他包如 `bufio` 来处理输入。
>
>    - **示例：**
>
>      ```go
>      package main
>      
>      import (
>          "bufio"
>          "fmt"
>          "os"
>      )
>      
>      func main() {
>          // 创建标准输入的扫描器
>          input := bufio.NewScanner(os.Stdin)
>      
>          fmt.Print("Enter something: ")
>      
>          // 从标准输入读取一行文本
>          input.Scan()
>      
>          // 获取输入的文本
>          text := input.Text()
>      
>          // 输出输入的文本
>          fmt.Println("You entered:", text)
>      }
>      ```
>
>    - **运行示例：**
>
>      ```bash
>      go run main.go
>      # 输入: Hello, World!
>      # 输出: You entered: Hello, World!
>      ```

#### d. bufio 包 : Scanner

> 继续来看 `bufio` 包，它使处理输入和输出方便又高效。
>
> `Scanner` 类型是该包最有用的特性之一，它读取输入并将其拆成行或单词；通常是处理**行形式**的输入最简单的方法。

>  程序使用短变量声明创建 `bufio.Scanner` 类型的变量 `input`。

```go
input := bufio.NewScanner(os.Stdin)
```

> - 该变量从程序的标准输入中读取内容。
> - 每次调用 `input.Scan()`，即读入下一行，**并移除行末的换行符**；
>
> - 读取的内容可以调用 `input.Text()` 得到。
>
> - `Scan` 函数在读到一行时返回 `true`，不再有输入时返回 `false`。

#### f. fmt.Printf

> 类似于 C 或其它语言里的 `printf` 函数，`fmt.Printf` 函数对一些表达式产生格式化输出。
>
> 该函数的首个参数是个**格式字符串**，指定后续参数被如何格式化。
>
> 各个参数的格式取决于“转换字符”（conversion character），形式为百分号后跟一个字母。
>
> - 举个例子，`%d` 表示以十进制形式打印一个整型操作数，而 `%s` 则表示把字符串型操作数的值展开。
>
> `Printf` 有一大堆这种转换，Go程序员称之为*动词（verb）*。下面的表格虽然远不是完整的规范，但展示了可用的很多特性：

```go
%d          十进制整数
%x, %o, %b  十六进制，八进制，二进制整数。
%f, %g, %e  浮点数： 3.141593 3.141592653589793 3.141593e+00
%t          布尔：true或false
%c          字符（rune） (Unicode码点)
%s          字符串
%q          带双引号的字符串"abc"或带单引号的字符'c'
%v          变量的自然形式（natural format）
%T          变量的类型
%%          字面上的百分号标志（无操作数）
```

> `dup1` 的格式字符串中还含有制表符`\t`和换行符`\n`。
>
> 字符串字面上可能含有这些代表不可见字符的**转义字符（escape sequences）**。默认情况下，`Printf` 不会换行。
>
> 按照惯例，以字母 `f` 结尾的格式化函数，如 `log.Printf` 和 `fmt.Errorf`，都采用 `fmt.Printf` 的格式化准则。
>
> 而以 `ln` 结尾的格式化函数，则遵循 `Println` 的方式，以跟 `%v` 差不多的方式格式化参数，并在最后添加一个换行符。（译注：后缀 `f` 指 `format`，`ln` 指 `line`。）

#### g. os.open

很多程序要么从标准输入中读取数据，如上面的例子所示，要么从一系列具名文件中读取数据。`dup` 程序的下个版本读取标准输入或是使用 `os.Open` 打开各个具名文件，并操作它们。