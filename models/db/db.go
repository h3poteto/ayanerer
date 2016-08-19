package db

import (
	"database/sql"
	"io/ioutil"
	"os"
	"path/filepath"

	// require mysql but is not referred
	_ "github.com/go-sql-driver/mysql"
	"gopkg.in/yaml.v2"
)

// Database is struct for database connection
type Database struct {
}

// Init is prepare database connection
func (u *Database) Init() *sql.DB {
	env := os.Getenv("ECHOENV")
	root := os.Getenv("ECHOROOT")
	path := filepath.Join(root, "db/dbconf.yml")
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	m := make(map[interface{}]interface{})
	err = yaml.Unmarshal(buf, &m)
	if err != nil {
		panic(err)
	}
	username := m[env].(map[interface{}]interface{})["user"].(string)
	password := m[env].(map[interface{}]interface{})["password"].(string)
	database := m[env].(map[interface{}]interface{})["name"].(string)
	host := m[env].(map[interface{}]interface{})["host"].(string)
	username = os.ExpandEnv(username)
	password = os.ExpandEnv(password)
	database = os.ExpandEnv(database)
	host = os.ExpandEnv(host)
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+host+":3306)/"+database+"?charset=utf8")
	if err != nil {
		panic(err)
	}
	return db
}
