package controller

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ichtrojan/thoth"
	_ "github.com/joho/godotenv/autoload"
)

func database() *sql.DB{

	user := getEvn("DB_USER")
	pwd := getEvn("DB_PASS")
	host := getEvn("DB_HOST")
	dataSrc := fmt.Sprintf("%s:%s@(%s:3306)/equip",user,pwd,host)
	fmt.Println(dataSrc)
	db, err := sql.Open("mysql",dataSrc)
	if err != nil{
		log.Fatal("connect err:",err)
	}
	return db
}
func getEvn(key string) string{
	lgr, _:= thoth.Init("log")
	value, exist := os.LookupEnv(key)
	if !exist{
		msg := fmt.Sprintf("%s not found.",key)
		lgr.Log(errors.New(msg))
		log.Fatal(msg)
	}
	return value
}
