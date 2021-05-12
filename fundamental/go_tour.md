# 00. What is Go

## - Go Developers

### * Robert Griesemer

구글 JS V8 엔진 개발 참여자

### * Rob Pike

벨 연구소, UTF-8 개발자이자 유닉스 개발자

### * Ken Thompson

벨 연구소, 유닉스 개발, B언어 개발 

<br><br><br>

## - Go 특징

 - 컴파일 언어
 - GC 기능 지원
 - Concurrent 프로그래밍 지원 

<br><br><br>

## - Go 환경

### * go env

GO 환경변수 확인

`GOROOT` : GO 실행파일(bin)과 표준 패키지(pkg)가 있음 [참고](https://golang.org/pkg/)

`GOPATH` : 서드파티 패키지, 사용자 정의 패키지 위치 설정

### * go run

컴파일과 동시에 실행. 실행 파일은 만들어지지 않는다.

### * go build

컴파일 + 실행 파일 만듬

### * Workspace 구조

#### GOROOT

 - bin : 실행파일
 - pkg : 표준 패키지
 - src : 소스코드

#### GOPATH : 서드파티

<br><br><br><br><br>
<br><br><br><br><br>

# 01. Hello Wolrd

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
}
```

<br><br><br><br><br>
<br><br><br><br><br>

# 02. 변수, 상수 그리고 자료형


## - 변수

```go
var a int
var f float32 = 11.0
var i, j, k int
var i, j, k int = 1, 2, 3
```

### * 변수 선언 축약

타입은 알아서 매칭해준다.

```go
intro := "Hello World"
```

<br><br><br>

## - 상수

```go
const c int = 10
const s string = "Hello World"
const (
    a = 10
    b = "string"
    c = 3.14
)
```

<br><br><br>

## - 자료형

 - bool
 - string (immutable)
 - int, int8, int16, int32, int64
 - uint, uint8, uint16, uint32, uint64, uintptr
 - float32, float64, complex64, complex128
 - byte = uint8
 - rune = int32

### * 자료형 변환

```go
str := "ABC"
bytes := string(str)
str2 := string(bytes)
```

<br><br><br>

### - 문자열
 - Back Quote(\`  \`) : raw
 - (" ") : string

<br><br><br><br><br>
<br><br><br><br><br>

# 03. 연산자

## - 산술 연산자

<br><br><br>

## - 관계 연산자

<br><br><br>

## - 논리 연산자

<br><br><br>

## - Bitwise 연산자

<br><br><br>

## - 할당 연산자 + 복합 할당 연산자

<br><br><br>


## - 포인터 연산자(*), 주소 연산자(&)


<br><br><br><br><br>
<br><br><br><br><br>

# 04. 제어문

## - 조건문

### * if ~ else

```go
num := 20
if num == 20 {
    println("It is true")
}else{
    println("It is false")
}
```

### * else if

```go
score := 78
if score >= 90 {
    println("수")
}else if score >= 80{
    println("우")
}else if score >= 70{
    println("미")
}else if score >= 60{
    println("양")
}else{
    println("가")
}
```

### * switch ~ case ~ default

```go
choice := 2

switch choice {
    case 1:
        println("selected one")
    case 2:
        println("selected two")
    case 3:
        println("selected third")
    case 4:
        println("selected four")
    default:
        println("default!!!")
}
```

<br><br><br>

## - 반복분 

### * for 초기화 ; 조건식 ; 증감식

```go
sum := 0

for i := 1 ; i <= 10 ; i++ {
    sum += i
}

println(sum) //5
```

### * for 조건식 (while 처럼 사용)

```go
sum := 0
i := 0

for i < 100 {
    i++
    sum += i
}

