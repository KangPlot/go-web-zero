# go web zero（quickly start）

[source video](https://www.bilibili.com/video/BV1Xv411k7Xn?p=7)

## p6

query 操作
方法1：通过map，返回字符串切片
id := query["id"]

方法2：通过Get方法，返回key对应的第一个值
name := query.Get("name")

~~~go
package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		url := r.URL
		query := url.Query()

		// 方法1：通过map，返回字符串切片
		id := query["id"]
		log.Println(id)

		// 方法2：通过Get方法，返回key对应的第一个值
		name := query.Get("name")
		log.Println(name)

	})
	http.ListenAndServe("localhost:8080", nil)
}
~~~


## p7 
[source video](https://www.bilibili.com/video/BV1Xv411k7Xn?p=7)

![img.png](img.png)

~~~go
package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		// r.Form 出现情况，url 里面的key 和 表单里面相同的话，会一起显示出来
		fmt.Fprintln(w, r.PostForm)
	})
	server.ListenAndServe()
}
~~~
结果： 使用r.Form

![img_1.png](img_1.png)

结果：使用r.PostForm

![img_2.png](img_2.png)


index.html
~~~html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>test Form</title>
</head>
<body>
<form action="http://localhost:8080/process?first_name=Tom" method="post" enctype="application/x-www-form-urlencoded">
    first name
    <input type="text" name="first_name" />
    last name
    <input type="text" name="last_name" />
    <input type="submit" />
</form>

</body>
</html>
~~~

## p8 上传文件
方法1
![img_3.png](img_3.png)



~~~go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(1024)

	fileHeader := r.MultipartForm.File["uploaded"][0] // 获取上传的第一个文件
	file, err := fileHeader.Open()
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
~~~


成功读取文件
![img_5.png](img_5.png)

方法2
![img_4.png](img_4.png)

~~~go
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//r.ParseMultipartForm(1024)
	//
	//fileHeader := r.MultipartForm.File["uploaded"][0] // 获取上传的第一个文件
	//file, err := fileHeader.Open()

	file, _, err := r.FormFile("uploaded") // 传入的html form里面对应的key
	// 只会返回第一个文件，适合单文件上传
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}

}
func main() {
	server := http.Server{
		Addr: "localhost:8080",
	}
	http.HandleFunc("/process", process)
	server.ListenAndServe()
}
~~~

index.html
~~~html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>load file</title>
</head>
<body>
    <form action="http://localhost:8080/process?hello=world&thread=123" method="post" enctype="multipart/form-data">
        first name
        <input type="text" name="hello" value="gogo"/>
        last name
        <input type="text" name="post" value="zhuanma"/>
        file
        <input type="file" name="uploaded">
        <input type="submit" />
    </form>

</body>
</html>
~~~

解析JSON（视频里留坑，没有讲）

![img_6.png](img_6.png)

