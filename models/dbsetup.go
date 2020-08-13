package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"madhyam/logging"
	"os"

	_ "github.com/go-sql-driver/mysql"
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
		logging.WriteLog(logging.ERROR, "Couldn't find database credentials file.")
	}

	bytearr, err := ioutil.ReadAll(file)
	if err != nil {
		logging.WriteLog(logging.ERROR, "Error occoured while reading the credential file.")
	}
	dbc := DBCreds{}

	json.Unmarshal(bytearr, &dbc)
	host := dbc.Mysqldb.Host
	port := dbc.Mysqldb.Port
	database := dbc.Mysqldb.Database
	username := dbc.Mysqldb.Username
	password := dbc.Mysqldb.Password

	connStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)

	return connStr

}

func ConDB() *sql.DB {
	db, err := sql.Open("mysql", ConnectionStr())
	if err != nil {
		fmt.Println("INTERNAL SERVER ERROR: Couldn't connect to the database!")
	}

	if err := db.Ping(); err != nil {
		fmt.Println("INTERNAL SERVER ERROR: The connection has been closed!")
		fmt.Println(err)
		return nil
	}
	return db
}