println(sum)
```

### * 무한루프

```go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var choice int
	isExit := false

	for {
		fmt.Print("입력하세요 : ")
		fmt.Scan(&input)

		choice, _ = strconv.Atoi(input)

		switch choice {
		case 1:
			println("selected one")
		case 2:
			println("selected two")
		case 3:
			println("selected third")
		case 4:
			isExit = true
		default:
			println("잘못 입력!")
		}

		if isExit {
			println("종료합니다.")
			break
		}
	}
	println("---종료됨---")
}
```

### * for range

```go
package main

func main() {
	nums := []int{10, 20, 30, 40, 50}

	for index, num := range nums {
		println(index, " : ", num)
	}
}
```

### * break, continue

<br><br><br><br><br>
<br><br><br><br><br>

# 05. 함수

## - 함수 선언과 사용

### * return 

```go
package main

import "fmt"

func main() {
	add, sub, mul, div := calc(10, 3)
	fmt.Println(add, " : ", sub, " : ", mul, " : ", div) // 13  :  7  :  30  :  3
}

func calc(num1 int, num2 int) (int, int, int, int) {
	return num1 + num2, num1 - num2, num1 * num2, num1 / num2
}
```

### * 가변인자

```go
package main

import "fmt"

func main() {
	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) //10 55
}

func add(nums ...int) (count int, total int) {
	for _, n := range nums {
		total += n
	}
	count = len(nums)
	return
}
```

<br><br><br>

## - call by value, call by reference

### * call by value

```go
package main

func main() {
	msg := "Hello"
	changeMsg(msg)
	println(msg) //Hello
}

func changeMsg(msg string) {
	msg = "Hello World"
}
```

### * call by refernce

```go
package main

func main() {
	msg := "Hello"
	changeMsg(&msg)
	println(msg) //Hello World
}

func changeMsg(msg *string) {
	*msg = "Hello World"
}
```

<br><br><br>

## - 익명함수

 - case 1

```go
package main

import "fmt"

