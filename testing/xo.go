package main

import (
	"../utils"
	"errors"
	"fmt"
	_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	"os"
	//_ "github.com/mattn/go-sqlite3"
	"time"
)

var (
	engine *xorm.Engine
)

type User struct {
	Id            int64
	Email         string
	Password      string
	Nickname      string
	Realname      string
	Avatar        string
	Avatar_min    string
	Avatar_max    string
	Birth         time.Time
	Province      string
	City          string
	Company       string
	Address       string
	Postcode      string
	Mobile        string
	Website       string
	Sex           int64
	Qq            string
	Msn           string
	Weibo         string
	Ctype         int64
	Role          int64
	Created       time.Time
	Hotness       float64
	Hotup         int64
	Hotdown       int64
	Hotscore      int64
	Views         int64
	LastLoginTime time.Time
	LastLoginIp   string
	LoginCount    int64
}

//category,Pid:root
type Category struct {
	Id             int64
	Pid            int64
	Uid            int64
	Ctype          int64
	Title          string
	Content        string
	Attachment     string
	Created        time.Time
	Hotness        float64
	Hotup          int64
	Hotdown        int64
	Hotscore       int64
	Views          int64
	Author         string
	NodeTime       time.Time
	NodeCount      int64
	NodeLastUserId int64
}

//node,Pid:category
type Node struct {
	Id              int64
	Pid             int64
	Uid             int64
	Ctype           int64
	Title           string
	Content         string
	Attachment      string
	Created         time.Time
	Updated         time.Time
	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Hotscore        int64
	Views           int64
	Author          string
	TopicTime       time.Time
	TopicCount      int64
	TopicLastUserId int64
}

//topic,Pid:node
type Topic struct {
	Id              int64
	Cid             int64
	Nid             int64
	Uid             int64
	Ctype           int64
	Title           string
	Content         string
	Attachment      string
	Created         time.Time
	Updated         time.Time
	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Hotscore        int64
	Views           int64
	Author          string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64
	Pid        int64 //Topic id
	Ctype      int64
	Content    string
	Attachment string
	Created    time.Time
	Hotness    float64
	Hotup      int64
	Hotdown    int64
	Hotscore   int64
	Views      int64
	Author     string
	Email      string
	Website    string
}

type File struct {
	Id              int64
	Cid             int64
	Nid             int64
	Uid             int64
	Pid             int64
	Ctype           int64
	Filename        string
	Content         string
	Hash            string
	Location        string
	Url             string
	Size            int64
	Created         time.Time
	Updated         time.Time
	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Hotscore        int64
	Views           int64
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

// k/v infomation
type Kvs struct {
	Id int64
	/*
		Cid int64
		Nid int64
		Tid int64
		Rid int64
	*/
	K string
	V string
}

//Subscribe 订阅之邮件列表
type Subscribe struct {
	Id      int64
	Email   string
	Ctype   int64 //0为接受订阅，1为用户取消订阅。
	Created time.Time
}

func main() {

	//engine.ShowSQL = true

	var tp = Topic{Title: " haha!"}
	PostTopic(tp)
	for i := 0; i < 3; i++ {
		fmt.Println(GetTopic(int64(i)))
	}
}

func init() {
	engine = ConDb()
	defer engine.Close()
	engine.CreateTables(&User{})
	engine.CreateTables(&Category{})
	engine.CreateTables(&Node{})
	engine.CreateTables(&Topic{})
	engine.CreateTables(&Reply{})
	engine.CreateTables(&Kvs{})
	engine.CreateTables(&File{})
	engine.CreateTables(&Subscribe{})
}

func ConDb() *xorm.Engine {

	/*
		engine, _ = xorm.NewEngine("sqlite3", "./test.db")
	*/
	engine, _ = xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")
	return engine
}

func GetTopic(id int64) (tp Topic) {

	engine = ConDb()
	defer engine.Close()
	tp.Id = id

	engine.Get(&tp)

	return tp
}

func PostTopic(tp Topic) (int64, error) {

	engine = ConDb()
	defer engine.Close()
	id, err := engine.Insert(&tp)

	return id, err
}

func PutTopic(tid int64, tp Topic) error {
	/*
				user := User{Name:"xlw"}
		rows, err := engine.Update(&user, &User{Id:1})
		// rows, err := engine.Where("id = ?", 1).Update(&user)
		// or rows, err := engine.Id(1).Update(&user)
	*/
	engine = ConDb()
	defer engine.Close()
	rows, err := engine.Update(&tp, &Topic{Id: tid})
	fmt.Println(rows)
	return err

}

func DelTopic(id int64) error {

	engine = ConDb()
	defer engine.Close()

	topic := new(Topic)
	engine.Id(id).Get(&topic)

	if topic.Attachment != "" {

		if utils.Exist("." + topic.Attachment) {
			if err := os.Remove("." + topic.Attachment); err != nil {
				//return err
				//可以输出错误，但不要反回错误，以免陷入死循环无法删掉
				fmt.Println("DEL TOPIC", id, err)
			}
		}
	}

	//不管实际路径中是否存在文件均删除该数据库记录，以免数据库记录陷入死循环无法删掉
	if topic.Id == id {

		if _, err := engine.Delete(topic); err != nil {

			fmt.Println(err)
		} else {
			return err
		}

	}
	return errors.New("无法删除不存在的TOPIC ID:" + string(id))
}
