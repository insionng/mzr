package main

import (
	"fmt"
	"github.com/coocood/qbs"
	_ "github.com/lib/pq"
	"net/http"
	"time"
)

type Process struct {
	Id     int64
	Name   string
	CtTime time.Time `qbs:"created"`
}

func main() {
	dsn := new(qbs.DataSourceName)
	dsn.Password = "jn!@#$%^&*"
	dsn.Username = "postgres"
	dsn.DbName = "pgsql"
	dsn.Dialect = qbs.NewPostgres()
	qbs.RegisterWithDataSourceName(dsn)
	dsn.Append("sslmode", "disable")
	qbs.SetConnectionLimit(10, false)
	mg, err := qbs.GetMigration()
	if err != nil {
		fmt.Println(err)
	}
	mg.Log = true
	mg.CreateTableIfNotExists(new(Process))
	http.HandleFunc("/", do)
	http.ListenAndServe(":80", nil)
}

func do(w http.ResponseWriter, r *http.Request) {
	q, _ := qbs.GetQbs()
	if q == nil {
		return
	}
	defer q.Close()
	proc := new(Process)
	proc.Id = 100
	proc.Name = "Process"
	q.Save(proc)
	w.Write([]byte("Process"))
}
