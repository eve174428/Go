package route

import (
	"equip_mgt.eve/controller"
	"github.com/gorilla/mux"
)

func Init() *mux.Router{
	r := mux.NewRouter()
	r.HandleFunc("/",controller.HomeHandler)
	r.HandleFunc("/{id}",controller.EquipHandler)
	return r
}
