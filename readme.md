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

# Test
Tabel drive test

```go
tests := []struct {
    a, b, c int32
}{
    {1, 2, 3},
    {0, 2, 2},
}

for _, test := range tests {
    if actual := add(test.a, test.b); actual != test.c {
        ...
    }
}
```

Benchmark:

```go
func BenchmarkSubStr(b *testing.B) {
	s := "黑化肥挥发发灰会花飞灰化肥发发"

	for i := 0; i < 13; i++ {
		s = s + s
	}

	ans := 7

	b.Logf("len(s) = %d", len(s))

	for i := 0; i < b.N; i++ {
		_, actual := longestSubStr(s)
		if actual != ans {
			b.Errorf("got %d for input %s; " +
				"expected %d",
					actual, s, ans)
		}
	}
}
```


cpuprofile

```go
go test -bench . -cpuprofile cpu.out
go tool pprof cpu.out
(pprof) web
```

## http server test

```go
package main

import (
	"testing"
	"net/http/httptest"
	"net/http"
	"io/ioutil"
	"strings"
	"os"
	"errors"
	"fmt"
)

func errPanic(writer http.ResponseWriter, request *http.Request) error {
	panic(122)
}

type testingUserError string

func (e testingUserError) Error() string {
	return e.Message()
}

func (e testingUserError) Message() string {
	return string(e)
}

func errUserError(writer http.ResponseWriter,
	request *http.Request) error {
	return testingUserError("user error")
}

func errNoPermission(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrPermission
}

func errUnknown(writer http.ResponseWriter,
	request *http.Request) error {
	return errors.New("unknown error")
}

func errNotFound(writer http.ResponseWriter,
	request *http.Request) error {
	return os.ErrNotExist
}

func noError(writer http.ResponseWriter,
	request *http.Request) error {
	fmt.Fprintln(writer, "no error")
	return nil
}

var tests = []struct {
	h appHandler
	code int
	message string
}{
	{errPanic, 500, "Internal Server Error"},
	{errUserError, 400, "user error"},
	{errNoPermission, 403, "Forbidden"},
	{errNotFound, 404, "Not Found"},
	{errUnknown, 500, "Internal Server Error"},
	{noError, 200, "no error"},
}

func TestErrWrapper (t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		response := httptest.NewRecorder()
		request := httptest.NewRequest(
			http.MethodGet,
			"http://www.baidu.com", nil)
		f(response, request)

		verifyResponse(response.Result(), tt.code, tt.message, t)
	}
}

// a real server
func TestErrWrapperInServer(t *testing.T) {
	for _, tt := range tests {
		f := errWrapper(tt.h)
		server := httptest.NewServer(http.HandlerFunc(f))
		response, _ := http.Get(server.URL)
		verifyResponse(response, tt.code, tt.message, t)
	}
}

func verifyResponse(response *http.Response,
	expectedCode int, expectedMsg string, t *testing.T) {
	b, _ := ioutil.ReadAll(response.Body)
	body := strings.Trim(string(b), "\n")
	if response.StatusCode != expectedCode ||
		body != expectedMsg {
		t.Errorf("expect (%d, %s); " +
			"got (%d, %s)", expectedCode, expectedMsg, response.StatusCode, body)
	}
}
```

# doc
'''
godoc -http :6060
'''

example:

```go
package queue

import "fmt"

func ExampleQueue_Pop() {
	q := Queue{1}
	q.Push(1)
	q.Push(2)
	fmt.Println(q.Pop())
	fmt.Println(q.Pop())
	fmt.Println(q.IsEmpty())

	// Output:
	// 1
	// 1
	// false
}
```

## Test Summary
* table drive
* coverage
* profile
* http testing
* doc and example

# goroutine
## coroutine
* lightweight
* non-preemptive multi-task processing, return control by itself
* multi-task on compiler/analyser/virtual-machine level, not on OS

e.g. The goroutine cannot return the control, so the process is in infinite loop

