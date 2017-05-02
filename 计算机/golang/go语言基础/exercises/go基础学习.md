# 学习笔记
## 基本结构和数据类型
### 符号
- **文件名**   小写，使用`_`分割，以`.go`为扩展名；源文件没有大小限制；
- **标识符**   以UTF8字符或`_`开头，后跟0或多个字符或Unicode数字;关键字不能做标识符，汉字被看做小写
- **关键字**   25个
```
break       default         func        interface       select
case        defer           go          map             struct
chan        else            goto        package         switch
const       fallthrough     if          range           type
continue    for             import      return          var
```

- **预定义标识符**   由基本类型和内置函数组成,共36个
    - 基本类型：nil、iota；bool、true、false；byte；rune；int、int8、int16、int32、int64、uint、uint8、uint16、uint32、uint64；float32、float64；string；uintptr；complex、complex64、complex128、imag、real； 
    - 内置函数：append、cap、close、copy、len、make、new、panic、recover、print、println  

- **符号**
    + 分割符  `()`、`[]`、`{}`
    + 标点符号 `.` 、 `,`、 `;`、 `:`、 `...`.  `;`这个不用于语句结束，但是当多个语句子在一行时需要用它隔开

- **运算符**
    ```
    优先级 	运算符
     7         ^ !
     6         * / % << >> & &^
     5         + - | ^
     4         == != < <= >= >
     3         <-
     2         &&
     1         ||
    ```
    注：
    1. `&^ ` 位清楚操作 0xF &^ 0 = 0xE
    2. `^` 有异或和按位补足作用；按位补足 :question:
    ### go程序的基本组成

1. **包** 程序的机构化组织,类似命名空间
    - `package 包名`  源文件所属包，非注释首行； `package main` 只能有一个
    - `import "包名"` 或者`import ("包名1" \r 包名2)`  `import ("fmt"; "os")`  导入包,包路径:`/` 绝对路径、 `./`当前路径、 未指定 全局路径；如果导入未使用则错误 `imported and not used: os`
    - 可见性  大写字母开头的都是对外可见，小写字母开头的都只能包内可见

2. **函数** 功能块,必现且只能有一个main函数

3. **注释** `//` 行注释,`/*  ...  */`块注释
    - 注释不被编译但可以被godoc使用
    - package 语句之前的注释默认当做包说明
    - 全局作用域的类型、函数、变量、常量都应该有个合理的注释说明，而且注释应当以表示开头，可以被用于生产说明文档

4. **类型** 数据的抽象
    - 类型别名  `type 新类型名  类型名`  
    - 类型转换  `var i = int(1.5)`  golang只能强制类型转换  

5. **常量** 在整个程序运行过程中都不会改变的存储单元
    - 字面常量 如：数字 9527 3.1415、 字母 'a' 'd'、字符串 "hello world"
    - 符号常量 使用`const`声明符号常量
```golang
    const PI = 3.1415926
    const PI float32 = 3.14
    const (
        周一 = iota   //从零开始 自动增长；iota可以改变值 iota+1则从1开始；iota*2 + 1 则是1,3,5,7
        周二
        周三
        周日 = 7
        周四          //如果没有在知道为 iota，之后所有的常量都将是7；如果再次指定为iota，周四为4，从当前第几个开始，之后的继续递增
        周五
    )
```

6. **变量** 值可以被改变的存储单元
```golang
    var num int  //定义
    num = 10    //赋值
    var num int = 10 //可以推导类型 一般不适用
    var num = 10
    num := 10    //定义并赋初始值的简化形式 
```
    注：go中有值类型和引用类型两类数据，零值和默认值不同

7. **语句** 表达计算过程

### 基本类型
1. 布尔 bool true false
2. 整型 int int8 int16 int32 int64 uint uint8 uint16 uint32 uint64
3. 浮点型 float32 float64
4. 复数 complex64 complex128 (imag、real)
5. 字符 rune  32位的一个unicode码
6. 字符串 string
7. 指针 uintptr

------------------------------
## 控制结构
1. 顺序 按顺序执行的语句
2. 选择 根据条件选择不同执行语句
    - if 语句
        1. if 条件 {...}
        2. if 初始化语句;条件{...}
        3. if 条件 {...} else {...}
        4. if 条件 {...} else if {...}
    - switch 语句 case执行后默认退出，需要继续向下执行需要添加`fallthrough`
    ```golang
            switch [表达式] { //表达式可以省略
                case 条件:    //确定的值，为表达式类型
                    执行语句1
                case 条件2,条件3：//可以同时从在多个条件
                    执行语句2
                case i > 10      //可以试逻辑表达式
                    执行语句3
                default:        //默认语句不是必须有的
                    执行语句4
     ```
    - select 语句 主要用于从多通道中选择一个科执行的

