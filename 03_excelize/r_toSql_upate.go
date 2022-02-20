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
	var col_title []string
	var syntax string

	for i, row := range rows {

		if i == 1 {
			for _, col := range row {
				col_title = append(col_title, col)
			}
		}
		if i > 1{
			//upate
			r := row[1:len(row)]
			for i = 0;  i < len(r); i++ {
				syntax += fmt.Sprint(`
				UPDATE equip.equip SET `+ col_title[i+1] +`= '`+ r[i] +`'
				WHERE `+ col_title[0] +`= '`+ row[0] +`';
				`)
			}

		}
	}
	defer nf.Close()
	io.Copy(nf,strings.NewReader(syntax))
}
