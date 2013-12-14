package handler

import (
	"fmt"
	"html/template"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type MainHandler struct {
	lib.BaseHandler
}

func (self *MainHandler) Get() {
	//fmt.Println("im MainHandler")
	self.Data["catpage"] = "home"
	self.TplNames = "main.html"

	ipage, _ := self.GetInt(":page")
	page := int(ipage)

	tab := self.GetString(":tab")

	url := "/"
	if tab == "lastest" {
		url = "/lastest/"
		tab = "id"
		self.Data["tab"] = "lastest"
	} else if tab == "hotness" {
		url = "/hotness/"
		tab = "hotness"
		self.Data["tab"] = "hotness"
	} else {
		url = "/hotness/"
		tab = "hotness"
		self.Data["tab"] = "hotness"
	}

	pagesize := 30
	results_count, err := model.GetTopicsCount(0, pagesize, tab)
	if err != nil {
		return
	}
	pages, page, beginnum, endnum, offset := helper.Pages(int(results_count), page, pagesize)

	if tps, err := model.GetTopics(offset, pagesize, tab); err == nil {
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
			self.Data["pagesbar"] = helper.Pagesbar(url, "", results_count, pages, page, beginnum, endnum, 0)
		}

	} else {
		fmt.Println("首页推荐榜单 数据查询出错", err)
	}

	//侧栏九宫格推荐榜单
	//先行取出最热门的9个节点 然后根据节点获取该节点下最热门的话题
	/*
		if nd, err := model.GetNodes(0, 9, "hotness"); err == nil {
			if len(*nd) > 0 {
				for _, v := range *nd {

					i := 0
					output_start := `<ul class="widgets-popular widgets-similar clx">`
					output := ""
					if tps := model.GetTopicsByNid(v.Id, 0, 1, 0, "hotness"); err == nil {

						if len(*tps) > 0 {
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
					}
					output_end := "</ul>"
					if len(output) > 0 {
						output = output_start + output + output_end
						self.Data["topic_hotness_9_module"] = template.HTML(output)
					} else {
						self.Data["topic_hotness_9_module"] = nil
					}

				}
			}
		} else {
			fmt.Println("节点数据查询出错", err)
		}
	*/

	//侧栏九宫格推荐榜单
	//根据最热的1个分类查找该分类下级9个话题
	/*
		if cats, err := model.GetCategorys(0, 1, "hotness"); err == nil && len(cats) > 0 {
			for _, v := range cats {
				output_start := `<ul class="widgets-popular widgets-similar clx">`
				output := ""
				i := 0
				//根据CID 获取同分类下的最热的9个话题
				if tps := model.GetTopicsByCid(v.Id, 0, 9, 0, "hotness"); len(*tps) > 0 {

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
			}

		}
	*/

	//根据最热的1个节点查找该节点的上级分类下的9个话题
	if nd, err := model.GetNodes(0, 1, "hotness"); err == nil && len(*nd) > 0 {
		for _, v := range *nd {

			output_start := `<ul class="widgets-popular widgets-similar clx">`
			output := ""
			i := 0
			//根据节点的上级PID 获取同分类下的最热话题
			if tps := model.GetTopicsByCid(v.Pid, 0, 9, 0, "hotness"); len(*tps) > 0 {

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

		}
	}
}