3. 循环 只有for循环
    - for {...}       //死循环
    - for 条件 {...}  //条件成立循环
    - for 初始化;条件;结束执行 {...}  //三个语句可省略但分号不能少
    - for k,v  := range 集合 {...}     //用于 数组 切片 字典 字符串等多元元素遍历,k为索引，v为元素拷贝，注：字符串索引不一定连续

4. 跳转
    - 标签  代码中一个执行位置,用于跳转；`lable:`
    - break
        - 结束当前循环
        - break lable 结束到标签
    - contiune
        - 结束本次循环继续下次循环
    - contiune 结束本次循环到标签
    - goto 程序跳转到指定标签


------------------------------
## 数组 （array）
> 大小固定 相同类型的一组变量 较少使用
### 概念
数组：数组是一个由 **固定长度** 的 **特定类型** 元素组成的序列
### 变量定义

```golang
   var arr [3]int = [3]int {1, 2, 3}  //错误 var arr [3]int = {1, 2, 3} var arr [...]int = [3]int{1, 2, 3}
   // 正确 var arr [3]int = [...]int{1, 2, 3}  arr := [3]int{1, 2, 3}  var arr = [3]int{1, 2, 3}  arr:= [3]int{} 
   // var  arr [3]int = [3]int{1:5}  指定特定的元素赋值
   // 注意 1，数组的长度是数组的基本属性， [3]int 和 [4]int 是不同的类型；
   //      2, ... 在数组元素中用于自动推导个数，不能用于变量类型说明中，如：  var  arr [...]int
   //      3，var arr []int = [3]int {1, 2, 3} var arr []int = []int {1, 2, 3} 中arr 不是数组而是切片了
```

## 操作
1. 遍历
使用 `for i:=0; i<len(arr); i++{}` 遍历元素，len 内置函数，求的数组的长度
使用 `for ... range arr {}` 自动遍历所有元素，`i, v := range arr` i为下标，v是当前元素值的副本，不是当前元素，不能通过修改v修改元素实际的值。 range是关键字不是函数，不需要括号。

2. 比较
数组可以使用 `==` 比较，只有数组类型和数值全都相同时，才为真。   `!=` 同理只要类型或值不相等则为真

3. 求长度  
len() 可以求的数组的长度
cap() 可以得到数组的容量  数组的长度就是它的最大容量

### 注意项
1、做入参传入函数时是值传递，复制整个数组到函数内部。
   但也可以传入数组指针，从而在函数内直接修改数组内容

```golang
   func zero(ptr *[32]byte) { //不是数组元素的基本类型指针，而是一个数组类型的指针
    for i := range ptr {     //数组的长度是固定的
        ptr[i] = 0  // 直接修改原数组    
     }
   }
```
2、数组的大小是固定的不能删除或添加元素

3、数组可以整体重新赋值

```golang
    var arr [3]int = [3]int {1, 2, 3}
    arr = [3]int{}   // 将变为0，但是不能赋值不同长度类型的，如 [4]int {3,4}
```

----------------

## 切片 （slice）
> 一个数组的片段引用。

### 概念
Slice（切片）代表变长的序列，元素为同一类型。切片类型记作[]T，T代表slice中元素的类型。本质是对数组一段的一个引用，它由指针、长度和容量组成，指针指向切片第一个元素，长度是切片元素的多少，容量是切片第一个元素到底层数组最后的元素之间的大小。多个切片可以引用同一底层数组，且范围可以重叠

### 切片的定义

