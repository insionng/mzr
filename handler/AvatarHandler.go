package handler

import (
	"mzr/lib"
	"mzr/model"
)

type AvatarHandler struct {
	lib.BaseHandler
}

func (self *AvatarHandler) Get() {

	username := self.GetString(":username")
	filename := self.GetString(":filename")

	if usr, e := model.GetUserByUsername(username); e == nil && usr != nil && username != "" && filename != "" {

		if filename == "72x72" {
			if usr.Avatar != "" {

				self.Redirect(usr.Avatar, 301)
			} else {
				self.Redirect("/static/img/usr_72x72.png", 302)
			}
		} else {
			self.Redirect("/static/img/usr_72x72.png", 302)
		}
	} else {
		self.Abort("404")
	}
}
