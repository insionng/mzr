package handler

import (
	"fmt"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
	"strings"
)

type NewReplyHandler struct {
	lib.BaseHandler
}

func (self *NewReplyHandler) Post() {
	tid, _ := self.GetInt(":tid")
	sess_userid, _ := self.GetSession("userid").(int64)

	author := self.GetString("author")
	email := self.GetString("email")
	website := self.GetString("website")
	rc := self.GetString("comment")

	//不等于0,即是注册用户或管理层 此时把ctype设置为1 主要是为了区分游客
	if sess_userid != 0 {
		if tid > 0 && rc != "" {

			if usr, err := model.GetUser(sess_userid); err == nil {
				//为安全计,先行保存回应,顺手获得rid,在后面顺手再更新替换@通知的链接
				if rid, err := model.AddReply(tid, sess_userid, 1, rc, usr.Username, usr.Content, usr.Email, usr.Website); err != nil {
					fmt.Println("#", rid, ":", err)
				} else {

					//如果回应内容中有@通知 则处理以下事件
					if users := helper.AtUsers(rc); len(users) > 0 {
						if tp, err := model.GetTopic(tid); err == nil {
							todo := []string{}
							for _, v := range users {
								//判断被通知之用户名是否真实存在
								if u, e := model.GetUserByUsername(v); e == nil && u != nil {
									//存在的则加入待操作列
									todo = append(todo, v)
									//替换被通知用户的用户名带上用户主页链接
									rc = strings.Replace(rc, "@"+v,
										"<a href='/user/"+u.Username+"/' title='"+u.Nickname+"' target='_blank'><span>@</span><span>"+u.Username+"</span></a>", -1)

									//发送通知内容到用户的 时间线
									model.AddTimeline(usr.Username+"在「"+tp.Title+"」的回应里提到了你~",
										rc+"[<a href='/topic/"+self.GetString(":tid")+"/#reply-"+strconv.Itoa(int(rid))+"'>"+tp.Title+"</a>]",
										tp.Cid, tp.Nid, u.Id, usr.Username, usr.Content)

								}

							}
							if len(todo) > 0 {
								model.SetReplyContentByRid(rid, rc)
							}

						}
					}

					self.Redirect("/topic/"+self.GetString(":tid")+"/#reply-"+strconv.Itoa(int(rid)), 302)
					return
				}
			}
			self.Redirect("/topic/"+self.GetString(":tid")+"/", 302)
		} else if tid > 0 {
			self.Redirect("/topic/"+self.GetString(":tid")+"/", 302)
		} else {
			self.Redirect("/", 302)
		}
	} else { //游客回应 此时把ctype设置为-1   游客不开放@通知功能
		if author != "" && email != "" && tid > 0 && rc != "" {
			if rid, err := model.AddReply(tid, sess_userid, -1, rc, author, "", email, website); err != nil {
				fmt.Println("#", rid, ":", err)
				self.Redirect("/topic/"+self.GetString(":tid")+"/", 302)
			} else {
				self.Redirect("/topic/"+self.GetString(":tid")+"/#reply-"+strconv.Itoa(int(rid)), 302)
			}
		} else if tid > 0 {
			self.Redirect("/topic/"+self.GetString(":tid")+"/", 302)
		} else {
			self.Redirect("/", 302)
		}

	}

}
