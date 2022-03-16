package main

import (
	"net/http"
	"io"
)


func main(){
	http.HandleFunc("/",dog)
	http.Handle("/image/",http.StripPrefix("/image",http.FileServer(http.Dir("./img"))))
	//start svar
	http.ListenAndServe(":8080",nil)
}

func dog(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w,`<img src="/img/gopher.jpg">)
}
