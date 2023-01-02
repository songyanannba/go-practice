package main

import (
	"fmt"
	"net/http"
)

//web-beego

func main() {

	syn := func(resp http.ResponseWriter, r *http.Request) {
		fmt.Println("request ", r)

		fmt.Fprint(resp, "欢迎。。")
	}

	/*http.Cookie{
		Name:       "",
		Value:      "",
		Path:       "",
		Domain:     "",
		Expires:    time.Time{},
		RawExpires: "",
		MaxAge:     0,
		Secure:     false,
		HttpOnly:   false,
		SameSite:   0,
		Raw:        "",
		Unparsed:   nil,
	}*/

	http.HandleFunc("/time", syn)

	http.ListenAndServe(":9999", nil)

}
