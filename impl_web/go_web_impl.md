# 00. 사전준비

## - `static` 폴더 경로 안에

### * HTML
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="stylesheet" href="./static/style.css">
</head>
<body>
    <h1>Hello World</h1>
    <script src="./static/script.js"></script>
</body>
</html>
```
### * CSS

```css
h1 {
    color : red;
}
```
### * JavaScript
```js
console.log("자바스크립트 불러오기 성공")
```


<br><br><br><br><br>
<br><br><br><br><br>

# 01. Hello World


## - 웹페이지 불러오기
```go
package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {

	//파일시스템 포인팅 : css, js
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)

    //웹페이지
    welcome_page := "static/index.html"

	//라우터
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		webpage, _ := ioutil.ReadFile(welcome_page)
		res.Header().Set("Content-Type", "text.html")
		res.Write(webpage)
	})

	log.Println("서버 시작")
	http.ListenAndServe(":5000", nil)
}
```

<br><br><br>

## - JSON 가져오기

### * JSON 인코딩 & 디코딩

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Data struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

func main() {
	dataToJson := Data{"Alice", 100}

	fmt.Println(dataToJson) //{Alice 100}

	//인코딩
	jsonBytes, err := json.Marshal(dataToJson)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonBytes) //[123 34 115 116 114 34 58 34 65 108 105 99 101 34 44 34 110 117 109 34 58 49 48 48 125]

	//디코딩
	var jsonToData Data
	err = json.Unmarshal(jsonBytes, &jsonToData)
	if err != nil {
		panic(err)
	}

	fmt.Println(jsonToData) //{Alice 100}
}
```

<br><br><br>

## -  JSON 보내기

```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

func main() {

	//파일시스템 포인팅 : css, js
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)

	//웹페이지
	welcome_page := "static/index.html"

	//라우터
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {

		webpage, _ := ioutil.ReadFile(welcome_page)
		res.Header().Set("Content-Type", "text.html")
		res.Write(webpage)
	})

	http.HandleFunc("/json", func(res http.ResponseWriter, req *http.Request) {
		doc := Data{"Hello", 100}

		//JSON 인코딩
		body, _ := json.Marshal(doc)

		res.Header().Set("Content-Type", "application/json")
		res.Write(body)
	})

	log.Println("서버 시작")
	http.ListenAndServe(":5000", nil)
}
```

<br><br><br>

## -  핸들러 정리하기


```go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type Data struct {
	Str string `json:"str"`
	Num int    `json:"num"`
}

//웹페이지
var welcome_page string = "static/index.html"

func main() {

	//파일시스템 포인팅 : css, js
	fs := http.StripPrefix("/static/", http.FileServer(http.Dir("static")))
	http.Handle("/static/", fs)

	//라우터
	http.HandleFunc("/", mainPage)
	http.HandleFunc("/json", sendJson)

	//서버가동
	log.Println("서버 시작")
	http.ListenAndServe(":5000", nil)
}

func mainPage(res http.ResponseWriter, req *http.Request) {
	webpage, _ := ioutil.ReadFile(welcome_page)
	res.Header().Set("Content-Type", "text.html")
	res.Write(webpage)
}

func sendJson(res http.ResponseWriter, req *http.Request) {
	doc := Data{"Hello", 100}

	//JSON 인코딩
	body, _ := json.Marshal(doc)

	res.Header().Set("Content-Type", "application/json")
	res.Write(body)
}
```

<br><br><br>

## -  비동기 처리하기

파일 수정하기

### * HTML
```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
    <link rel="stylesheet" href="./static/style.css">
</head>
<body>
    <h1>Hello World</h1>
    <p class="show-json"></p>
    <script src="./static/script.js"></script>
</body>
</html> 
```
### * CSS

```css
h1 {
    color : red;
}

.show-json {
    background-color: yellow;
    color : blue;
}
```
### * JavaScript
```js
console.log("자바스크립트 불러오기 성공")

const show_json = document.getElementsByClassName("show-json")[0];
const newDIV = document.createElement( 'div' );

fetch('http://localhost:5000/json')
  .then(function(response) {
    return response.json();
  })
  .then(function(data) {
    console.log(data) //{str: "Hello", num: 100}
    
    newDIV.innerHTML = data.str + ", " + data.num;
    newDIV.setAttribute("class","myDiv");
    show_json.appendChild(newDIV);
    
  });
```