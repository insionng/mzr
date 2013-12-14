package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"time"
)

type DeleteTopicHandler struct {
	lib.AuthHandler
}

func (self *DeleteTopicHandler) Get() {

	flash := beego.NewFlash()
	tid, _ := self.GetInt(":tid")
	uid, _ := self.GetSession("userid").(int64)
	role, _ := self.GetSession("userrole").(int64)
	if tid > 0 {

		if e := model.DelTopic(tid, uid, role); e != nil {

			self.TplNames = "error.html"
			flash.Error("删除 Topic id:" + strconv.Itoa(int(tid)) + "出现错误 " + fmt.Sprintf("%s", e) + "!")
			flash.Store(&self.Controller)

			return
		}
	}
	self.Redirect("/?ver="+strconv.Itoa(int(time.Now().UnixNano())), 302)
}
