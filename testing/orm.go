package main

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	//_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// 设置数据库路径
	_DB_NAME = "./data.db"
	// 设置数据库名称
	_SQLITE3_DRIVER = "sqlite3"
)

type User struct {
	Id      int `orm:"auto"` // 设置为auto主键
	Name    string
	Profile *Profile `orm:"rel(one)"` // OneToOne relation
}

type Profile struct {
	Id   int `orm:"auto"`
	Age  int16
	User *User `orm:"reverse(one)"` // 设置反向关系(可选)
}

func init() {
	// 数据库别名
	name := "default"
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User), new(Profile))

	//orm.RegisterDriver("postgres", orm.DR_Postgres)
	//orm.RegisterDataBase("default", "postgres", "user=postgres password=jn!@#$%^&* dbname=pgsql sslmode=disable", 30)
	orm.RegisterDataBase(name, _SQLITE3_DRIVER, _DB_NAME, 10)

	// drop table 后再建表
	force := true

	// 打印执行过程
	verbose := false

	// 遇到错误立即返回
	err := orm.RunSyncdb(name, force, verbose)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	o := orm.NewOrm()

	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	profile := &Profile{}

	profile.Age = 30

	user := &User{}
	user.Profile = profile
	user.Name = "insion"
	i, _ := o.Insert(profile)
	d, _ := o.Insert(user)
	fmt.Println(i)
	fmt.Println(d)
}
