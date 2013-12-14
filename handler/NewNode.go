package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"mzr/lib"
	"mzr/model"
	"time"
)

type NewNodeHandler struct {
	lib.AuthHandler
}

func (self *NewNodeHandler) Get() {
	self.TplNames = "new-node.html"
}

func (self *NewNodeHandler) Post() {
	self.TplNames = "new-node.html"
	flash := beego.NewFlash()

	cid, _ := self.GetInt("category")
	uid, _ := self.GetSession("userid").(int64)
	nid_title := self.GetString("title")
	nid_content := self.GetString("content")

	if nid_title != "" && nid_content != "" && cid != 0 {

		if nid, err := model.AddNode(nid_title, nid_content, cid, uid); err != nil {
			flash.Error(fmt.Sprint(err))
			flash.Store(&self.Controller)
			return
		} else {

			if nd, err := model.GetNode(nid); err != nil {
				flash.Error(fmt.Sprint(err))
				flash.Store(&self.Controller)
				return
			} else {

				catmap := &map[string]interface{}{
					"node_time":         time.Now(),
					"node_count":        model.GetNodeCountByPid(cid),
					"node_last_user_id": uid}

				if e := model.UpdateCategory(cid, catmap); e != nil {
					fmt.Println("NewNode model.UpdateCategory errors:", e)
				}

				nd.Order = nd.Id * 10
				if row, err := model.PutNode(nid, nd); row != 0 && err != nil {
					flash.Error(fmt.Sprint(err))
					flash.Store(&self.Controller)
					return
				} else {
					self.Redirect("/", 302)
				}
			}
		}
	} else {
		flash.Error("分类不存在或节点标题、节点内容为空,请纠正错误!")
		flash.Store(&self.Controller)
		return
	}
}
