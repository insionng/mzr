package handler

import (
	"fmt"
	"mzr/lib"
	"mzr/model"
)

type HotnessHandler struct {
	lib.BaseHandler
}

func (self *HotnessHandler) Get() {

	self.TplNames = "cat-hotness.html"

	if tps, err := model.GetTopics(0, 30, "hotness"); tps != nil && err == nil {
		self.Data["topics_hotness_30"] = *tps

	} else {
		fmt.Println("hotness headline!err:", err)
	}

	//v, _ := self.RenderString()
	//self.Ctx.WriteString(libs.Cached("default", v, 300))
}