```golang
    var slice []int = []int{}
    fmt.Printf("%T, %d, %d, %v\n", slice, len(slice), cap(slice), slice)  //输出 []int, 0, 0, []

    var slice []int = []int{1, 2, 3, 4, 5}   //错误 var slice []int = [10]int{1, 2, 3, 4, 5} 类型不同
    fmt.Printf("%T, %d, %d, %v\n", slice, len(slice), cap(slice), slice) //输出 []int, 5, 5, [1 2 3 4 5]

    var array = new([5]int)
    var slice []int = array[:]
    fmt.Printf("%T, %d, %d, %v\n", slice, len(slice), cap(slice), slice)

    var array = new([10]int)      //新建一个数组，new出来的数组是*[10]int类型
    var slice []int = array[:6]   //slice 应用数组的一部分
    slice[4] = 15       //改变了原数组元素的值   slice[7]错误 超出范围
    fmt.Printf("%T, %d, %d, %v, %v\n", slice, len(slice), cap(slice), slice, *array)
    //输出[]int, 6, 10, [0 0 0 0 15 0], [0 0 0 0 15 0 0 0 0 0]

    //还有以下形式
    var slice = []int{1, 2, 3, 4, 5}
    slice := []int{1, 2, 3, 4, 5}
    var slice = make([]int, 5)   // 没有make([]int)
    var slice = make([]int, 5, 10)
    slice := make([]int, 5)
    slice := make([]int, 5, 10)
```

### 操作
1. 遍历 切片的遍历和数组相同，for  range 遍历，或者使用标遍历
2. len() 和 cap() 获取长度和容量
3. make(Type,Len) 和 make(Type, Len, Cap) 创建切片
4. append(s S, x ...T) S 追加元素
```golang
    array := [5]int{}
    slice := array[:3]
    fmt.Println(slice) //输出 [0, 0, 0]
    fmt.Println(array) //输出 [0, 0, 0, 0, 0]
    slice1 := append(slice, 4, 5)
    fmt.Println(slice1) //输出 [0, 0, 0, 4, 5] //追加到slice后，覆盖原数组元素内容
    fmt.Println(array)  //输出 [0, 0, 0, 4, 5] //没有超过容量时，覆盖原数组
    slice2 := append(slice1, 6, 7)
    fmt.Println(slice2) //输出 [0, 0, 0, 4, 5, 6, 7]  超过原数组容量后，则会自动重新创建一个更大的数组存放元素
    fmt.Println(array) //输出 [0, 0, 0, 4, 5]   原数组不受影响
    slice = slice[0:5] // 未超过底层数组容量，可重新切片
    slice = slice[0:7] //超容量，运行时错误 panic: runtime error: slice bounds out of range
```
5. copy(dst, src []T) int 和 copy(dst []byte, src string) int  切片复制
```
    array := [5]int{1, 2, 3, 4, 5}
    var slice = []int{5, 4, 3}
    n := copy(slice, array[0:5])  // n=3 slice = [1, 2, 3] 不能给clice扩容
```

### 注意项

1、数组可以直接专为切换传给函数

```golang
    array := [3]int{1, 2, 3}
    reserse(array[:])  //翻转  原数组可以被改变
```
2、 slice 不能用 `==` 判断是否相等，因为slice是引用类型。
唯一能和nil比较，一个nil值的slice的长度和容量都是0。
测试一个slice是否是空的，应使用len(s) == 0来判断，而不应该用s == nil来判断，因为 []int{} != nil。
字节类型可以通过标准库中 `bytes.Equal` 比较，其他类型需要自己写函数比较，如：
```golang
func equal(x, y []string) bool {
    if len(x) != len(y) {
        return false
    }
    for i := range x {
        if x[i] != y[i] {
            return false
        }
    }
    return true
}
```
3、slice 在追加元素时如果超过了底层数组容量，则会重新创建一个更大容量的副本数组。

## make 和 new的区别
new(T) \*T 分配一个新的空间，返回的是空间的地址,空间内容为类型对应的初始化值
```golang
    var n = new(int)
    fmt.Printf("%T %v %v\n", n, n, *n)  //输出 *int 0xc042008230 0
```
make() 只用于 slice、map、channel类型，返回的是值。

```
Call             Type T     Result

make(T, n)       slice      slice of type T with length n and capacity n
make(T, n, m)    slice      slice of type T with length n and capacity m

make(T)          map        map of type T
make(T, n)       map        map of type T with initial space for n elements

make(T)          channel    unbuffered channel of type T
make(T, n)       channel    buffered channel of type T, buffer size n
```
建立值有不同的初始化条件，类似于函数的重载吧，因为slice、map、channel是内置类型，所以需要这样一个内置函数来创建对应的值。


----------------
## 字典(map)
> key/value对的集合

### 概念
它是一个无序的key/value对的集合,所有的key都不同，是哈希表的引用。可在常数时间复杂内更新key对应的value。

