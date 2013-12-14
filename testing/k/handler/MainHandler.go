package handler

import (
	"github.com/astaxie/beego"
)

type MainHandler struct {
	beego.Controller
}

func (self *MainHandler) Get() {

	self.Ctx.Output.Context.WriteString("okay!!!")
}
