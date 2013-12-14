package main

import (
	"errors"
	"fmt"
	//_ "github.com/insionng/pqpooling"
	//_ "github.com/bylevel/pq"
	"github.com/lunny/xorm"
	_ "github.com/mattn/go-sqlite3"
	"time"
	"veryhour/helper"
)

const (
	dbtype = "sqlite"
)

var (
	Engine *xorm.Engine
)

type User struct {
	Id            int64
	Email         string
	Password      string
	Username      string
	Nickname      string
	Realname      string
	Content       string `xorm:"text"`
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
	Updated       time.Time
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
	Order          int64
	Ctype          int64
	Title          string
	Content        string `xorm:"text"`
	Attachment     string `xorm:"text"`
	Created        time.Time
	Updated        time.Time
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
	Order           int64
	Ctype           int64
	Title           string
	Content         string `xorm:"text"`
	Attachment      string `xorm:"text"`
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
	Order           int64
	Ctype           int64
	Title           string
	Content         string `xorm:"text"`
	Attachment      string `xorm:"text"`
	Created         time.Time
	Updated         time.Time
	Hotness         float64
	Hotup           int64
	Hotdown         int64
	Hotscore        int64
	Views           int64
	Author          string
	Category        string
	Node            string
	ReplyTime       time.Time
	ReplyCount      int64
	ReplyLastUserId int64
}

//reply,Pid:topic
type Reply struct {
	Id         int64
	Uid        int64
	Pid        int64 //Topic id
	Order      int64
	Ctype      int64
	Content    string `xorm:"text"`
	Attachment string `xorm:"text"`
	Created    time.Time
	Updated    time.Time
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
	Uid             int64
	Cid             int64
	Nid             int64
	Pid             int64
	Order           int64
	Ctype           int64
	Filename        string
	Content         string `xorm:"text"`
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

type Image struct {
	Id              int64
	Uid             int64
	Cid             int64
	Nid             int64
	Pid             int64
	Order           int64
	Ctype           int64
	Fingerprint     string
	Filename        string
	Content         string `xorm:"text"`
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
type Kv struct {
	Id int64
	K  string `xorm:"text"`
	V  string `xorm:"text"`
}

type Timeline struct {
	Id      int64
	Ctype   int64
	Content string `xorm:"text"`
	Created time.Time
	Updated time.Time
}

func main() {

	SetEngine()
	CreatTables()
	initData()
}

func ConDb() (*xorm.Engine, error) {
	switch {
	case dbtype == "sqlite":
		return xorm.NewEngine("sqlite3", "./sqlite.db")

	case dbtype == "mysql":
		return xorm.NewEngine("mysql", "user=mysql password=jn!@#9^&* dbname=mysql")

	case dbtype == "pgsql":
		// "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable maxcons=10 persist=true"
		return xorm.NewEngine("postgres", "user=postgres password=LhS88root dbname=pgsql sslmode=disable")
		//return xorm.NewEngine("postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable")
	}
	return nil, errors.New("尚未设定数据库连接")
}

func SetEngine() (*xorm.Engine, error) {
	var err error
	Engine, err = ConDb()
	//Engine.SetMaxConns(5)
	//Engine.ShowSQL = true
	return Engine, err
}

func CreatTables() error {

	return Engine.CreateTables(&User{}, &Category{}, &Node{}, &Topic{}, &Reply{}, &Kv{}, &File{}, &Image{}, &Timeline{})
}

func initData() {
	/*
		需求:
		检测管理员账号是否存在
		如若第一次发现不存在后建立管理员账号
		第二次由于条件已经不存在,检测后不会再执行创建账号行为

		此问题在postgreSQL上发现,换为sqlite3也是如此,我相信mysql应该也是如此....好坑的bug..

	*/

	//以下是最开始的做法,这是直接从qbs源码改过来的逻辑,人类的逻辑在xorm终于撞上南墙了!!
	//压根就阻挡不了第2次啊!!

	for i := 0; i < 3; i++ {
		fmt.Println("循环测试1")
		if usr := GetUserByRole(-1000); usr.Role != -1000 {

			fmt.Println("usr:", usr)
			fmt.Println("usr role:", usr.Role)
			fmt.Println("usr id:", usr.Id)
			if id, err := AddUser("root@veryhour.com", "root", "root", "root", helper.Encrypt_hash("rootpass", nil), -1000); err == nil {
				fmt.Println("Default Email:root@veryhour.com ,Username:root ,Password:rootpass Userid:", id)
			} else {
				fmt.Print("create root got errors:", err)
			}

		}
	}

}

func AddUser(email string, username string, nickname string, realname string, password string, role int64) (int64, error) {
	id, err := Engine.Insert(&User{Email: email, Password: password, Username: username, Nickname: nickname, Realname: realname, Role: role, Created: time.Now()})

	return id, err

}

func GetUserByRole(role int64) (*User, error) {
	user := new(User)
	if has, err := Engine.Where("role=?", role).Get(user); has {
		return user, err
	}
	return nil, err
}

func GetUserByRole33(role int64) (user User) {
	Engine.Where("role=?", role).Get(&user)
	return user
}

func GetUserByRole2(role int64) (*User, error) {

	user := &User{Role: role}
	has, err := Engine.Get(user)
	if has {
		return user, err
	} else {

		return nil, err
	}
}