### 变量定义
map的类型为 `map[keyType]ValueType`,key 的类型必现是可比较的。map没有容量的概念，不能使用cap(map)
```golang
    var week map[uint8]string    //声明一个week map，week默认为空(nil),week == nil, len(week)==0
    var week map[uint8]string = map[uint8]string{}  //定义week一个空内容的map，week != nil, len(week)==0,可添加元素week[0] = "Sunday"
    var week = map[uint8]string{}  //同上类型自动推导
    var week map[uint8]string = map[uint8]string{0: "Sunday", 1: "Monday"}  //初始化内容

    //map 定义大多使用make
    week := make(map[uint8]string)  //为空但不nil
    //week := make(map[uint8]string){0: "Sunday", 1: "Monday"} //错误 make是一个函数返回的是一个map
    week := map[uint8]string{0: "Sunday", 1: "Monday"}
    week := make(map[uint8]string, 10)
    //初始化一个预估map容量，map增大时会重新分配内存；
    //给预估容量是让map初次分配预估的容量，元素个数不超过容量时，是不会重新分配内存的，重新分配内存会增加额外的系统消耗。
    //所以在可以预估map时，可以指定预估容量，从而提高效率；
    //指定预估容量，不会改变map的大小，它还是一个空map；
    //空map和 nil map是不同，空map是没有内容，但已经引用了一个哈希表，而nil map还没有引用一个哈希表，就是还不存在一个具体的map
```

### 操作
1. 遍历 使用for range循环遍历所有元素
```golang
    for name, age := range ages {
        fmt.Printf("%s\t%d\n", name, age)
    }
```
2. 添加修改元素 通过key和value直接添加或修改元素 `week[key] = value`,如果key存在则是修改元素，如果key不存在则添加元素
3. 删除元素  `delete(week,key)` 从map week中删除key对应的元素
4. 判断元素是否存在
如果key存在则返回对应的value，如果key不存在则返回value类型的默认值,但是默认也是有效值则不能直接判断，需要通过返回的错误内容判断，如：
```golang
    age, ok := ages["bob"]
    if !ok { /* "bob" is not a key in this map; age == 0. */ }
    //与上面两行等同
    if age, ok := ages["bob"]; !ok { /* ... */ }
```

### 注意项
1. 不宜使用浮点型数据做key，虽然浮点型也支持 `==` 相等运算比较，但是有可能出现NaN和任何浮点数都不相等。
2. map中的元素不是一个变量，不能对map的元素进行取址操作，这是由于随着map内容的增加，可能重新分配内存空间。
3. map的迭代遍历是随机顺序的，和初始化时的顺序没有关系。每次都故意的使用随机顺序的遍历，这就强制要求程序不能依赖具体的哈希函数实现。
如果要顺序遍历，需要对key按顺序遍历value,如：
```golang
    var names []string
    for name := range ages {
        names = append(names, name)
    }
    sort.Strings(names)
    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
```
4. 不能向nil map中存入元素，将导致一个panic，如：
```golang
    var week map[uint8]string
    week[0] = "Sunday"   // panic: assignment to entry in nil map`
```
5. map不能通过 `==` 直接判断是否相等，只能和nil比较；如果要判断两个map是否相同，需要通过循环遍历去判断，如：
```golang
func equal(x, y map[string]int) bool {
    if len(x) != len(y) {
        return false
    }
    for k, xv := range x {
        if yv, ok := y[k]; !ok || yv != xv {
            return false
        }
    }
    return true
}
```
6. 可以通过辅助函数将不能直接比较的类型数据转为可以直接比较的值，然后用其做key。

----------------
## 函数（function）
> 功能模块化，提高代码复用

### 概念
函数：有输入和输出的代码功能块。方便代码复用和结构化。

### go函数特点
- 不持嵌套 (nested)
- 不持重载 (overload)
- 不持默认参数 (default parameter)
- 无需声明原型
- 支持不定变参
- 支持多返回值
- 支持命名返回参数
- 支持匿名函数和闭包
- 支持函数递归

### 函数声明
函数声明包含：关机字 `func` 、【函数名】、【函数输入参数列表】、【函数返回值列表】、花括号、【函数体】，带括号表示可选
```golang
func 函数名(输入参数 string) (返回值名 bool) { //返回值可以命名，如果只有一个返回类型可以省略小括号，
            //如果没有函数名则是一个匿名函数，一般直接调用（在函数体后直接加小括号和参数去调用）
    if 输入参数 != "" {
        fmt.Println(输入参数)  //输出参数的内容
        返回值名 = true  //可对返回变量赋值
    } else {
        return false    //  直接返回的值将赋值到返回变量
    }
    return   //有命名返回变量，可以不带值 也可以带对应列表参数的多个值；没有命名返回值，则需要和函数返回值列表相对应
}
func main() {
    fmt.Println(函数名("hello")) //输出 hello 和 true
    fmt.Println(函数名(""))      //输出false
}

