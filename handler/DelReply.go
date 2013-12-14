package handler

import (
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type DeleteReplyHandler struct {
	lib.AuthHandler
}

func (self *DeleteReplyHandler) Get() {
	rid, _ := self.GetInt(":rid")
	tid, _ := self.GetInt(":tid")
	if rid > 0 && tid > 0 {
		model.DelReply(rid)
		self.Redirect("/topic/"+strconv.Itoa(int(tid))+"/", 302)
	} else {
		self.Redirect("/", 302)
	}
}
