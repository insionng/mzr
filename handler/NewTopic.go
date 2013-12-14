package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"html/template"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"time"
)

type NewTopicHandler struct {
	lib.AuthHandler
}

func (self *NewTopicHandler) Get() {
	self.TplNames = "new-topic.html"
	//侧栏九宫格推荐榜单
	//先行取出最热门的一个节点 然后根据节点获取该节点下最热门的话题
	if nd, err := model.GetNodes(0, 1, "hotness"); err == nil {
		if len(*nd) == 1 {
			for _, v := range *nd {

				if tps := model.GetTopicsByNid(v.Id, 0, 9, 0, "hotness"); err == nil {

					if len(*tps) > 0 {
						i := 0
						output := `<ul class="widgets-popular widgets-similar clx">`
						for _, v := range *tps {

							i += 1
							if i == 3 {
								output = output + `<li class="similar similar-third">`
								i = 0
							} else {
								output = output + `<li class="similar">`
							}
							output = output + `<a target="_blank" href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `">
												<img src="` + v.ThumbnailsSmall + `" wdith="70" height="70" />				
											</a>
										</li>`
						}
						output = output + `</ul>`
						self.Data["topic_hotness_9_module"] = template.HTML(output)
					}
				} else {
					fmt.Println("推荐榜单(9)数据查询出错", err)
				}
			}
		}
	} else {
		fmt.Println("节点数据查询出错", err)
	}

}

func (self *NewTopicHandler) Post() {
	self.TplNames = "new-topic.html"

	flash := beego.NewFlash()
	nid, _ := self.GetInt("nodeid")

	nd, err := model.GetNode(nid)
	if err != nil || nid == 0 {

		flash.Error("节点不存在,请创建或指定正确的节点!")
		flash.Store(&self.Controller)
		return
	} else {

		cid := nd.Pid
		uid, _ := self.GetSession("userid").(int64)
		sess_username, _ := self.GetSession("username").(string)
		tid_title := self.GetString("title")
		tid_content := self.GetString("content")

		if tid_title != "" && tid_content != "" {

			tp := new(model.Topic)
			tp.Title = tid_title
			tp.Content = tid_content
			tp.Cid = cid
			tp.Nid = nid
			tp.Uid = uid
			tp.Node = nd.Title
			tp.Author = sess_username
			tp.Created = time.Now()

			if s, e := helper.GetBannerThumbnail(tid_content); e == nil {
				tp.Attachment = s
			}

			if thumbnails, thumbnailslarge, thumbnailsmedium, thumbnailssmall, e := helper.GetThumbnails(tid_content); e == nil {
				tp.Thumbnails = thumbnails
				tp.ThumbnailsLarge = thumbnailslarge
				tp.ThumbnailsMedium = thumbnailsmedium
				tp.ThumbnailsSmall = thumbnailssmall
			}

			if cat, err := model.GetCategory(cid); err == nil {
				tp.Category = cat.Title
			}

			nodezmap := &map[string]interface{}{
				"topic_time":         time.Now(),
				"topic_count":        model.GetTopicCountByNid(nid),
				"topic_last_user_id": uid}

			if e := model.UpdateNode(nid, nodezmap); e != nil {
				fmt.Println("NewTopic model.UpdateNode errors:", e)
			}

			if tid, err := model.PostTopic(tp); err == nil {
				model.SetRecordforImageOnPost(tid, uid)
				self.Redirect("/"+strconv.Itoa(int(tid))+"/", 302)
			} else {

				flash.Error(fmt.Sprint(err))
				flash.Store(&self.Controller)
				return
			}
		} else {
			flash.Error("话题标题或内容为空!")
			flash.Store(&self.Controller)
			return
		}
	}
}
