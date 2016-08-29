package service

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

type db struct {
	User     string
	Password string
	Host     string
	Name     string
	conn     *dbr.Connection
	sess     *dbr.Session
}

// Connect establishes a connection to the database
func (d *db) Connect() {
	var err error
	d.conn, err = dbr.Open("mysql", d.dsn(), nil)
	if err != nil {
		panic(fmt.Sprintf("service-db: %v", err))
	}
	d.sess = d.conn.NewSession(nil)
	err = d.sess.Ping()
	if err != nil {
		panic(fmt.Sprintf("service-db: %v", err))
	}
}

// Select returns a select query builder
func (d *db) Select(column ...string) *dbr.SelectBuilder {
	return d.sess.Select(column...)
}

// dns generates the connection data source name
func (d *db) dsn() string {
	return fmt.Sprintf("%s:%s@%s/%s?parseTime=true", d.User, d.Password, d.Host, d.Name)
}
