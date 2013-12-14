package handler

import (
	"fmt"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"strings"
)

type TimelineHandler struct {
	lib.BaseHandler
}

func (self *TimelineHandler) Get() {

	self.TplNames = "timeline.html"
	username := self.GetString(":username")
	uid, _ := self.GetInt(":userid")

	switch {
	case username != "":

		if usr, err := model.GetUserByUsername(username); usr != nil && err == nil {
			if tls, err := model.GetTimelines(0, 0, "created", usr.Id); err == nil {
				self.Data["timelines_username"] = username
				self.Data["timelines"] = *tls
				return
			}

		} else {
			self.Redirect("/timeline/", 302)
		}

	case uid != 0:
		if usr, err := model.GetUser(uid); usr != nil && err == nil {
			if tls, err := model.GetTimelines(0, 0, "hotness", usr.Id); err == nil {
				self.Data["timelines"] = *tls
				return
			}
		} else {
			self.Redirect("/timeline/", 302)
		}
	case uid == 0 && username == "": //首页
		if tls, err := model.GetTimelines(0, 0, "hotness", 0); err == nil {
			self.Data["timelines"] = *tls
			return
		} else {
			self.Redirect("/timeline/", 302)
		}
	default:
		self.Redirect("/timeline/", 302)
	}
}

func (self *TimelineHandler) Post() {

	self.TplNames = "timeline.html"
	sess_userid, _ := self.GetSession("userid").(int64)
	sess_username, _ := self.GetSession("username").(string)

	tl := self.GetString("timeline")

	//不等于0,即是注册用户或管理层
	if sess_userid != 0 {
		if tl != "" {
			//获取当前用户信息
			if usr, err := model.GetUser(sess_userid); usr != nil && err == nil {

				//前一条是记录发送者自己的timeline  所以这里的接受者是自己  记录完自己的timeline后  再处理自己timeline的内容里的@通知
				if lid, err := model.AddTimeline("", tl, 0, 0, sess_userid, usr.Username, usr.Content); err != nil {
					fmt.Println("#", lid, ":", err)
				} else {
					//如果 内容tl中有@通知 则处理以下事件
					if users := helper.AtUsers(tl); len(users) > 0 {
						//todo := []string{}
						//被通知列表 k是uid v是username
						todolist := map[int64]string{}
						for _, v := range users {
							//判断被通知之用户名是否真实存在
							if u, e := model.GetUserByUsername(v); e == nil && u != nil {
								//存在的则加入待操作列
								//todo = append(todo, v)
								todolist[u.Id] = u.Username
								//替换被通知用户的用户名带上用户主页链接
								tl = strings.Replace(tl, "@"+v,
									"<a href='/user/"+u.Username+"/' title='"+u.Nickname+"' target='_blank'><span>@</span><span>"+u.Username+"</span></a>", -1)

								//发送通知内容到被通知用户的 时光记录 注意这里的uid不再是sess_userid 而是u.Id
								model.AddTimeline(usr.Username+"提到了你~",
									tl+"[<a href='/user/"+usr.Username+"/#timeline-"+strconv.Itoa(int(lid))+"'>"+usr.Username+"</a>]", 0, 0, u.Id, usr.Username, usr.Content)

							}

						}
						//如果有@通知操作 则重新替换一次发送者已存档的内容
						if len(todolist) > 0 {
							model.SetTimelineContentByRid(lid, tl)
						}
					}

					//处理@link
					if atpagez, _ := helper.AtPages(tl); len(atpagez) > 0 {
						tid := int64(0)
						if tid, tl, err = model.AtLinksPostImagesOnTopic(tl); err == nil {

							model.SetTimelineContentByRid(lid, tl+" <a href='/topic/"+strconv.Itoa(int(tid))+"/' target='_blank'>[#美图合辑("+strconv.Itoa(int(tid))+")#]</a>")
						}
					}
					self.Redirect("/user/"+sess_username+"/#timeline-"+strconv.Itoa(int(lid)), 302)
					return
				}
			}
			self.Redirect("/user/"+sess_username+"/", 302)
		} else {
			self.Redirect("/timeline/", 302)
		}
	} else {
		self.Redirect("/timeline/", 302)

	}

}
