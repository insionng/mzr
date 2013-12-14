package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"mzr/lib"
	"mzr/model"
)

type NewCategoryHandler struct {
	lib.AuthHandler
}

func (self *NewCategoryHandler) Get() {
	self.TplNames = "new-category.html"
}

func (self *NewCategoryHandler) Post() {
	self.TplNames = "new-category.html"
	flash := beego.NewFlash()

	t := self.GetString("title")
	c := self.GetString("content")

	if t != "" && c != "" {
		if id, err := model.AddCategory(t, c); err != nil {
			flash.Error("AddCategory Id", id, ":", fmt.Sprint(err))
			flash.Store(&self.Controller)
		} else {
			//分类创建成功后跳转到新创建的分类id去
			self.Redirect("/", 302)
		}
	} else {
		flash.Error("分类标题或内容为空,请纠正错误!")
		flash.Store(&self.Controller)
		return
	}

}
