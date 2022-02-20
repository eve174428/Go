package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func main(){
	nf, err := os.Create("syntax.txt")
	if err != nil {
		fmt.Println(err.Error())
	}
	f, err := excelize.OpenFile("r_equip.xlsx")
	if err != nil{
		panic(err)
	}
	rows := f.GetRows("equip")

	var syntax string
	var titles string

	for i, row := range rows {

		if i == 1 {
			for _, col := range row {
				titles += fmt.Sprint(``+ col +`,`)
			}
		}
		if i > 1{
			//INSERT
			var values string
			for _, col := range row {
				values += fmt.Sprintf("'%s',",col)
			}

			syntax += fmt.Sprintf("INSERT INTO equip(%s) VALUES(%s); \n",strings.TrimSuffix(titles,","),strings.TrimSuffix(values,","))
		}
	}
	defer nf.Close()
	io.Copy(nf,strings.NewReader(syntax))
}
