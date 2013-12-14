package handler

import (
	"mzr/lib"
	"mzr/model"
	"strconv"
	"time"
)

type EditNodeHandler struct {
	lib.RootHandler
}

func (self *EditNodeHandler) Get() {

	self.TplNames = "edit-node.html"

}

func (self *EditNodeHandler) Post() {

	cid, _ := self.GetInt("categoryid")
	nid, _ := self.GetInt("nodeid")

	nd_title := self.GetString("title")
	nd_content := self.GetString("content")

	if cid != 0 && nid != 0 && nd_title != "" && nd_content != "" {

		nd := new(model.Node)
		nd.Id = int64(nid)
		nd.Pid = int64(cid)
		nd.Title = nd_title
		nd.Content = nd_content
		nd.Created = time.Now()

		if _, err := model.PutNode(nd.Id, nd); err == nil {

			self.Redirect("/node/"+strconv.Itoa(int(nid))+"/", 302)
		} else {
			self.Redirect("/", 302)
		}

	} else {
		self.Ctx.Redirect(302, "/")
	}
}
