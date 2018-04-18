# 变量定义

1. `var` 但不指定初值
```go
var a
```

2. `var` 指定初值
```go
var a = 1
```

3. `:=`
```go
a := 1
```
(上面这种方式不能再函数外部)


4. `var ()`


```go
var (
  a = 1
  b = 2
)
```


# 内建变量类型

* bool
* string
* (u)int, (u)int8, (u)int16, (u)int32, (u)int64, uintptr
* byte
* rune 字符型，4 个字节，等于是 int32 , 至于为什么不是 unit32，可以参考(https://stackoverflow.com/questions/24714665/why-is-rune-in-golang-an-alias-for-int32-and-not-uint32)
* float32, float64, complex64, complex128

# 强制类型转换
不能隐式转换

# 常量与枚举

* 枚举

```go
const (
  a = 0
  b = 1
)
```

简化写法：

```go
const (
  a = iota
  b
)
```

`iota` 是行计数器，每个 const 都从 0 开始。也可以跳过一个：

```go
const (
  a = iota
  _
  javascript
)
```

下面来一个更牛逼的例子：

```go
const (
  b = 1 << (10 * iota)
  kb
  mb
  gb
  tb
  pb
)
```

# 条件语句
`if` 条件不需要括号。

紧凑写法：

```go
if content, err := ioutil.ReadFile(filename); err != nil {
  ...
```

# switch
* 不需要break

```go
switch {
case score < 60:
  g = "F"
case score < 80:
  g = "C"
default:
  panic(fmt.Sprintf("error"))
}
```

or 

```go
switch score {
case 60:
  g = "F"
case 80:
  g = "C"
default:
  panic(fmt.Sprintf("error"))
}
```

# 循环
* No while, use for to replace

Only stop condition:

```go
for scanner.Scan() {
  
}
```

Nothing:

```go
for {

}
```

# Function
* multi-return
* annonymous function
* variable arguments

```go
func sum(numbers ...int) int {
  
}
```

* augument passing

value ? reference ?

go only have **value passing**.



# Pointer
* cannot computer


# array, slick, container
## array
```go
var arr [5]int
arr := [3]int{1, 2, 3}
arr := [...]int{2, 3, 4, 5}
```

traverse array:

```go
for i, v := range array {

}
```

Array is value type. When used as func's argument, must specify the length. And the change made to array in func will not change the array itself.


## slice

```go
arr := [...]int{0,1,2,3,4,5,6}
s := arr[2:6]
```

slice's capacity:

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6] // length is 4, the capacity is 6
// s1 := arr[2:6:6] // length is 4, the capacity is 4
s2 := s1[3:5] // length is 2, capacity is 3
fmt.Println(s1, s2)
```

append item to slice:

```go
arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7}
s1 := arr[2:6]
s2 := s1[3:5]

// s2 = [5, 6, *7]
s3 := append(s2, 10) // capacity of s2 is 3 and length is 2, so can append item
fmt.Println(s3, arr)
s4 := append(s3, 11) // capacity and length of s3 is 3, can not append item, to append must create a new array, the old array may be collected
fmt.Println(s4, cap(s4), arr) // now the capacity of s4 is 6 ( 2 times of s3)
s5 := append(s4, 12)
fmt.Println(s5, cap(s5), arr)
```


create slice:

```go
var s []int
// or
s := make([]int, 2, 4) // length is 2, capacity is 4
```

copy slice:

```go
copy(des, src) // the copied items are determined by des's length
```

delete item form slice:

```go
// delete no.3 item from slice
s2 = append(s2[:3], s2[4:]...)
```

pop from front of slice:

```go
s = s[1:]
```

pop from tail of slice:

```go
s = s[:len(s)-1]
```


# map

```go
m := map[string]string {
    "name": "ccmouse",
    "course": "golang",
}
// or
m := make(map[string]int) // empty map
// or
var m3 map[string]int // nil
```

traverse

```go
for k, v := range m {

}
```

get

```go
value, ok := m["key"]
// or

