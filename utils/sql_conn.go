package utils

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func SelectStatus(orderNumber string) string {
	db, err := sql.Open("mysql", "123123:123123@tcp(localhost:3306)/weidiango_coretrans")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.SetMaxIdleConns(10)
	db.SetConnMaxLifetime(1000)
	//smt,_ := db.Prepare(orderNumber)
	//rows,_:=smt.Query(orderNumber)
	//if rows == nil {
	//	fmt.Println("sql result is [NULL]")
	//}
	//var status string
	//for rows.Next()  {
	//	rows.Scan(&status)
	//}
	//defer rows.Close()
	res, _ := db.Exec(orderNumber)
	fmt.Print(res)
	return "1"
}