//常见形式
func add(x1, x2 int) int {}
func add(x1, x2 int) (sum int) {}
func add(x1, _ int) (sum int) {}   //强调参数未被使用，多用于占位。
func divide(x1, x2 int)(int, error)()
func (x1, x2 int)int {} (2, 4) // 匿名函数直接执行
func Sin(x float64) float //implemented in assembly language，表示不是GO实现的函数，是在说明
```

### 函数的特殊用法
1. 作为值传递
函数也有值和类型，可以赋值给其他变量、可以作为函数参数、可以被函数返回
```golang
func square(n int) int { return n * n }
fn := square  //定义并初始化fn变量，fn推导出的类型为 func(int) int , 但square是不能被赋值的
var fn func(int) int  //声明 fn 默认零值是nil;调用fn(3)则发生panic
fn = square   //给fn赋值
callback(fn)  // 作为值传给函数调用，一般用作回调函数
```
2. 匿名函数
没有名字的函数叫匿名函数(anonymous function),函数值字面量是一种表达式，它的值被成为匿名函数；
重要的是通过这种方式定义的函数可以访问完整的词法环境（lexical environment）；
Go使用闭包（closures）技术实现函数值，Go程序员也把函数值叫做闭包；
函数值是引用类型，可以引用父函数的上下文变量；
匿名函数也可以递归，先赋值给一个变量，再在自己内部调用函数变量，从而形成递归调用
```golang
// squares返回一个匿名函数。
// 该匿名函数每次被调用时都会返回下一个数的平方。
func squares() func() int {
    var x int
    return func() int {
        x++  //匿名函数像表达式一样，可以访问所在作用域，并且是引用访问
        return x * x
    }
}
func main() {
    f := squares()
    fmt.Println(f()) // "1"
    fmt.Println(f()) // "4" 对上次的x还存在引用
    fmt.Println(f()) // "9"
    fmt.Println(f()) // "16"
}
```

3. 工厂函数
一个返回值为另一个函数的函数可以被称之为工厂函数，这在创建一系列相似的函数的时候非常有用：书写一个工厂函数而不是针对每种情况都书写一个函数。
以下例子演示动态添加文件后缀
```
// 创建一个添加后缀的函数
func MakeAddSuffix(suffix string) func(string) string {   //工厂函数，可以创建相似的函数
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}
func main() {
    addBmp := MakeAddSuffix(".bmp")   //创建添加.bmp的函数
    addJpeg := MakeAddSuffix(".jpeg")  //创建添加.jpeg的函数
    fmt.Println(addBmp("filename"))   //输出filename.bmp
    fmt.Println(addJpeg("filename1")) //输出filename1.jpeg
}
```

4. 变参函数 在最后参数类型之前加上 `...`
```golang
func sum(vals...int) int {
    total := 0
    for _, val := range vals { //vals可以类似的看做切片
        total += val
    }
    return total
}
fmt.Println(sum())
fmt.Println(sum(1, 2, 3, 4)) // "10"
values := []int{1, 2, 3, 4}
fmt.Println(sum(values...)) // "10" 直接传入切片时，需要在后面加上省略号
func errorf(linenum int, format string, args ...interface{}){} //interfac{}表示函数的最后一个参数可以接收任意类型
```
5. defer 当前函数结束后调用
defer 关键字 延迟函数执行，在当前函数结束后才被执行，主要用于完成函数执行完成后的善后工作； 
多用于成对的操作，如：打开、关闭、连接、断开连接、加锁、释放锁； 
可以有多个，按声明反向顺序执行； 
在循环中使用时，需要注意是在函数执行完成后执行，而不是在循环结束后执行，可将循环体修改为函数去执行； 
还可以追踪函数离开 `defer untrace(s string) { fmt.Println("leaving:", s) }`  
6. 内置函数  
**close**  用于管道通信   
**len、cap** len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）   
**new、make**  new 和 make 均是用于分配内存：new 用于值类型和用户定义的类型，如自定义结构，make 用户内置引用类型（切片、map 和管道）。它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int) 。make(T) 返回类型 T 的初始化之后的值，因此它比new 进行更多的工作    
**copy、append**   用于复制和连接切片  
**panic、recover**  两者均用于错误处理机制  
**print、println** 底层打印函数，在部署环境中建议使用 fmt 包  
**complex、real imag**  用于创建和操作复数  
### 注意项
1. 函数参数是值传递，但是对于slice、map、channel、function等引用类型，可能会被间接修改实参
2. 捕获迭代变量 在使用闭包时，循环中的变量会被引用到匿名函数中，因此可能始终访问的是同一个变量,变量逃逸，分配在堆上
```golang
var rmdirs []func()
for _, d := range tempDirs() {
    dir := d // NOTE: necessary! 重新创建一个变量
    os.MkdirAll(dir, 0755) // creates parent directories too
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) //匿名中使用的是引用
    })
}
// ...do some work…
for _, rmdir := range rmdirs {
    rmdir() // clean up
}
//下面这个方法时错误的
var rmdirs []func()
for _, dir := range tempDirs() {
    os.MkdirAll(dir, 0755)
    rmdirs = append(rmdirs, func() {
        os.RemoveAll(dir) // NOTE: incorrect!  删除的是同一个目录
    })
}

