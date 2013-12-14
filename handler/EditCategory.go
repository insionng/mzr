package handler

import (
	"mzr/lib"
	"mzr/model"
	//"strconv"
	"time"
)

type EditCategoryHandler struct {
	lib.RootHandler
}

func (self *EditCategoryHandler) Get() {
	self.TplNames = "edit-category.html"

}

func (self *EditCategoryHandler) Post() {

	cid, _ := self.GetInt("categoryid")

	cat_title := self.GetString("title")
	cat_content := self.GetString("content")

	if cid != 0 && cat_title != "" && cat_content != "" {
		cat := new(model.Category)
		cat.Id = int64(cid)
		cat.Title = cat_title
		cat.Content = cat_content
		cat.Created = time.Now()
		model.PutCategory(cid, cat)
		self.Ctx.Redirect(302, "/category/"+self.GetString("categoryid")+"/")
	} else {
		self.Ctx.Redirect(302, "/")
	}
}
