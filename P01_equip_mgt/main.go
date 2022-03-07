package main

import (
	"errors"
	"log"
	"net/http"
	"os"

	"equip_mgt.eve/route"

	"github.com/ichtrojan/thoth"
	"github.com/joho/godotenv"
)
func main(){
	lg, _ := thoth.Init("log")
	if err := godotenv.Load(); err != nil{
		logMsg(lg,".env file not found.")
	}

	port, exist := os.LookupEnv("PORT")
	if !exist{
		logMsg(lg,"PORT not set in .evn")
	}
	log.Fatal(http.ListenAndServe(":"+port, route.Init()))
}

func logMsg(lg thoth.Config, msg string){
	lg.Log(errors.New(msg))
	log.Fatal(msg)
}