if value, ok := m["key"]; ok {

} else {

}
```

delete

```go
delete(m, "name")
```

## map's key

* must be comparable
* expert `slice`, `map`, `function`
* `struct` type without types above

## demo

find longest substr doesn't contain duplicated character

```go
func longestSubStr(s string) (int, int) {
	lastOccured := make(map[rune]int)
	start := 0
	maxLength := 0
	maxStart := 0
	for i, ch := range []rune(s) {
		lastI, ok := lastOccured[ch]
		if ok && lastI >= start {
			start = lastOccured[ch] + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
			maxStart = start
		}
		lastOccured[ch] = i
	}
	return maxStart, maxLength
}
```

## rune

```go
s := "abc萨克国际"

for _, b := range []byte(s) {
    fmt.Printf("%X ", b)
}
fmt.Println()
// 61 62 63 E8 90 A8 E5 85 8B E5 9B BD E9 99 85

for i, ch := range s { // ch is a rune
    fmt.Printf("(%d %x)", i, ch)
}
fmt.Println()
// (0 61)(1 62)(2 63)(3 8428)(6 514b)(9 56fd)(12 9645)

fmt.Println("Rune count:", utf8.RuneCountInString(s))
// Rune count: 7

bytes := []byte(s)
for len(bytes) > 0 {
    ch, size := utf8.DecodeRune(bytes)
    bytes = bytes[size:]
    fmt.Printf("%c ", ch)
}
fmt.Println()
// a b c 萨 克 国 际

for i, ch := range []rune(s) {
    fmt.Printf("(%d %c %X)", i, ch, ch)
}
fmt.Println()
// (0 a 61)(1 b 62)(2 c 63)(3 萨 8428)(4 克 514B)(5 国 56FD)(6 际 9645)
```


# struct and function
* only support encapsulation
* donn't support inherit and polymorphic

```go
type treeNode struct {
	value int
	left, right *treeNode
}

var root treeNode

root = treeNode{value:3}
root.left = &treeNode{}
root.right = &treeNode{5, nil, nil}
root.right.left = new(treeNode)

nodes := []treeNode {
    {value: 3},
    {},
    {6, nil, &root},
}
```

or use factory function

```go
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}
```

note that it return the local pointer


Where the struct create at, stack or heap?
Answer: You donn't need to know.



receiver

```go
func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(value int) {
	node.value = value
}
```


nil can call function too



## when to use value receiver or pointer receiver

* if you want to change content, use pointer
* if struct is too large, use pointer



# package and encapsulation


# extend existed type
## alias
```go
type Queue []int

func (q *Queue) Push(v int) {
	*q = append(*q, v)
}

func (q *Queue) Pop() int {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
```

## assembly

for example, we want to extend the TreeNode

```go
type MyTreeNode struct {
	node *TreeNode
}

func (myNode *MyTreeNode) postOrder () {
	if myNode == nil || myNode.node == nil {
		return
	}
	MyTreeNode{myNode.node.Left}.postOrder()
	MyTreeNode{myNode.node.Right}.postOrder()
	myNode.node.Print()
}
```


# orientating interface programing

```go
type Retriever interface {
	Get(url string) string
}

func download(r Retriever) string {
	return r.Get("www.baidu.com")
}

type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
	return r.Contents
}

func main() {
	r := MockRetriever{"I am content."}
	fmt.Println(download(r))
}
```



c ++ duck typing

```c++
template <class R>
string download(const R& retriever) {
    return retriever.get('www.baidu.com')
}
```

Dont know if has get function until compiling


java

```java
<R extends Retriever>
String download(R r) {
    return r.get("www.baidu.com")
}
```


Use switch to decide type

```go
switch r.(type) {
case MockRetriever:
    fmt.Println("I am mock retriever")
case RealRetriever:
    fmt.Println("I am real retriever")
}
```


Or use assert

```go
if realRetriever, ok := r.(RealRetriever); ok {
    fmt.Println(realRetriever.UserAgent)
} else {
    fmt.Println("not real retriever")
}
```


Interface combination

```go
// interface
type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type PosterRetriever interface {
	Retriever
	Poster
}

func session(s PosterRetriever, url string) {
...
}

func download(r Retriever, url string) string {
...
}

func post(poster Poster, url string) {
...
}

// implement interface, the MockRetriever implement Get Post, so it is implement all the interfaces above
type MockRetriever struct {
	Contents string
}

