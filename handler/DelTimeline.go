package handler

import (
	"fmt"
	"mzr/lib"
	"mzr/model"
)

type DeleteTimelineHandler struct {
	lib.AuthHandler
}

func (self *DeleteTimelineHandler) Get() {
	lid, _ := self.GetInt(":lid")
	sess_username := self.GetSession("username").(string)

	if lid > 0 && sess_username != "" {
		if e := model.DelTimeline(lid); e != nil {
			fmt.Println("DeleteTimelineHandler errors:", e)
		}
		self.Redirect("/user/"+sess_username+"/", 302)
	} else {
		self.Redirect("/timeline/", 302)
	}
}
