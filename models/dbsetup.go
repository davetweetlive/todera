package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type DBCreds struct {
	Mysqldb `json:"mysqldb"`
}

type Mysqldb struct {
	Name      string `json:"name"`
	Connector string `json:"connector"`
	Host      string `json:"host"`
	Port      int    `json:"port"`
	Database  string `json:"database"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

func ConnectionStr() string {
	file, err := os.Open("config/dbcred.json")
	if err != nil {
		fmt.Println("Couldn't open the database credential file")
	}

	bytearr, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Couldn't parse the file into byte slice")
	}
	dbc := DBCreds{}

	json.Unmarshal(bytearr, &dbc)
	host := dbc.Mysqldb.Host
	port := dbc.Mysqldb.Port
	database := dbc.Mysqldb.Database
	username := dbc.Mysqldb.Username
	password := dbc.Mysqldb.Password

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	// fmt.Println(connStr)
	return connStr
	// return ""
}

func ConDB() *sql.DB {
	db, err := sql.Open("mysql", ConnectionStr())
	if err != nil {
		fmt.Println("ERROR:")
	}
	return db
	// stm, err := db.Prepare(`CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));`)
	// if err != nil {
	// 	fmt.Println("ERRPR!")
	// }
	// fmt.Println(stm)
}
