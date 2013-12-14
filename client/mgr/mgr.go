package main

import (
	"errors"
	"fmt"
	_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	"mzr/helper"
	"mzr/model"
	"path"
	//_ "github.com/mattn/go-sqlite3"
	"os"
	"strings"
	"time"
)

const (
	dbtype = "pgsql"
	//dbtype = "sqlite"
)

var (
	Engine *xorm.Engine
)

type MgUrl struct {
	Id      int64
	Url     string `xorm:"varchar(300) not null index"`
	SiteNum int
	IsUsed  bool
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
	return Engine.Sync(&Topic{})
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

func GetSpiderData(offset int, limit int, path string) (*[]MgUrl, error) {
	tls := new([]MgUrl)
	err := Engine.Limit(limit, offset).Desc(path).Find(tls)
	return tls, err
}

func AddTopic(title string, content string, thumbnails string, thumbnailslarge string, thumbnailsmedium string, thumbnailssmall string, cid int64, nid int64, uid int64) (int64, error) {

	id, err := Engine.Insert(&Topic{Cid: cid, Nid: nid, Uid: uid, Title: title, Content: content, Attachment: thumbnails, Thumbnails: thumbnails, ThumbnailsLarge: thumbnailslarge, ThumbnailsMedium: thumbnailsmedium, ThumbnailsSmall: thumbnailssmall, Author: "root",
		Category: "美女", Node: "清新系", Created: time.Now()})
	if thumbnails != "" {
		model.AddImage(thumbnails, id, 1, 1)
	}
	return id, err
}

func main() {

	if imgs, e := GetSpiderData(0, 0, "id"); e == nil {
		for k, v := range *imgs {
			fmt.Println("#", k, ":", v.Url)
			if fpath, err := Download(v.Url); err == nil {
				fmt.Println("fpath:", fpath)
				fmt.Println(helper.Local2url(fpath))
				if thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, e := helper.MakeThumbnails(helper.Local2url(fpath)); e == nil {

					title := "清新系 美女季 2013年度/第" + fmt.Sprint(time.Now().Format("0102-150405")) + "期"
					id, err := AddTopic(title, "<p><img src=\""+helper.Local2url(fpath)+"\" alt=\""+title+"\"/></p>", thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, 1, 1, 1)
					if err != nil {
						fmt.Println("###################发布话题", id, "出错####################", err)
					}
				} else {
					fmt.Println("@@@@@@@@@@@@@@@处理缩略图出错@@@@@@@@@@@@@@@@", err)
				}
			}
		}

	}

}

func Download(url string) (string, error) {

	ext := strings.ToLower(path.Ext(url))
	filename := helper.MD5(time.Now().String()) + ext
	ipath := "./file/" + "download/" //time.Now().Format("03/04/")
	if !helper.Exist(ipath) {
		os.MkdirAll(ipath, 0644)
	}

	path := ipath + filename
	if e := helper.GetFile(url, path, "default", "http://huaban.com/"); e == nil {

		filehash, _ := helper.Filehash(path, nil)
		fname := helper.Encrypt_hash(filehash+"1", nil)
		if ext == "" {
			ext = ".jpg"
		}
		opath := "./file/" + time.Now().Format("03/04/")
		if !helper.Exist(opath) {
			os.MkdirAll(opath, 0644)
		}

		finalpath := opath + fname + ext

		fmt.Println("path:", path)
		fmt.Println("finalpath:", finalpath)
		if e := helper.MoveFile(path, finalpath); e == nil {
			return finalpath, e
		}

	} else {
		return "", errors.New("下载错误!")
	}
	return "", errors.New("Download函数出现错误!")
}