func main() {
	add := func(nums ...int) (count int, total int) {
		for _, n := range nums {
			total += n
		}
		count = len(nums)
		return
	}

	fmt.Println(add(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) //10 55
}
```
 - case 2

```go
package main

import "fmt"

func main() {
	fmt.Println(func(nums ...int) (count int, total int) {
		for _, n := range nums {
			total += n
		}
		count = len(nums)
		return
	}(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)) //10 55
}
```

<br><br><br>

## - Delegate

함수의 원형을 정의하고 타 메서드에 전달해 리턴받는 기능을 Delegate라 부른다.

### * 매개 변수의 함수

```go
package main

import "fmt"

func main() {
	add := func(nums []int) (count int, total int) {
		total = 0
		for _, n := range nums {
			total += n
		}
		count = len(nums)
		return
	}

	mul := func(nums []int) (count int, total int) {
		total = 1
		for _, n := range nums {
			total *= n
		}
		count = len(nums)
		return
	}

	fmt.Println(calc(add, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(calc(mul, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func calc(f func([]int) (int, int), nums ...int) (int, int) {
	return f(nums)
}
```

### * Type 정의

```go
package main

import "fmt"

type calculator func([]int) (int, int)

func main() {

	add := func(nums []int) (count int, total int) {
		total = 0
		for _, n := range nums {
			total += n
		}
		count = len(nums)
		return
	}

	mul := func(nums []int) (count int, total int) {
		total = 1
		for _, n := range nums {
			total *= n
		}
		count = len(nums)
		return
	}

	fmt.Println(calc(add, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(calc(mul, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func calc(f calculator, nums ...int) (int, int) {
	return f(nums)
}
```

<br><br><br>

## - 클로저(Closure)

```go
package main

type increment func() int

func main() {
	next := nextValue()

	println(next()) //1
	println(next()) //2 
	println(next()) //3

	anotherNext := nextValue()

	println(anotherNext()) //1
	println(anotherNext()) //2
	println(anotherNext()) //3

}

func nextValue() increment {
	i := 0
	return func() int {
		i++
		return i
	}
}
```
# 06. 컬렉션

## - 배열

```go
package main

import "fmt"

func main() {
	var nums [5]int
	nums[0] = 10
	nums[1] = 20
	nums[2] = 30
	nums[3] = 40
	nums[4] = 50

	// println(nums) // ERROR
	fmt.Println(nums) //[10 20 30 40 50]
}
```

```go
package main

import "fmt"

func main() {

	var nums = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(nums) // [[1 2 3] [4 5 6]]
}
```

<br><br><br>

## - Slice

배열처럼 크기가 지정되어 있지않고 동적으로 변한다

```go
var nums []int
nums = []int{1, 2, 3}
println(len(nums), cap(nums)) //3 3
```

### * append and copy

 - append()

`size`와 `capacity`를 정할 수 있으며 append를 통해 요소를 확장할 수 있고 `size`가 `capacity`를 넘으면 `capacity`가 2배 늘어난다.

```go
nums := make([]int, 5, 7)
println(len(nums), cap(nums)) //5, 7
fmt.Println(nums)

nums = append(nums, 10, 20)
println(len(nums), cap(nums)) //7, 7
fmt.Println(nums)

nums = append(nums, 30)
println(len(nums), cap(nums)) //8, 14
fmt.Println(nums)
```

 - copy()

```go
package main

import "fmt"

func main() {

	nums := make([]int, 0, 10)

	nums = append(nums, 10, 20, 30, 40, 50)

	fmt.Println(nums) //[10, 20, 30, 40, 50]

	copyNums := []int{1, 2}

	copy(copyNums, nums)

	fmt.Println(copyNums) //[10, 20]

	copyNums2 := make([]int, len(nums), cap(nums)*2)

	copy(copyNums2, nums)

	fmt.Println(copyNums2) //[10 20 30 40 50]

	//5 10  :  2 2  :  5 20
	fmt.Println(
		len(nums), cap(nums),
		" : ",
		len(copyNums), cap(copyNums),
		" : ",
		len(copyNums2), cap(copyNums2),
	)

}
```

<br><br><br>


## - Map

### * make and delete

```go
var calc map[string]calculator //[key]value

calc = make(map[string]calculator)

calc["add"] = func(num1 int, num2 int) int {
    return num1 + num2
}

calc["mul"] = func(num1 int, num2 int) int {
    return num1 * num2
}

fmt.Println(calc["add"](100, 200)) //300
fmt.Println(calc["mul"](22, 33))   //726

delete(calc, "mul")

if calc["mul"] == nil { //There is no "mul" key
    println("There is no \"mul\" key")
}
```

### * 선언 축약 방식과 Map 체크 및 열거

```go
myMap := map[int]string{
    10: "TEN",
    5:  "FIVE",
    3:  "THREE",
}

fmt.Println(myMap) //map[3:THREE 5:FIVE 10:TEN]

val, isExist := myMap[10]

fmt.Println(val, isExist) // TEN true

for key, val := range myMap {
    fmt.Println(key, " : ", val)
}

val, isExist = myMap[11]

fmt.Println(val, isExist) //   false
if val == "" {            //val은 ""
    fmt.Println("val은 \"\"")
}
```

<br><br><br><br><br>
<br><br><br><br><br>

# 07. 패키지

## - package `main `

main 패키지는 Go Compiler에 의해 특별하게 인식됨. (Entry Point)

실행 프로그램으로 만들어지게 됨

<br><br><br>

## - import

- 표준 라이브러리 위치 : `GOROOT/pkg` (설치시 자동으로 환경변수 설정됨)
- 사용자 패키지 위치 : `GOPATH/pkg` (사용자가 따로 환경변수 설정해줘야 함)

## - Scope

 - Public : 첫글자가 대문자
 - Private : 첫글자가 소문자(패키지 내에서만 사용 가능)

## - `init()`

패키지 실행시 처음으로 호출되는 init() 함수

```go
import(
	alice "mylib/function/good"
	bob "mylib/function/good"
	_ "mylib/function/initialize"
)
```

alias중에서 `_`은 `init()`만 실행하고 싶을 떄 사용한다.


<br><br><br><br><br>
<br><br><br><br><br>

# 08. 구조체

## - 사용자 자료형

 - case 1
```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p := person{}

	p.name = "RHIE"
	p.age = 20

	fmt.Println(p)
}
```
- case 2
```go
package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var p person
	p = person{"Alice", 10}
	fmt.Println(p)
}
```

## - 메서드

func와 함수명 사이에 어떤 struct의 메서드인지 표시

```go
package main

type Rect struct {
	width, height int
}

func (r Rect) area() int {
	return r.height * r.width
}

func main() {
	rect := Rect{10, 20}
	area := rect.area()
	println(area) //200
}
```

### * Value Receiver & Pointer Receiver

```go
package main

import "fmt"

type Number struct {
	num int
}

func (n Number) valueRecv() int {
	n.num++
	return n.num
}

func (n *Number) pointerRecv() int {
	n.num++
	return n.num
}

func newNum() *Number {
	var n Number
	n.num = 10
	return &n
}

func main() {
	vr := newNum()
	pr := newNum()

	vr_result := vr.valueRecv()
	pr_result := pr.pointerRecv()

	fmt.Println(vr_result, vr.num) //11 10
	fmt.Println(pr_result, pr.num) //11 11
}
```

<br><br><br><br><br>
<br><br><br><br><br>

# 09. 인터페이스

구현해야하는 메서드 원형들의 집합

```go
package main

type MyStr string

type MyInt int

func (s MyStr) show() {
	println("----", s, "----")
}

func (i MyInt) show() {
	println("<", i, ">")
}

type Prnt interface {
	show()
}

func main() {
	var s MyStr = "I love Golang"
	var i MyInt = 100

	s.show()
	i.show()

	s = "Really"
	i = 200

	var p Prnt
	p = s
	p.show()
	p = i
	p.show()

	s = "OK"
	i = 300

	printer(s, i)
}

func printer(p ...Prnt) {
	for _, v := range p {
		v.show()
	}
}
```

OUTPUT
```
---- I love Golang ----
< 100 >
---- Really ----
< 200 >
---- OK ----
< 300 >
```

## - Object와 같은 Empty Interface와 타입 설정

```go
package main

import "fmt"

func main() {
	var obj interface{}
	obj = "I love Golang"
	obj = 100

	fmt.Println(obj) //100
	println(obj)     //(0x9c2aa0,0x9fec58)

	other := obj

	fmt.Println(other) //100
	println(other)     //(0x9c2aa0,0x9fec58)

	other2 := other.(int) /* 타입 설정 */
	fmt.Println(other2 + 1) //101
	println(other2 + 1)     //101

}
```

<br><br><br><br><br>
<br><br><br><br><br>

# 10. 에러처리

내장으로 error라는 타입이 있음
```go
type error interface{
	Error() string
}
```

내장 패키지인 log는 이 error 인터페이스를 받아서 구현함

```go
package main

import (
	"fmt"
	"log"
)

type myErr struct {
	errMsg string
}

func (n myErr) Error() string {
	return n.errMsg
}

func main() {

	cuserr := myErr{errMsg: "HEY!!! It's error!!!"}
	log.Fatal(cuserr.Error()) // 2021/05/12 22:25:42 HEY!!! It's error!!!
	//log.Fatal()은 os.Exit(1)을 호출하여 프로그램을 종료시킨다.
	fmt.Println("안녕하세요") //출력되지 않는다.

}
```

`nil`인지 아닌지 판단하여 예외처리를 하자

## - 사용자 에러 만들기와 예외처리


[Continue](http://golang.site/go/article/19-Go-%EC%97%90%EB%9F%AC%EC%B2%98%EB%A6%AC)