//一般我们会引入一个与循环变量同名的局部变量，作为循环变量的副本，来解决这个问题
for _, dir := range tempDirs() {
    dir := dir // declares inner dir, initialized to outer dir
    // ...
}
```
----------------
## 结构体（struct）
> 组装各种基本类型成为一个新的类型

### 概念
结构体是一种聚合的数据类型，结构体是复合类型（composite types），是由零个或多个任意类型的值聚合成的实体。每个值称为结构体的成员或字段。Go语言中结构体还可以有方法。

### 结构体类型声明
- 结构体类型声明：创建一种新的类型
```golang
// Person 人类
type Person struct {  //必须要有type和结构体名，type 用于重新给类型起个新名字 alias types
	Name    string  //必须要有字段名，如果不使用可以用_省略
	Age     uint8   //字段可以是任何类型，包括自身类型、函数、接口
	Sex     uint8   //字段名不能重复
	Stature uint8
    
}
//可以是匿名结构类型
type House struct {  //房子
    Bed struct{ //床  这是一个匿名结构类型
        Width  int //床宽
        Length int //床的长度
   }
}
//可以定义字段标签用于反射读取
type u1 struct { name string "username" }

//定义空类型节省空间
var null struct{}
set := make(map[string]struct{})
set["a"] = null  //只关注key

//匿名字段 匿名字段不过是一种语法糖 被匿名嵌入的可以是任何类型
type User struct { name string }
type Manager struct {
    User 
    title string
}
m := Manager{
    User: User{"Tom"}, // 匿名字段的显式字段名，和类型名相同。
    title: "Administrator",
}
m.name //可以像普通字段那样访问匿名字段成员,编译器从外向内逐级查找所有层次的匿名字段，直到发现目标或出错
//有同名字段时，外层同名字段会遮蔽嵌入字段成员，同层则需要使用显示字段名
//不能同时嵌入某一类型和其指针类型，因为它们名字相同
//匿名主要是用来扩展类型，有类似特例的关系，增加一个新特性，如人->官，官比人多一个职务，是人的一个特种。
```

- 声明结构体变量，和其他普通类型一样
```golang
var jack Person
var jack Person = Person{"jack", 18, Male, 170} //初始化必须包含所有字段
var jack  = Person{"jack", 18, Male, 170}
var jack  = Person{Stature:170, Sex:Male, Age:18, Name:"jack"} //可以用标签指定对应字段的值
var jack Person = Person{"jack", 18} //错误 必须指定全部字段
jack := Person{} //可以全空
jack := new(Person)  // jack这时是一个指针
```

### 操作
1. 访问字段  通过`.`访问字段
jack.Age = 20 //如果是指针也可以通过点直接访问
2. 添加方法  `func (p Person) Say(str string) {fmt.Println(p.Name + ":" + str)}`
- 方法名不能和原有字段重复
- 如果需要修改自身数据需要传入指针类型对象，如`func (p *Person) AddAge() {p.Age++}`
- 调用方法  `jack.Say("你好")` //输出jack:你好
3. 支持 "=="、 "!=" 相等操作符
4. 类型的 String() 方法和格式化描述符
    对类型一种可阅读性和打印性的输出，可以直接在printf 中使用%v 输出。
```golang
    func (p *Person) String() string {
	    return "[" + p.Name + ":" + strconv.Itoa(p.Age) + "]"  //转换为字符串
    }
 ```
### 注意项
1. 匿名字段类型中的相同字段会被覆盖，同级的需要明确指定字段
2. 不能有相同字段，包括方法，方法也不能和字段相同

### :question:疑问？
1. 结构体中的补位字段有什么作用？

----------------
## 接口（interface）
> 一组功能的合约

### 概念
接口类型是对其它类型行为的抽象和概括。接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

作用：对不同类型相同操作的一种提炼，如，`进食` 是猪、狗、牛、羊、马这些不同类型的动物，但是他们都有`进食`这个方法。所有可以将进食抽象为一个接口，这个接口就可以处理所有这些类型的对象。每一种动物的进食由自己实现，接口只是个抽象。

接口值：包含接口所对应的具体类型和对应类型的具体值，接口零值指类型和具体值都为nil

### go接口特点
- 接口没有字段
- 接口只有方法签名没有实现
- 接口可以嵌套其他接口
- 类型可以实现多个接口
- 类型只要实现了接口的所有方法，则认为类型实现了该接口，不需要显示的说明。
- Go语言中的接口都很简短，通常它们会包含 0 个、最多 3 个方法。
- Go 语言中接口可以有值，一个接口类型的变量或一个接口值是一个多字（multiword）数据结构
- 接口可以引用类型的变量，即将变量赋值给接口变量；接口变量里包含了接收者实例的值和指向对应方法表的指针

### 接口声明 变量定义

```golang
type 图形 interface {    //声明接口 '图形'
	面积() float32       //方法标签 不需要 'func' 关键字
	周长() float32       //要有名字 参数 返回值  方法的顺序没有关系
}

