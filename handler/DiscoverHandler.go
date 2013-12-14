package handler

import (
	"fmt"
	"mzr/lib"
	"mzr/model"
)

type DiscoverHandler struct {
	lib.BaseHandler
}

func (self *DiscoverHandler) Get() {
	self.TplNames = "discover-topic.html"

	if tps, err := model.GetTopics(0, 3, "hotness"); tps != nil && err == nil {
		self.Data["topics_hotness_headline"] = *tps

	} else {
		fmt.Println("hotness headline!err:", err)
	}

	//v, _ := self.RenderString()
	//self.Ctx.WriteString(libs.Cached("default", v, 300))
}
