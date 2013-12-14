package main

import (
	"errors"
	"fmt"
	_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	"mzr/helper"
	"mzr/model"
	"runtime"
	//_ "github.com/mattn/go-sqlite3"
	"os"
)

const (
	dbtype = "pgsql"
	//dbtype = "sqlite"
)

var (
	Engine *xorm.Engine
)

func init() {
	SetEngine()
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

	runtime.GOMAXPROCS(1)
	//遍历图集
	if mz, er := model.GetImagesByCtype(1); er == nil && len(*mz) > 0 {
		for k, v := range *mz {

			if thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, e := helper.MakeThumbnails(helper.Local2url(v.Location)); e == nil {
				fmt.Println("#", k, ":", thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, e)
			} else {
				fmt.Println("@@@@@@@@@@@@@@@处理缩略图出错@@@@@@@@@@@@@@@@", er)
			}
		}
	}
	//覆写缩略图

}
