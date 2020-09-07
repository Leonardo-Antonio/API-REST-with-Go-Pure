package storage

import (
	"database/sql"
	"log"
	"sync"

	//!!
	_ "github.com/go-sql-driver/mysql"
)

var (
	once     sync.Once
	instance *Mysql
)

// Mysql -> connection on database mysql
type Mysql struct {
	db *sql.DB
}

// NewMysql -> method contructor of the class Mysql
func NewMysql() *Mysql {
	once.Do(func() {
		db, err := sql.Open("mysql", "leo:chester@tcp(localhost:3306)/BD_GO")
		if err != nil {
			log.Fatalf("Error al intentar acceder al bd -> %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("Error en el ping -> %v", err)
		}
		instance = createSingleInstance(db)
	})
	return instance
}

func createSingleInstance(db *sql.DB) *Mysql {
	return &Mysql{db}
}

// Pool .
func (mysql *Mysql) Pool() *sql.DB {
	return mysql.db
}
