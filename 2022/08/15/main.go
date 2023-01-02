package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHello(w http.ResponseWriter , r *http.Request) {
	_ = r.ParseForm() //解析参数 默认不会解析
	fmt.Println("form: ",r.Form)
	fmt.Println("Path: " , r.URL.Path)
	fmt.Println("Host: " , r.Host)

	for k, v := range r.Form {
		fmt.Println("key: " , k)
		fmt.Println("val: " , strings.Join(v,""))
	}

	_, _ = fmt.Fprintf(w, "hello web ,%s", r.Form.Get("name"))

}


type User struct {
	Name string
	Hab []string
}

func write(w http.ResponseWriter , r *http.Request) {
	w.Header().Set("Content-Type" , "application/json")
	w.Header().Set("X-Custom-Header" , "custom")

	w.WriteHeader(222)

	user := User{
		Name: "sssooonnnggg",
		Hab:  []string{"sss" ,"yyy" ,"nnn"},
	}

	marshal, _ := json.Marshal(user)
	w.Write(marshal)

}

//构建服务器
func main() {
	//http.HandleFunc("/" , sayHello)
	http.HandleFunc("/write" , write)
	err := http.ListenAndServe(":8880" ,nil)
	if err != nil {
		log.Fatal("监听服务器报错： ",err)
	}



}