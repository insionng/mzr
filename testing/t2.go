package main

import (
	"fmt"
	"github.com/coocood/qbs"
	//_ "github.com/lib/pq" //当某时间字段表现为0001-01-01 07:36:42+07:36:42形式的时候 会读不出数据
	_ "github.com/bylevel/pq"
	"time"
)

type Topic struct {
	Id              int64
	Cid             int64 `qbs:"index"`
	Nid             int64 `qbs:"index"`
	Uid             int64 `qbs:"index"`
	Ctype           int64
	Title           string
	Content         string
	Attachment      string
	Created         time.Time
	Updated         time.Time
	Hotness         float64 `qbs:"index"`
	Hotup           int64   `qbs:"index"`
	Hotdown         int64   `qbs:"index"`
	Views           int64   `qbs:"index"`
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

func main() {
	qbs.Register("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable", "pgsql", qbs.NewPostgres())
	m, _ := qbs.GetMigration()
	//m.Log = true
	//m.DropTable(new(Topic))
	m.CreateTableIfNotExists(new(Topic))

	q, _ := qbs.GetQbs()
	//q.Log = true
	i := new(Topic)
	i.Content = "i am test!!!!!!"
	q.Save(i)

	nana, _ := FindNaById(q, 1)
	fmt.Println(nana)
}

func FindNaById(q *qbs.Qbs, id int64) (*Topic, error) {
	na := new(Topic)
	//na.Id = id
	err := q.WhereEqual("id", id).Find(na)
	return na, err
}
