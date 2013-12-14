package handler

import (
	"fmt"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"time"
)

type ViewHandler struct {
	lib.BaseHandler
}

func (self *ViewHandler) Get() {
	name := self.GetString(":name")
	id, _ := self.GetInt(":id")

	if name != "" && id > 0 {
		if name == "topic" {

			if tp, err := model.GetTopic(id); tp != nil && err == nil {
				tp.Views = tp.Views + 1
				tp.Hotup = tp.Hotup + 1
				tp.Hotscore = helper.Hotness_Score(tp.Hotup, tp.Hotdown)
				tp.Hotness = helper.Hotness(tp.Hotup, tp.Hotdown, time.Now())

				if row, e := model.PutTopic(id, tp); e != nil {
					fmt.Println("ViewHandler更新话题ID", id, "访问次数数据错误,row:", row, e)
					self.Abort("500")
				} else {
					self.Ctx.Output.Context.WriteString(strconv.Itoa(int(tp.Views)))
				}
			}
		} else {
			self.Abort("501")
		}

	} else {
		self.Abort("501")
	}

}
