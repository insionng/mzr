package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	//"encoding/json"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type UploaderHandler struct {
	lib.BaseHandler
}

func (self *UploaderHandler) Post() {

	flash := beego.NewFlash()
	targetFolder := "/file/"
	self.TplNames = "editor-tinymce-ajax-result.html"
	file, handler, e := self.GetFile("userfile")

	uid := int64(0)
	if self.GetSession("userid") != nil {
		uid = self.GetSession("userid").(int64)
	} else {
		flash.Error("UploaderHandler获取UID错误0!")
		flash.Store(&self.Controller)

		self.Data["result"] = "UploaderHandler获取UID错误0!"
		self.Data["resultcode"] = "failed"
		return
	}

	if e != nil {
		fmt.Println("UploaderHandler获取文件错误1!")
		flash.Error("UploaderHandler获取文件错误1!")
		flash.Store(&self.Controller)

		self.TplNames = "editor-tinymce-ajax-result.html"
		self.Data["result"] = " "
		self.Data["resultcode"] = "failed"
	} else {

		if handler != nil {

			ext := strings.ToLower(path.Ext(handler.Filename))
			filename := helper.MD5(time.Now().String()) + ext

			ipath := targetFolder + time.Now().Format("03/04/")
			//ipath := targetFolder + helper.FixedpathByNumber(2, 2)
			os.MkdirAll("."+ipath, 0644)
			path := ipath + filename
			f, err := os.OpenFile("."+path, os.O_WRONLY|os.O_CREATE, 0644)

			if err != nil {

				fmt.Println("UploaderHandler获取文件错误2!")
				flash.Error("UploaderHandler获取文件错误2!")
				flash.Store(&self.Controller)

				self.TplNames = "editor-tinymce-ajax-result.html"
				self.Data["result"] = " "
				self.Data["resultcode"] = "failed"
			} else {
				io.Copy(f, file)
				defer file.Close()
				defer f.Close()
				input_file := "." + path
				output_file := "." + path
				output_size := "696"
				output_align := "center"
				background := "white"
				newpath := ""
				//所有上传的图片都会被缩略处理
				if err := helper.Thumbnail("resize", input_file, output_file, output_size, output_align, background); err != nil {

					fmt.Println("UploaderHandler生成缩略图出错:", err)
					flash.Error(fmt.Sprint(err))
					flash.Store(&self.Controller)

					if e := os.Remove(helper.Url2local(path)); e != nil {
						fmt.Println("UploaderHandler清除残余文件出错:", e)
					}
					self.TplNames = "editor-tinymce-ajax-result.html"
					self.Data["result"] = err
					self.Data["resultcode"] = "failed"
				} else {
					f.Close() //手动关闭  不然下面会导致重命名文件出错~
					watermark_file := helper.Url2local(helper.GetTheme()) + "/static/mzr/img/watermark.png"

					if e := helper.Watermark(watermark_file, input_file, output_file, "SouthEast"); e == nil {
						//所有文件以该加密方式哈希生成文件名  从而实现针对到用户个体的文件权限识别
						filehash, _ := helper.Filehash(helper.Url2local(path), nil)

						fname := helper.Encrypt_hash(filehash+strconv.Itoa(int(uid)), nil)

						newpath = ipath + fname + ext

						if err := os.Rename(helper.Url2local(path), helper.Url2local(newpath)); err != nil {
							fmt.Println("重命名文件出错:", err)
						}

						//文件权限校验 通过说明文件上传转换过程中没发生错误
						//首先读取被操作文件的hash值 和 用户请求中的文件hash值  以及 用户当前id的string类型  进行验证

						if fhashed, _ := helper.Filehash(helper.Url2local(newpath), nil); helper.Validate_hash(fname, fhashed+strconv.Itoa(int(uid))) {

							//用户上传图片的记录
							//ctype为0表示没上传文件
							//收到的图片存储都设置ctype为 -1  证明用户上传了文件,但尚未正式使用
							//当用户edit话题或new话题,在进行posting的时候,检查image表,如存在同样文件,则顺手修改ctype为1 表示该文件正在使用
							//并修改相关tid uid等等信息进image表 留待以后或许有用~
							if _, err := model.AddImage(helper.Url2local(newpath), 0, -1, uid); err != nil {
								fmt.Print("model.AddImage:", err)
							}

							self.TplNames = "editor-tinymce-ajax-result.html"
							self.Data["result"] = "file_uploaded"
							self.Data["resultcode"] = "ok"
							self.Data["file_name"] = newpath
						} else {

							fmt.Println("UploaderHandler校验图片不正确!")
							flash.Error("UploaderHandler校验图片不正确!")
							flash.Store(&self.Controller)

							self.TplNames = "editor-tinymce-ajax-result.html"
							self.Data["result"] = " "
							self.Data["resultcode"] = "failed"

							if e := os.Remove(helper.Url2local(newpath)); e != nil {
								fmt.Println("UploaderHandler清除错误文件", newpath, "出错:", e)
							}
						}

						//hash, _ := utils.Filehash(output_file)
						//fileInfo, err := os.Stat(output_file)
						//var fsize int64 = 0
						//if err == nil {
						//	fsize = fileInfo.Size() / 1024
						//}
					} else {
						fmt.Println("UploaderHandler 图片添加水印失败!")
						flash.Error("UploaderHandler 图片添加水印失败!")
						flash.Store(&self.Controller)

						self.TplNames = "editor-tinymce-ajax-result.html"
						self.Data["result"] = " "
						self.Data["resultcode"] = "failed"

						if e := os.Remove(helper.Url2local(newpath)); e != nil {
							fmt.Println("UploaderHandler清除错误水印遗留文件", newpath, "出错:", e)
						}
					}

				}

			}

		} else {

			fmt.Println("UploaderHandler获取文件错误3!")
			flash.Error("UploaderHandler获取文件错误3!")
			flash.Store(&self.Controller)

			self.TplNames = "editor-tinymce-ajax-result.html"
			self.Data["result"] = " "
			self.Data["resultcode"] = "failed"
		}
	}

}
