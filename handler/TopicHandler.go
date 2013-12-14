package handler

import (
	"fmt"
	"html/template"
	"mzr/lib"
	"mzr/model"
	"strconv"
)

type TopicHandler struct {
	lib.BaseHandler
}

func (self *TopicHandler) Get() {
	//fmt.Println("im TopicHandler")
	self.TplNames = "topic.html"

	tid, _ := self.GetInt(":tid")

	if tid > 0 {

		if tp, err := model.GetTopic(tid); tp != nil && err == nil {

			self.Data["article"] = *tp
			self.Data["replys"] = *model.GetReplysByPid(tid, 0, 0, 0, "id")

			//性能瓶颈 亟待优化!!!
			/*
				if tps := model.GetTopicsByCid(tp.Cid, 0, 0, 0, "asc"); *tps != nil && tid != 0 {
					if len(*tps) > 0 {
						self.Data["topics_cat"] = *tps
						//self.Data["topics_cat"] = model.GetTopicsByHotnessCategory(3, 3)

					}

					for i, v := range *tps {

						if v.Id == tid {
							//两侧的翻页按钮参数 初始化 s
							prev := i - 1
							next := i + 1
							//两侧的翻页按钮参数 初始化 e

							//话题内容部位 页码  初始化 s
							ipagesbar := `<div class="link_pages">`
							h := 5
							ipagesbar_start := i - h
							ipagesbar_end := i + h
							j := 0
							//话题内容部位 页码  初始化 e
							for i, v := range *tps {
								//两侧的翻页按钮 s
								if prev == i {
									self.Data["previd"] = v.Id
									self.Data["prev"] = v.Title
								}
								if next == i {
									self.Data["nextid"] = v.Id
									self.Data["next"] = v.Title
								}
								//两侧的翻页按钮 e

								//话题内容部位 页码 s
								if ipagesbar_start == i {
									ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span><</span></a>`
								}
								if i > ipagesbar_start && i < ipagesbar_end {
									j += 1
									if v.Id == tid { // current

										ipagesbar = ipagesbar + `<span>` + strconv.Itoa(int(v.Id)) + `</span>`

									} else { //loop

										ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span>` + strconv.Itoa(int(v.Id)) + `</span></a>`

									}
									if j > (2 * h) {
										break
									}
								}
								if ipagesbar_end == i {
									ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span>></span></a>`
								}
								//话题内容部位 页码 e
							}
							self.Data["ipagesbar"] = template.HTML(ipagesbar + "</div>")
						}
					}
				}
			*/

			if tps := model.GetTopicsByCidOnBetween(tp.Cid, tid-5, tid+5, 0, 11, 0, "asc"); tps != nil && tid != 0 && len(tps) > 0 {

				for i, v := range tps {

					if v.Id == tid {
						//两侧的翻页按钮参数 初始化 s
						prev := i - 1
						next := i + 1
						//两侧的翻页按钮参数 初始化 e

						//话题内容部位 页码  初始化 s
						ipagesbar := `<div class="link_pages">`
						h := 5
						ipagesbar_start := i - h
						ipagesbar_end := i + h
						j := 0
						//话题内容部位 页码  初始化 e
						for i, v := range tps {
							//两侧的翻页按钮 s
							if prev == i {
								self.Data["previd"] = v.Id
								self.Data["prev"] = v.Title
							}
							if next == i {
								self.Data["nextid"] = v.Id
								self.Data["next"] = v.Title
							}
							//两侧的翻页按钮 e

							//话题内容部位 页码 s
							if ipagesbar_start == i {
								ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span><</span></a>`
							}
							if i > ipagesbar_start && i < ipagesbar_end {
								j += 1
								if v.Id == tid { // current

									ipagesbar = ipagesbar + `<span>` + strconv.Itoa(int(v.Id)) + `</span>`

								} else { //loop

									ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span>` + strconv.Itoa(int(v.Id)) + `</span></a>`

								}
								if j > (2 * h) {
									break
								}
							}
							if ipagesbar_end == i {
								ipagesbar = ipagesbar + `<a href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `"><span>></span></a>`
							}
							//话题内容部位 页码 e
						}
						self.Data["ipagesbar"] = template.HTML(ipagesbar + "</div>")
					}
				}
				// (tps []*Topic)

			}
			//侧栏你可能喜欢 推荐同节点下的最热话题
			if tps := model.GetTopicsByNid(tp.Nid, 0, 6, 0, "hotness"); *tps != nil {

				if len(*tps) > 0 {
					i := 0
					output := `<ul class="widgets-similar clx">`
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
					self.Data["topic_sidebar_hotness_6_module"] = template.HTML(output)
				}

			}

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

			//底部六格推荐
			//推荐同一作者的最热话题
			if tps := model.GetTopicsByUid(tp.Uid, 0, 6, 0, "hotness"); len(*tps) > 0 {
				i := 0
				output := `<ul class="widgets-similar clx">`
				for _, v := range *tps {

					i += 1
					if i == 3 {
						output = output + `<li class="likesimilar likesimilar-3">`
						i = 0
					} else {
						output = output + `<li class="likesimilar">`
					}
					output = output + `<a target="_blank" href="/` + strconv.Itoa(int(v.Id)) + `/" title="` + v.Title + `" alt="` + v.Title + `" class="likeimglink">
												<img src="` + v.ThumbnailsMedium + `" wdith="150" height="150" />
												<span class="bg">` + v.Title + `</span>			
											</a>
										</li>`
				}
				output = output + `</ul>`
				self.Data["topic_hotness_6_module"] = template.HTML(output)
			} else {
				fmt.Println("六格推荐查询出错:", err)
			}

		} else {
			self.Redirect("/", 302)
		}
	} else {
		self.Redirect("/", 302)
	}

}
