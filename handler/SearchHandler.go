package handler

import (
	"fmt"
	"html/template"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type SearchHandler struct {
	lib.BaseHandler
}

func (self *SearchHandler) Get() {
	self.TplNames = "search.html"
	keyword := self.GetString(":keyword")
	if keyword == "" {
		keyword = self.GetString("keyword")
	}

	ipage, _ := self.GetInt(":page")
	if ipage <= 0 {
		ipage, _ = self.GetInt("page")
	}
	page := int(ipage)

	limit := 9 //每页显示数目

	//如果已经登录登录
	sess_username, _ := self.GetSession("username").(string)
	if sess_username != "" {
		limit = 30
	}

	if keyword != "" {
		if rc, err := model.SearchTopic(keyword, 0, 0, "id"); err == nil {

			rcs := len(*rc)
			pages, pageout, beginnum, endnum, offset := helper.Pages(rcs, page, limit)

			if st, err := model.SearchTopic(keyword, offset, limit, "hotness"); err == nil {
				results_count := len(*st)
				if results_count > 0 {
					i := 1
					output := ""
					for _, v := range *st {

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

			}

			if k := self.GetString("keyword"); k != "" {
				self.Data["search_keyword"] = k
			} else {
				self.Data["search_keyword"] = keyword
			}

			self.Data["pagebar"] = helper.Pagesbar("/search/", keyword, rcs, pages, pageout, beginnum, endnum, 0)
		} else {
			fmt.Println("SearchTopic errors:", err)
			return
		}

		//侧栏九宫格推荐榜单
		//根据用户的关键词推荐
		nds, ndserr := model.SearchNode(keyword, 0, 9, "hotness")
		cats, catserr := model.SearchCategory(keyword, 0, 9, "hotness")
		//如果在节点找到关键词
		if (ndserr == catserr) && (ndserr == nil) {

			output_start := `<ul class="widgets-popular widgets-similar clx">`
			output := ""
			i := 0
			if len(*nds) >= len(*cats) && len(*nds) > 0 {
				for _, v := range *nds {

					if tps := model.GetTopicsByNid(v.Id, 0, 1, 0, "hotness"); len(*tps) > 0 {

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
			} else if len(*cats) > len(*nds) && len(*cats) > 0 {
				for _, v := range *cats {

					if tps := model.GetTopicsByCid(v.Id, 0, 1, 0, "hotness"); len(*tps) > 0 {

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
			}
			output_end := "</ul>"
			if len(output) > 0 {
				output = output_start + output + output_end
				self.Data["topic_hotness_9_module"] = template.HTML(output)
			} else {
				self.Data["topic_hotness_9_module"] = nil
			}
		}
		//侧栏九宫格推荐榜单
		//先行取出最热门的9个节点 然后根据节点获取该节点下最热门的话题
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
	} else {
		self.Redirect("/", 302)
	}

}
