package main

import (
	"code.google.com/p/mahonia"
	"errors"
	"fmt"
	_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	"mzr/helper"
	"mzr/model"
	"path"
	"strconv"
	//_ "github.com/mattn/go-sqlite3"
	"github.com/PuerkitoBio/goquery"
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

	Urls []string
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
		//return xorm.NewEngine("postgres", "host=110.76.39.205 user=postgres password=LhS88root dbname=mzr sslmode=disable")
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

func GetSpiderData(offset int, limit int, path string) (*[]Spider, error) {
	tls := new([]Spider)
	err := Engine.Limit(limit, offset).Desc(path).Find(tls)
	return tls, err
}

func AddSpiderData(url string, title string) (int64, error) {
	if has, err := Engine.Where("url=?", url).Get(new(Spider)); !has || err != nil {
		id, err := Engine.Insert(&Spider{Url: url, Title: title})
		return id, err
	}
	return -1, nil

}

func UpdateTopic(id int64, title string, content string, thumbnails string, thumbnailslarge string, thumbnailsmedium string, thumbnailssmall string, cid int64, nid int64, uid int64) (int64, error) {

	id, err := Engine.Insert(&Topic{Id: id, Cid: cid, Nid: nid, Uid: uid, Title: title, Content: content, Attachment: thumbnails, Thumbnails: thumbnails, ThumbnailsLarge: thumbnailslarge, ThumbnailsMedium: thumbnailsmedium, ThumbnailsSmall: thumbnailssmall, Author: "root",
		Category: "美女", Node: "性感系", Created: time.Now()})
	if thumbnails != "" {
		ctype := int64(1)
		model.AddImage(thumbnails, id, ctype, uid)
	}
	return id, err
}

func AddTopic(title string, content string, thumbnails string, thumbnailslarge string, thumbnailsmedium string, thumbnailssmall string, cid int64, nid int64, uid int64) (int64, error) {

	id, err := Engine.Insert(&Topic{Cid: cid, Nid: nid, Uid: uid, Title: title, Content: content, Attachment: thumbnails, Thumbnails: thumbnails, ThumbnailsLarge: thumbnailslarge, ThumbnailsMedium: thumbnailsmedium, ThumbnailsSmall: thumbnailssmall, Author: "root",
		Category: "美女", Node: "性感系", Created: time.Now()})
	if thumbnails != "" {
		ctype := int64(1)
		model.AddImage(thumbnails, id, ctype, uid)
	}
	return id, err
}

func Download(url string) (string, error) {

	ext := strings.ToLower(path.Ext(url))
	filename := helper.MD5(time.Now().String()) + ext
	ipath := "./file/" + "download/" //time.Now().Format("03/04/")
	if !helper.Exist(ipath) {
		os.MkdirAll(ipath, 0644)
	}

	path := ipath + filename
	if e := helper.GetFile(url, path, "default", "http://www.fengyun5.com/"); e == nil {

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

func SelfPage(cururl string) {

	x, _ := goquery.NewDocument(cururl)
	title := x.Find("title").Text()

	if url, b := x.Find("#content a img").Attr("src"); b == true {

		if alt, b := x.Find("#content a img").Attr("alt"); b == true {
			title = alt
		}

		enc := mahonia.NewDecoder("gbk")
		if _, e := AddSpiderData(url, enc.ConvertString(title)); e != nil {
			fmt.Println(e)
		}
	}

}

func main() {

	urls := []string{}
	//循环所有页面查找所有图片的网页链接
	for i := 617; i <= 779; i++ {
		purl := "http://www.fengyun5.com/Tpimage/" + strconv.Itoa(i) + "/"
		x, _ := goquery.NewDocument(purl)
		x.Find(".zzz a").Each(func(idx int, s *goquery.Selection) {
			v, b := s.Attr("href")
			if b == true {
				urls = append(urls, purl+v)
			}
		})
	}

	//遍历所有图片网址 提取图片URL后保存到数据库
	for k, v := range urls {

		fmt.Println("<url #[", k, "]# url>")
		SelfPage(v) //单独处理网页
	}

	//读取图片集合并下载
	if imgs, e := GetSpiderData(0, 0, "id"); e == nil {
		j := int64(0)
		for k, v := range *imgs {
			fmt.Println("#>", k, ":", v.Url)
			if fpath, err := Download(v.Url); err == nil {

				if helper.Exist(fpath) {
					if thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, e := helper.MakeThumbnails(helper.Local2url(fpath)); thumbnails != "" && thumbnailslarge != "" && thumbnailsmedium != "" && thumbnailssmall != "" && e == nil {
						j += 1
						title := "性感系 " + v.Title + " 美女季/第" + fmt.Sprint(time.Now().Format("0102-150405")) + "期"
						cid := int64(1)
						nid := int64(4) //2
						uid := int64(1)
						role := int64(-1000)

						if j <= 1 {
							fmt.Println("###post topic A ,file:", fpath)
							model.DelTopic(25099, uid, role)
							id, err := UpdateTopic(j, title, "<p><img src=\""+helper.Local2url(fpath)+"\" alt=\""+v.Title+"\"/></p>", thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, cid, nid, uid)
							if err != nil {
								fmt.Println("###################发布话题", id, "出错####################", err)
							}
						} else {
							fmt.Println(":::Post topic B ,file:", fpath)
							id, err := AddTopic(title, "<p><img src=\""+helper.Local2url(fpath)+"\" alt=\""+v.Title+"\"/></p>", thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, cid, nid, uid)
							if err != nil {
								fmt.Println("###################发布话题", id, "出错####################", err)
							}
						}

					} else {
						fmt.Println("@@@@@@@@@@@@@@@处理缩略图出错@@@@@@@@@@@@@@@@", err)
						os.Remove(fpath)
						os.Remove(helper.Url2local(helper.SetSuffix(fpath, "_large.jpg")))
						os.Remove(helper.Url2local(helper.SetSuffix(fpath, "_medium.jpg")))
						os.Remove(helper.Url2local(helper.SetSuffix(fpath, "_small.jpg")))

					}
				}

			}
		}

	}

}