//类型必现实现接口的所有方法，才能算类型实现了接口
type 圆形 struct { 半径 float32 }  //声明圆形
func (圆 圆形) 面积() float32 {	return math.Pi * 圆.半径 * 圆.半径 }  //给圆形添加‘面积’方法
func (圆 圆形) 周长() float32 {	return 2 * math.Pi * 圆.半径  }      //给圆形添加‘周长’方法

type 矩形 struct {  //声明矩形
	长 float32
	宽 float32
}
func (长方形 矩形) 面积() float32 {  return 长方形.长 * 长方形.宽  }        //给矩形添加‘面积’方法
func (长方形 矩形) 周长() float32 {  return 2 * (长方形.长 + 长方形.宽)  }  //给矩形添加‘面积’方法

func show(t 图形) {  fmt.Printf("面积:%v  周长:%v\n", t.面积(), t.周长())  } //输出图形的面积和周长

func main() {
	var tx 图形 = 圆形{2} //定义图形接口变量tx，并赋值一个圆形的实例
	show(tx)          //输出   面积:12.566371  周长:12.566371                                                                             
	tx = 矩形{2, 3}  // 赋值一个矩形的实例
	show(tx)      //输出   面积:6  周长:10      
}
```
- 类型声明
type IF_Name interface { 方法标签列表 ... }
- 变量定义
var interf IF_Name //可以用实现接口类型的变量，给接口变量赋值
var interf IF_Name = IF_Name{} //错误  不支持空接口

### 操作
- 类型断言
```golang
if v, ok := tx.(圆形); ok {  // 判断tx接口是不指向圆形的实体，OK
    Process(v)    //v指向具体的值
    return
}
//  类型判断 type-switch
switch t := areaIntf.(type) {
case *Square:
	fmt.Printf("Type Square %T with value %v\n", t, t)
case *Circle:
	fmt.Printf("Type Circle %T with value %v\n", t, t)
case nil:
	fmt.Printf("nil value: nothing to check?\n")
default:
	fmt.Printf("Unexpected type %T\n", t)
}
```
- 测试接口实现
```golang
type Stringer interface {
    String() string
}
if sv, ok := v.(Stringer); ok { //v是一个值
    fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
}
```
### 注意项
1. Go 语言规范定义了接口方法集的调用规则：
    - 类型 \*T 的可调用方法集包含接受者为 \*T 或 T 的所有方法集
    - 类型 T 的可调用方法集包含接受者为 T 的所有方法
    - 类型 T 的可调用方法集不包含接受者为 \*T 的方法
2. 空接口
    - 任何接口都实现了空接口
    - 可以给空接口赋值任何类型的值
    - 主要用于通用类型构建
3. 反射  利用接口获得类型和值
func TypeOf(i interface{}) Type
func ValueOf(i interface{}) Value

----------------
## go协程（Goroutines）
> 并行执行的协程

### 概念
- 进程：进程是一个运行在自己内存地址空间里的独立执行体，进程由一个或多个操作系统线程组成
- 线程：具体的执行单元，共享进程内存空间地址。
- 协程：线程中更小的执行单元，自带CPU上下文，切换比线程更轻量，协程和线程没有一一对应关系。
- 并行：同一时刻，多个执行体同时在执行；只有在多核多处理器上才存在。
- 并发：一段时间段内，有多个执行体在执行。

golang在语言级别通过goroutines支持并发

### 启动一个go协程
```golang
go functionName() //这将创建一个go协程，go关键字后面跟一个函数调用
```

### 注意项
1. main中协程启动后继续执行，如果执行到main结束，则程序执行结束，所有需要等待协助执行完成。
----------------
## 通道（channels）
> 相当于一个类型化的消息队列 ，主要用于go协程间的同步通信。

### 概念
通道是golang的一个特殊类型，可以在协程间传递类型化的数据，避开所有内存共享导致的坑；通道保证了同步性；同一时刻只有一个协程能够访问，所有不会出现竞争。
通道可以传递任何类型的数据，包括空接口和通道
### 变量定义

```golang
    var ch  chan int   //ch是一个传递int类型数据的，它并没有初始化，是一个零值nil
    ch = make(chan int)  //通道是应用类型，使用make分配内存，创建其实体
    ch := make(chan int, 4)  //通道可以带缓存  
    var send_only chan<- int 		// channel can only send data
    var recv_only <-chan int		// channel can onley receive data
