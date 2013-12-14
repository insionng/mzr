package handler

import (
	"mzr/lib"
	"strconv"
	"time"
)

type SignoutHandler struct {
	lib.BaseHandler
}

func (self *SignoutHandler) Get() {

	//强制写空 避免被nginx缓存导致无法退出
	self.SetSession("userid", nil)
	self.SetSession("username", nil)
	self.SetSession("userrole", nil)
	self.SetSession("useremail", nil)
	self.SetSession("usercontent", nil)
	//GET状态退出，销毁session
	self.DelSession("userid")
	self.DelSession("username")
	self.DelSession("userrole")
	self.DelSession("useremail")
	self.DelSession("usercontent")

	self.Redirect("/?ver="+strconv.Itoa(int(time.Now().UnixNano())), 302)

}
