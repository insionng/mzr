package handler

import (
	"fmt"
	"html/template"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type NodeHandler struct {
	lib.BaseHandler
}

func (self *NodeHandler) Get() {
	//fmt.Println("im NodeHandler")

	self.TplNames = "node.html"

	tab := self.GetString(":tab")
	ipage, _ := self.GetInt(":page")
	page := int(ipage)
	nid, _ := self.GetInt(":nid")
	ndname := self.GetString(":name")
	if nid > 0 {

		if n, e := model.GetNode(nid); e == nil && n != nil {
			ndname = n.Title
		}
	}

	url := "/node/"
	if tab == "lastest" {
		url = "/node/lastest/" + ndname + "/"
		tab = "id"
		self.Data["tab"] = "lastest"
	} else if tab == "hotness" {
		url = "/node/hotness/" + ndname + "/"
		tab = "hotness"
		self.Data["tab"] = "hotness"
	} else {
		url = "/node/hotness/" + ndname + "/"
		tab = "hotness"
		self.Data["tab"] = "hotness"
	}

	pagesize := 30

	if ndname != "" {
		//检验节点是否存在
		if nd, err := model.GetNodeByTitle(ndname); err == nil && nd != nil {
			//更新节点统计信息
			nd.Views = nd.Views + 1
			self.Data["innode"] = *nd
			self.Data["catpage"] = nd.Title
			model.PutNode(nd.Id, nd)

			limit := 25
			rcs := len(*model.GetTopicsByNid(nd.Id, 0, 0, 0, "hotness"))
			pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, int(page), limit)

			//通过节点的名字获取下级话题
			if tps, err := model.GetTopicsByNode(ndname, offset, pagesize, tab); err == nil {
				results_count := len(*tps)
				if results_count > 0 {
					i := 1
					output := ""
					for _, v := range *tps {

						i += i
						if i == 3 {
							output = output + `<div id="pin-` + strconv.Itoa(int(v.Id)) + `" class="pin pin3">`
							i = 0
						} else {

							output = output + `<div id="pin-` + strconv.Itoa(int(v.Id)) + `" class="pin">`
						}
						output = output + `<div class="pin-coat">
									<a href="/` + strconv.Itoa(int(v.Id)) + `/" class='imageLink image loading' target='_blank'>
										<img src='/static/mzr/img/dot.png' original='` + v.ThumbnailsLarge + `' width='200' height='150' alt='` + v.Title + ` ` + v.Created.String() + ` ` + v.Node + ` ` + v.Category + `' oriheight='300' />
										<span class='bg'>` + v.Title + `</span>
									</a>
									<div class="pin-data clx">
										<span class="timer">
											<em></em>
											<span>` + helper.TimeSince(v.Created) + `</span>
										</span>
										<a href="/` + strconv.Itoa(int(v.Id)) + `/" class="viewsButton" title="阅读` + v.Title + `" target="_blank">
											<em></em>
											<span>` + strconv.Itoa(int(v.Views)) + ` views</span>
										</a>
									</div>
								</div>
							</div>`
					}
					self.Data["topics"] = output

				}

			} else {
				fmt.Println("节点推荐榜单 数据查询出错", err)
			}

			self.Data["pagesbar"] = helper.Pagesbar(url, "", rcs, pages, pageout, beginnum, endnum, 0)

			//侧栏九宫格推荐榜单
			output_start := `<ul class="widgets-popular widgets-similar clx">`
			output := ""
			i := 0
			//根据节点的上级PID 获取同分类下的最热话题
			if tps := model.GetTopicsByCid(nd.Pid, 0, 9, 0, "hotness"); len(*tps) > 0 {

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
			}
			output_end := "</ul>"

			if len(output) > 0 {
				output = output_start + output + output_end
				self.Data["topic_hotness_9_module"] = template.HTML(output)
			} else {
				self.Data["topic_hotness_9_module"] = nil
			}
		} else {
			self.Redirect("/", 302)
		}

	} else {
		self.Redirect("/", 302)
	}

}