```go
package main

import (
	"time"
	"fmt"
)

func main() {
	var a [10]int
	for i := 0; i < 10; i++ {
		go func(i int) {
			for {
				a[i]++
			}
		}(i)
	}
	time.Sleep(time.Second * 5)
	fmt.Println(a)
}
```

can use `runtime.Gosched()` to return the control


**`go run --race`**

Monitor race.


## switch point
* I/O, select
* channel
* waiting lock
* func calling
* runtime.Gosched()

# channel
## usage
```go
package main

import (
	"time"
	"fmt"
)

func worker (id int, c chan int) {
	//for {
	//	n, ok := <-c
	//	if !ok {
	//		break
	//	}
	//	fmt.Printf("Worker %d received %c\n", id, n)
	//}
	// or
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func chanDemo () {
	var channels [10]chan<- int
	for i := 0; i < 10; i++ {
		channels[i] = createWorker(i)
	}

	for i := 0; i < 10; i++ {
		channels[i] <- 'a' + i
	}
}

func bufferedChannel () {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
}

// if closed, can still receive data, but the value is 0, can use it to
// decide if the channel is closed.
// if not closed, it will block
func channelClose () {
	c := make(chan int, 3)
	go worker(0, c)
	c <- 'a'
	c <- 'b'
	c <- 'c'
	c <- 'd'
	close(c)
}

func main () {
	fmt.Println("Channel as first-class citizen.")
	//chanDemo()
	fmt.Println("Buffered channel")
	//bufferedChannel()
	fmt.Println("Channel close and range")
	channelClose()
	time.Sleep(time.Second)
}
```

**Communication Sequential Process**

**Don't communicate by sharing memory, share memory by communication.**

## wait goroutine done
### use channel
The code above use sleep, we want to remove it, so we change the code to:

```go
package main

import (
	"fmt"
)

func doWork (id int, c chan int, done chan bool) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		done <- true
	}
}

type worker struct {
	in chan int
	done chan bool
}

func createWorker(id int) worker {
	w := worker{in: make(chan int), done: make(chan bool)}
	go doWork(id, w.in, w.done)
	return w
}

func chanDemo () {
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		<-worker.done
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
		<-worker.done
	}
}

func main () {
	fmt.Println("Channel as first-class citizen.")
	chanDemo()
}
```

It add a done channel to every work, and when work done, it will send signal to our main goroutine.
But it will result in the problem:

```
Channel as first-class citizen.
Worker 0 received a
Worker 1 received b
Worker 2 received c
Worker 3 received d
Worker 4 received e
Worker 5 received f
Worker 6 received g
Worker 7 received h
Worker 8 received i
Worker 9 received j
Worker 0 received A
Worker 1 received B
Worker 2 received C
Worker 3 received D
Worker 4 received E
Worker 5 received F
Worker 6 received G
Worker 7 received H
Worker 8 received I
Worker 9 received J
```

The worker now work in sequence! We can change our code to:

```go
	for i, worker := range workers {
		worker.in <- 'a' + i
	}

    // block here
	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	for _, worker := range workers {
		<-worker.done
	}
```

But it will block in that! So we need to do this:

```go
...
go func() { done <- true }()
...
```


### use WaitGroup
```
package main

import (
	"fmt"
	"sync"
)

func doWork (id int, c chan int, wg *sync.WaitGroup) {
	for n := range c {
		fmt.Printf("Worker %d received %c\n", id, n)
		wg.Done()
	}
}

type worker struct {
	in chan int
	wg *sync.WaitGroup
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int), wg: wg}
	go doWork(id, w.in, wg)
	return w
}

func chanDemo () {
	var workers [10]worker
	var wg sync.WaitGroup
	wg.Add(20)

	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i, &wg)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
	}

	for i, worker := range workers {
		worker.in <- 'A' + i
	}

	// wait for all of them
	wg.Wait()
}

func main () {
	fmt.Println("Channel as first-class citizen.")
	chanDemo()
}
```

More, we can change our code to functional programming:

