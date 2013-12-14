package handler

import (
	"mzr/helper"
	"mzr/lib"
)

type UrlHandler struct {
	lib.BaseHandler
}

func (self *UrlHandler) Get() {

	if helper.IsSpider(self.Ctx.Request.UserAgent()) != true {
		url := self.GetString("localtion")
		if url != "" {
			self.Redirect(url, 302)
		} else {
			self.Abort("401")
		}

	} else {
		self.Abort("401")
	}

}
