package main

//https://github.com/GoesToEleven/golang-web-dev/tree/master/021_third-party-serveMux
import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	err := tpl.ExecuteTemplate(w, "index.html",nil)
	handlerError(w,err)
}
func Hello(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
    fmt.Fprintf(w, "hello, %s!\n", p.ByName("name"))
}
func handlerError(w http.ResponseWriter, err error){
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
func main(){
	tpl = template.Must(template.ParseGlob("templates/*"))
	router := httprouter.New()
	router.GET("/", index)//type Handle func(http.ResponseWriter, *http.Request, Params)
	router.GET("/hello/:name", Hello)
	http.ListenAndServe(":8080", router)
}

