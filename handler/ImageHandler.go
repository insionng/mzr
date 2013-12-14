package handler

import (
	"mzr/lib"
	"mzr/model"
)

type ImageHandler struct {
	lib.BaseHandler
}

func (self *ImageHandler) Get() {

	self.TplNames = "image.html"

	mid, _ := self.GetInt(":mid")

	if mid > 0 {
		if im, err := model.GetImage(mid); im != nil && err == nil {

			self.Data["image"] = *im
			self.Data["replys"] = *model.GetReplysByPid(mid, 1, 0, 0, "id")

			if ims, err := model.GetImagesByCtypeWidthCid(1, im.Cid); *ims != nil && mid != 0 && err == nil {
				for i, v := range *ims {

					if v.Id == mid {
						prev := i - 1
						next := i + 1

						for i, v := range *ims {
							if prev == i {
								self.Data["previd"] = v.Id
								self.Data["prev"] = "上张"
							}
							if next == i {
								self.Data["nexmid"] = v.Id
								self.Data["next"] = "下张"
							}
						}
					}
				}
			}

			if ims, err := model.GetImagesByCtypeWidthNid(1, im.Nid); *ims != nil && err == nil {
				self.Data["images_bynid"] = *ims
			}
		} else {
			self.Redirect("/", 302)
		}
	} else {
		self.Redirect("/", 302)
	}
}
