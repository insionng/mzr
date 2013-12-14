package main

import (
	"errors"
	"fmt"
	_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"

	"mzr/model"

	"strconv"
	//_ "github.com/mattn/go-sqlite3"

	"os"

	"time"
)

const (
	dbtype = "pgsql"
	//dbtype = "sqlite"
)

var (
	Engine *xorm.Engine
)

type Spider struct {
	Id    int64
	Url   string `xorm:"varchar(300) not null index"`
	Title string
}

type Topic struct {
	Id                int64
	Cid               int64
	Nid               int64
	Uid               int64
	Order             int64
	Ctype             int64
	Title             string
	Content           string `xorm:"text"`
	Attachment        string `xorm:"text"`
	Thumbnails        string //Original remote file
	ThumbnailsLarge   string //200x300
	ThumbnailsMedium  string //200x150
	ThumbnailsSmall   string //70x70
	Tags              string
	Created           time.Time
	Updated           time.Time
	Hotness           float64
	Hotup             int64
	Hotdown           int64
	Hotscore          int64 //Hotup  -	Hotdown
	Hotvote           int64 //Hotup  + 	Hotdown
	Views             int64
	Author            string
	Category          string
	Node              string
	ReplyTime         time.Time
	ReplyCount        int64
	ReplyLastUserId   int64
	ReplyLastUsername string
	ReplyLastNickname string
}

func init() {
	SetEngine()
	CreatTables()
}

func CreatTables() error {
	return Engine.Sync(&Spider{})
}

func ConDb() (*xorm.Engine, error) {
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", "./data/sqlite.db")

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", "user=mysql password=jn!@#9^&* dbname=mysql")

	case dbtype == "pgsql":
		// "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable maxcons=10 persist=true"
		//return xorm.NewEngine("postgres", "host=110.76.39.205 user=postgres password=LhS88root dbname=pgsql sslmode=disable")
		return xorm.NewEngine("postgres", "user=postgres password=LhS88root dbname=mzr sslmode=disable")
	}
	return nil, errors.New("尚未设定数据库连接")
}

func SetEngine() (*xorm.Engine, error) {

	var err error
	Engine, err = ConDb()
	//Engine.Mapper = xorm.SameMapper{}
	//Engine.SetMaxConns(5)
	//Engine.ShowSQL = true
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	Engine.SetDefaultCacher(cacher)

	Engine.ShowDebug = true
	Engine.ShowErr = true
	Engine.ShowSQL = true
	f, err := os.Create("./xsql.log")
	if err != nil {
		println(err.Error())
		panic("OMG!")
	}
	Engine.Logger = f

	return Engine, err
}

func main() {

	for i := 1; i <= 3541; i++ {

		if e := model.DelTopic(int64(i), 1, -1000); e != nil {
			fmt.Println("删除 Topic id:" + strconv.Itoa(i) + "出现错误 " + fmt.Sprintf("%s", e) + "!")

		}
	}

}