func (r MockRetriever) Get(url string) string {
    ...
}

func (r MockRetriever) Post(url string, form map[string]string) string {
    ...
}

var r PosterRetriever
r = MockRetriever{"I am content."}

session(r, url)
download(r, url)
post(r, url)
```


# functional programming

```go
func adder() func(int) int {
	sum := 0
	return func(i int) int {
		sum += i
		return sum
	}
}

func main() {
	a := adder()
	for i := 1; i <= 36; i++  {
		 fmt.Println(a(i))
	}
}
```


Or

```go
type iAdder func(int) (int, iAdder)

func adder2(base int) iAdder {
	return func(v int) (int, iAdder) {
		return base + v, adder2(base + v)
	}
}

func main() {
	a2 := adder2(0)
	for i := 0; i < 10; i++ {
		var s int
		s, a2 = a2(i)
		fmt.Println(s)
	}
}
```

python closure

```python
// python3
def adder():
    sum = 0

    def f (value):
        nonlocal sum
        sum += value
        return sum

    return f

// python2
def adder():
  obj = {
    "sum": 0
  }

  def f(value):
    obj["sum"] += value
    return obj["sum"]

  return f
```

## demo1 fabnaci

```go
func fibnaci() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}
```

implement interface for func

```go
type intGen func() int

func (g intGen) Read(p []byte) (n int, err error) {
	next := g()
	if next > 1000 {
		return 0, io.EOF
	}
	s := fmt.Sprintf("%d\n", next)
	return strings.NewReader(s).Read(p)
}

func fibnaci() intGen {
	a, b := 0, 1
	return func() int {
		a, b = b, a + b
		return a
	}
}

func printContents(reader io.Reader) {
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

```


# resource manage & handle error
## defer
```go
func f () int {
	a := 1

	defer func() {
		a = 2
		fmt.Println("defer", a)
	}()

	return a // 1 set the reture value (equal a) 2 execute defer func 3 return
}

func main() {
	fmt.Println(f()) // 1
}
```


the defer stack

```go
func tryDefer() {
	defer fmt.Println(1)
	defer fmt.Println(2)
	fmt.Println(3)
}

3
2
1
```

## handle error
```go
	file, err := os.OpenFile(filename, os.O_EXCL|os.O_CREATE, 0666)
	if err != nil {
		if pathError, ok := err.(*os.PathError); !ok {
			panic(err)
		} else {
			fmt.Println(pathError.Op, pathError.Path, pathError.Err)
		}
		return
	}
```

custom error

```go
errors.New("this is a custom error")
```


## panic and recover
* panic
1. stop execution of function
2. return to up, execute every defer
3. if did not meet recover, exit

* recover
1. only call in defer
2. get panic value
3. if cannot handle, can re-panic

usage

```go
func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		}
	}()
	b := 0
	a := 5 / b
	fmt.Println(a)
}

func main() {
	tryRecover()
}

```

re-panic

```go
func tryRecover() {
	defer func() {
		r := recover()
		if err, ok := r.(error); ok {
			fmt.Println("error occurred:", err)
		} else {
		    // cannot handle, panic up
			panic(r)
		}
	}()
	panic(123)
}

func main() {
	tryRecover()
}
```

## uniform error handling
decoration pattern

```go
package main

import (
	"net/http"
	"os"
	"io/ioutil"
	"strings"
	"log"
)

const prefix = "/list/"

type userErrorI interface {
	error
	Message() string
}

type userError string

func (e userError) Error() string {
	return e.Message()
}

func (e userError) Message() string {
	return string(e)
}

func HandleFileList (writer http.ResponseWriter, request *http.Request) error {
	if strings.Index(request.URL.Path, prefix) != 0 {
		return userError("Path must start with " + prefix)
	}

	path := request.URL.Path[len(prefix):]

	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	all, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	writer.Write(all)
	return nil
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()

		err := handler(writer, request)
		if err != nil {
			code := http.StatusOK

			if userError, ok := err.(userErrorI); ok {
				http.Error(writer, userError.Message(), http.StatusBadRequest)
				return
			}
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			case os.IsPermission(err):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(HandleFileList))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
```




