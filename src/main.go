package main

import (
	"net/http"
	"text/template"
	"fmt"
)

func main() {
	http.HandleFunc("/",test)
	http.ListenAndServe(":8080",nil)
}

func test(w http.ResponseWriter,r *http.Request)  {
	t,err := template.ParseFiles("./static/index.html")
	if err !=nil{
		fmt.Println(err)
		w.Write([]byte("failed"))
		return
	}

	t.Execute(w,nil)
}
