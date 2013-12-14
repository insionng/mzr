package handler

import (
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"time"
)

type HateHandler struct {
	lib.BaseHandler
}

func (self *HateHandler) Get() {

	if helper.IsSpider(self.Ctx.Request.UserAgent()) != true {
		name := self.GetString(":name")
		id, _ := self.GetInt(":id")

		if name == "topic" {

			tp, _ := model.GetTopic(id)
			tp.Hotdown = tp.Hotdown + 1
			tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
			tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

			model.PutTopic(id, tp)
			self.Ctx.WriteString(strconv.Itoa(int(tp.Hotscore)))

		} else if name == "timeline" {

			tl, _ := model.GetTimeline(id)
			tl.Hotdown = tl.Hotdown + 1
			tl.Hotscore = helper.Hotness_Score(tl.Hotup, tl.Hotdown)
			tl.Hotness = helper.Hotness(tl.Hotup, tl.Hotdown, time.Now())

			model.PutTimeline(id, tl)
			self.Ctx.WriteString(strconv.Itoa(int(tl.Hotscore)))

		} else if name == "node" {

			nd, _ := model.GetNode(id)
			nd.Hotdown = nd.Hotdown + 1
			nd.Hotscore = helper.Hotness_Score(nd.Hotup, nd.Hotdown)
			nd.Hotness = helper.Hotness(nd.Hotup, nd.Hotdown, time.Now())

			model.PutNode(id, nd)

			self.Ctx.WriteString("node hated")
		} else {
			self.Abort("401")
		}

	} else {
		self.Abort("401")
	}

}
