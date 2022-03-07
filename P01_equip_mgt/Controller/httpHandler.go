package controller

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"

	"equip_mgt.eve/model"
)

var (
	id string
	user string
	name string

	tpl = template.Must(template.ParseFiles("./View/index.html"))
	db = database()
)

func HomeHandler(w http.ResponseWriter,r *http.Request){
	_ = tpl.Execute(w,nil)
}

func EquipHandler(w http.ResponseWriter,r *http.Request){
	vars := mux.Vars(r)
	v := vars["id"]
	st := fmt.Sprintf(`SELECT id,user,name FROM equip.equip where id="%s"`,v)
	row, err :=db.Query(st)
	if err != nil{
		log.Fatal(err)
	}
	for row.Next(){
		err =row.Scan(&id,&user,&name)
		if err != nil{
			panic(err)
		}
	}
		data := model.Equip{
			ID: id,
			User: user,
			Name: name,
		}
	_ = tpl.Execute(w,data)
}