```go
func doWork (id int, w worker) {
	for n := range w.in {
		fmt.Printf("Worker %d received %c\n", id, n)
		w.done()
	}
}

type worker struct {
	in chan int
	done func()
}

func createWorker(id int, wg *sync.WaitGroup) worker {
	w := worker{in: make(chan int), done: func() {
		wg.Done()
	}}
	go doWork(id, w)
	return w
}
...
```


## Use channel to traverse tree
```go
func (node *TreeNode) TraverseWithChannel() chan *TreeNode {
	out := make(chan *TreeNode)
	go func() {
		node.TraverseFunc(func(node *TreeNode) {
			out <- node // like yield in python
		})
		close(out)
	}()
	return out
}
```

## select
```go
package main

import (
	"fmt"
	"time"
	"math/rand"
)

func generator() chan int {
	out := make(chan int)
	go func() {
		i := 0
		for {
			time.Sleep(time.Duration(rand.Intn(1500)) * time.Millisecond)
			out <- i
			i++
		}
	}()
	return out
}

func worker (id int, c chan int) {
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)
	}
}

func createWorker(id int) chan<- int {
	c := make(chan int)
	go worker(id, c)
	return c
}

func main() {
	var c1, c2 = generator(), generator()
	var worker = createWorker(0)

	n := 0
	hasValue := false
	for {
		var activeWorker chan<- int
		if hasValue {
			activeWorker = worker
		}
		select {
		case n = <-c1:
			hasValue = true
		case n = <-c2:
			hasValue = true
		case activeWorker <- n: // nil chan will not be selected
			hasValue = false
		}
	}
}

```

There is a problem, if worker's process is too long, for example:

```go
...
		time.Sleep(time.Second * 5)
		fmt.Printf("Worker %d received %d\n", id, n)
...
```

We will lose data:

```go
Worker 0 received 0
Worker 0 received 7
```

So we need to cache the data produced:
```go
	var values []int
	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: // nil chan will not be selected
			values = values[1:]
		}
	}
```

We can use `time.After` and `time.Tick` to show more info:

```go
var values []int
	tm := time.After(10 * time.Second)
	tick := time.Tick(time.Second)
	n := 0
	for {
		var activeWorker chan<- int
		var activeValue int
		if len(values) > 0 {
			activeWorker = worker
			activeValue = values[0]
		}
		select {
		case n = <-c1:
			values = append(values, n)
		case n = <-c2:
			values = append(values, n)
		case activeWorker <- activeValue: // nil chan   not be selected
			values = values[1:]
		case <-time.After(800 * time.Millisecond):
			fmt.Println("Timeout")
		case <-tick:
			fmt.Println("Queue:")
			fmt.Println(values)
		case <-tm:
			fmt.Println("Bye")
			return
		}
	}
```

## Ordinary Synchronous
* WaitGroup
* Mutex
* Cond

```go
package main

import (
	"fmt"
	"time"
	"sync"
)

type atomicInt struct {
	value int
	lock sync.Mutex
}

func (a *atomicInt) increment () {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.value++
}

func (a *atomicInt) get() int {
	a.lock.Lock()
	defer a.lock.Unlock()
	return int(a.value)
}

func main() {
	var a atomicInt
	a.increment()
	go func() {
		a.increment()
	}()
	time.Sleep(time.Millisecond)
	fmt.Println(a.get())
}
```

If you want a zone to be locked?

```go
	func() {
		a.lock.Lock()
		defer a.lock.Unlock()
		a.value++
	}()
```

# Standard Lib
## http
```go
request, err := http.NewRequest(http.MethodGet, "http://www.imooc.com", nil)
	request.Header.Add("User-Agent", "Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	client := http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			fmt.Println("Redirect:", req)
			return nil
		},
	}

	resp, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
```

## pprof
`net/http/pprof`
`http://xxxxx/debug/pprof` to monitor system


# Other Standard Lib
`godoc -http :8888`


# Breadth First Search
My version：

