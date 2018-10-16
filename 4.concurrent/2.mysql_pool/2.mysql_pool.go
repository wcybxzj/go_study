//数据库连接池测试
package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

var db *sql.DB

/*
SetMaxOpenConns用于设置最大打开的连接数，默认值为0表示不限制。
SetMaxIdleConns用于设置闲置的连接数。

设置最大的连接数，可以避免并发太高导致连接mysql出现too many connections的错误。
设置闲置的连接数则当开启的一个连接使用完成后可以放在池里等候下一次使用。
*/
func init() {
	db, _ = sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/zuji_order?charset=utf8")
	db.SetMaxOpenConns(2000)
	db.SetMaxIdleConns(1000)
	db.Ping()
}

func startHttpServer() {
	http.HandleFunc("/pool", pool)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func pool(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT id FROM order_goods_instalment  limit 10")
	defer rows.Close()
	checkErr(err)

	columns, _ := rows.Columns()
	scanArgs := make([]interface{}, len(columns))
	values := make([]interface{}, len(columns))
	for j := range values {
		scanArgs[j] = &values[j]
	}

	record := make(map[string]string)
	for rows.Next() {
		//将行数据保存到record字典
		err = rows.Scan(scanArgs...)
		for i, col := range values {
			if col != nil {
				record[columns[i]] = string(col.([]byte))
			}
		}
	}

	//fmt.Println(record)
	fmt.Fprintln(w, "finish")
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

//apt-get install apache2-utils

//watch -n1 "mysql -uroot -proot -e  'show processlist'"

//ab -c 100 -n 1000 'http://localhost:9090/pool'
func main() {
	startHttpServer()
}
