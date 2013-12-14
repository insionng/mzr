
package main

import (
	"net/http"
	"github.com/coocood/qbs"
	_ "github.com/lib/pq"
)

type Process struct {
	Id   int64
	Name string
}

func main(){
	dsn := qbs.DefaultPostgresDataSourceName("qbs_test&quot
	qbs.SetConnectionLimit(10, false)
	qbs.RegisterWithDataSourceName(dsn)
	mg, _ := qbs.GetMigration()
	mg.CreateTableIfNotExists(new(Process))
	http.HandleFunc("/do", do)
	http.ListenAndServe("localhost080", nil)
}

func do(w http.ResponseWriter, r *http.Request){
	q, _ := qbs.GetQbs()
	if q == nil {
		return
	}
	defer q.Close()
	proc := new(Process)
	proc.Id = 100
	proc.Name = "Process"
	q.Where("id=?", 100).Update(proc)
	w.Write([]byte("Process&quot)
} 