```
### 通道特性
1. 数据收发   操作符`<-`对通道进行收发,明显的表明了数据的方向
    - **发送**  ch <- 100   将100放入通道ch
    - **接受**  v = <- ch   将通道ch中的值取出来赋值给变量v
2. 关闭通道 close()
只有在当需要告诉接收者不会再提供新的值的时候，才需要关闭通道。只有发送者需要关闭通道，接收者永远不会需要。对已经关闭的通道发送数据或再次关闭已经关闭的通道都会导致panic。 `if v, ok := <-ch; ok {  process(v) }` 判断是否可以接收阻塞的通道或者关闭的通道；`for-range`可以判断通道是否关闭。
3. 获取通道容量  cap()，不是剩余可存放的空间
4. 通道阻塞
默认通道是不带缓存的，当发送时，通道中有数据则发送过程会阻塞，接收数据时，如果通道为空则接收过程阻塞；带缓存的进程在缓存未满时是非阻塞的。
5. 使用通道输出结果 为了知道计算何时完成，可以通过信道回报。 一直阻塞的有结果。
6. 通道可以在声明是指定方向，只发生或只接收
7. for...range 读取通道内容  for v range ch { fmt.Println(v) }
8. 实现信号量
9. select切换协程 从中选择一个可用的
```golang
select { //从中选择一个可用的执行，如果有多个则随机选择一个执行。都不可用则一直阻塞。break和return可以结束select
    case v := <-ch1:
        fmt.Printf("Received on channel 1: %d\n", v)
    case v := <-ch2:
        fmt.Printf("Received on channel 2: %d\n", v)
    default：//default永远执行
```
10. 超时和计时器（Ticker）
```golang
time.Tick(1e8) //time包中提供一个间隔通道发送数据
select { //在select中接受time中的间隔数据
    case <-ch:
        // a read from ch has occured
    case <-timeout:
        // the read from ch has timed out
        break
}
```
11. 协程恢复（recover）  协程发生panic时不会影响其他协程，只是自身被释放
```
func server(workChan <-chan *Work) {
    for work := range workChan {
        go safelyDo(work)   // start the goroutine for that work
    }
}

func safelyDo(work *Work) {
    defer func {
        if err := recover(); err != nil {
            log.Printf("Work failed with %s in %v", err, work)
        }
    }()
    do(work)
}
```

### 注意项 :bangbang:
1. 所有协程都在等待时将发送panic，即死锁，可以再运行时检测到死锁。

### 疑问 :question:
1. 缓存为1 和不带缓存通道的区别？
    - 不带缓存：即缓存为0，发送是阻塞的，只有发送后被其他goroutines接收后才能继续执行。主要用于同步。
    - 带缓存：发送不阻塞，发送后立即执行之后的语句。