```go
package main

import (
	"./queue"
	. "./point"
	"fmt"
)

// if the point can be add to stack
func ifAddPoint(x, y int, seen [][]bool, m [][]int, col, row int) bool {
	if x >= row || x < 0 || y >= col || y < 0 {
		return false
	}
	if m[x][y] == 1 || seen[x][y] {
		return false
	}
	return true
}

func neighPoints(x, y int) []Point {
	neighPoints := []Point{
		{X: x, Y: y - 1}, // up
		{X: x + 1, Y: y }, // right
		{X: x, Y: y + 1}, // bottom
		{X: x - 1, Y: y}, // left
	}

	return neighPoints
}

func bfs(m [][]int, col int, row int) [][]Point {
	// record the parent point of current point
	parentPointMap := [][]Point{}
	seen := [][]bool{}
	q := queue.Queue{}

	parentPointMap = make([][]Point, row)
	seen = make([][]bool, row)
	for i := 0; i < row; i++ {
		parentPointMap[i] = make([]Point, col)
		seen[i] = make([]bool, col)
	}

	q.Push(Point{0, 0})
	seen[0][0] = true

	for {
		item := q.Pop()
		if p, ok := item.(Point); ok {
			x, y := p.X, p.Y

			if x == row-1 && y == col-1 {
				break
      }

			for _, neighPoints := range neighPoints(x, y) {
				if (ifAddPoint(neighPoints.X, neighPoints.Y, seen, m, col, row)) {
					q.Push(neighPoints)
					nx, ny := neighPoints.X, neighPoints.Y
					q.Push(Point{X: nx, Y: ny})
					seen[nx][ny] = true
					parentPointMap[nx][ny] = Point{x, y}
				}
			}
		} else {
			break
		}
	}

	return parentPointMap
}


func main() {
	m := [][]int{
		{0, 1, 0, 0, 0},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
		{1, 1, 1, 0, 0},
		{0, 1, 0, 0, 1},
		{0, 1, 0, 0, 0},
	}
	col := 5
	row := 6
	res := bfs(m, col, row)

	x, y := 5, 4
	for {
		fmt.Printf("(%d, %d)", x, y)
		if x == 0 && y == 0 {
			break
		}
		fmt.Print("<-")
		point := res[x][y]
		x, y = point.X, point.Y
	}
}
```

Lesson version:

```go
package main

import (
	"fmt"
	"os"
)

func readMaze(filename string) [][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	var row, col int
	fmt.Fscanf(file, "%d %d", &row, &col)

	maze := make([][]int, row)
	for i := range maze {
		maze[i] = make([]int, col)
		for j := range maze[i] {
			fmt.Fscanf(file, "%d", &maze[i][j])
		}
	}

	return maze
}

var dirs = [4]point{
	{-1, 0}, {0, -1}, {1, 0}, {0, 1},
}

type point struct {
	i, j int
}

func (p point) add(r point) point {
	return point{p.i + r.i, p.j + r.j}
}

func (p point) at(grid [][]int) (int, bool) {
	if p.i < 0 || p.i >= len(grid) {
		return 0, false
	}

	if p.j < 0 || p.j >= len(grid[p.i]) {
		return 0, false
	}

	return grid[p.i][p.j], true
}

func walk(maze [][]int, start, end point) [][]int {
	steps := make([][]int, len(maze))
	for i := range steps {
		steps[i] = make([]int, len(maze[i]))
	}

	// queue
	Q := []point{start}

	for len(Q) > 0 {
		cur := Q[0]
		Q = Q[1:]

		if cur == end {
			break
		}

		for _, dir := range dirs {
			next := cur.add(dir)

			// maze at next is 0
			// and steps at next is 0 (not walked)
			// and next != start
			val, ok := next.at(maze)
			if !ok || val == 1 {
				continue
			}

			val, ok = next.at(steps)
			if !ok || val != 0 {
				continue
			}

			if next == start {
				continue
			}

			curSteps, _ := cur.at(steps)
			steps[next.i][next.j] = curSteps + 1

			Q = append(Q, next)
		}
	}

	return steps
}

func main() {
	maze := readMaze("./maze.in")

	steps := walk(maze, point{0, 0}, point{len(maze) - 1, len(maze[0]) - 1})

	for _, row := range steps {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}
```

# jsonrpc
```go
```

