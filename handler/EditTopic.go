package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type EditTopicHandler struct {
	lib.AuthHandler
}

func (self *EditTopicHandler) Get() {
	self.TplNames = "edit-topic.html"
	flash := beego.NewFlash()

	tid, _ := self.GetInt(":tid")

	if tid_handler, err := model.GetTopic(tid); err == nil {

		self.Data["topic"] = *tid_handler
		self.Data["inode"], _ = model.GetNode(tid_handler.Nid)

	} else {

		flash.Error(fmt.Sprint(err))
		flash.Store(&self.Controller)
		return
	}
}

func (self *EditTopicHandler) Post() {
	self.TplNames = "edit-topic.html"
	flash := beego.NewFlash()

	tid, _ := self.GetInt(":tid")
	nid, _ := self.GetInt("nodeid")

	if nd, err := model.GetNode(nid); nd != nil && err == nil {

		uid, _ := self.GetSession("userid").(int64)
		tid_title := self.GetString("title")
		tid_content := self.GetString("content")

		if tid_title != "" && tid_content != "" {

			if tp, err := model.GetTopic(tid); tp != nil && err == nil {

				tp.Title = tid_title
				tp.Uid = uid

				//删去用户没再使用的图片
				helper.DelLostImages(tp.Content, tid_content)
				tp.Content = tid_content

				if s, e := helper.GetBannerThumbnail(tid_content); e == nil {
					tp.Attachment = s
				}

				if cat, err := model.GetCategory(nd.Pid); err == nil {
					tp.Category = cat.Title
				}

				if row, err := model.PutTopic(tid, tp); row == 1 && err == nil {
					model.SetRecordforImageOnEdit(tid, uid)
					self.Redirect("/"+strconv.Itoa(int(tid))+"/", 302)
				} else {

					flash.Error("更新话题出现错误:", fmt.Sprint(err))
					flash.Store(&self.Controller)
					return
				}
			} else {

				flash.Error("无法获取根本不存在的话题!")
				flash.Store(&self.Controller)
				return
			}
		} else {

			flash.Error("话题标题或内容为空!")
			flash.Store(&self.Controller)
			return
		}
	} else {
		flash.Error(fmt.Sprint(err))
		flash.Store(&self.Controller)
		return
	}
}
