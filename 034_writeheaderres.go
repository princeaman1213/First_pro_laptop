package main

import (
	"net/http"
	"fmt"
	"html/template"
)

type dog int

func(d dog) ServeHTTP(w http.ResponseWriter,r *http.Request){  //(lne1)    //this is the signature of handler interface
	w.Header().Set("Key","from me")
	w.Header().Set("Content-type","text/html ; char utf-8")
	fmt.Fprintln(w,"<h1>Any code here</h1>")
}

var t *template.Template

func init(){
	t=template.Must(template.ParseFiles("index.gohtml"))
}

func main() {
	var d dog
	http.ListenAndServe(":8080",d)        //d here is handler so it is handled by (lne1)
}

//what is the difference in using get or post