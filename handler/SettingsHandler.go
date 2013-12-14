package handler

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"mzr/helper"
	"mzr/lib"
	"mzr/model"
	"os"
	"path"
	"strconv"
	"strings"
	"time"
)

type Settings struct {
	lib.AuthHandler
}

func (self *Settings) Get() {

	self.TplNames = "settings.html"
	beego.ReadFromRequest(&self.Controller)

	sess_userid, _ := self.GetSession("userid").(int64)

	if usr, err := model.GetUser(sess_userid); usr != nil && err == nil {
		self.Data["usr"] = *usr
		return
	}
}

func (self *Settings) Profile() {

	flash := beego.NewFlash()
	self.TplNames = "settings.html"

	sess_userid, _ := self.GetSession("userid").(int64)

	username := self.GetString("username")
	email := self.GetString("email")

	nickname := self.GetString("nickname")
	realname := self.GetString("realname")

	content := self.GetString("content")
	birth := self.GetString("birth")
	fmt.Println(birth)
	province := self.GetString("province")
	city := self.GetString("city")
	company := self.GetString("company")
	address := self.GetString("address")

	postcode := self.GetString("postcode")
	mobile := self.GetString("mobile")
	website := self.GetString("website")
	sex, _ := self.GetInt("sex")
	qq := self.GetString("qq")
	msn := self.GetString("msn")
	weibo := self.GetString("weibo")

	if username == "" {
		flash.Error("用户名不能为空!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if email == "" {
		flash.Error("Email是你的主账号,和主要联系方式,不能留空~")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if content == "" {
		flash.Error("为了让别人更了解你,请务必填写你的个人签名~")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if helper.CheckUsername(username) == false {
		flash.Error("用户名包含非法字符,或不合符规格(限4~30个字符)~")
		flash.Store(&self.Controller)

		return

	}

	if helper.CheckEmail(email) == false {
		flash.Error("Email格式不合符规格~")
		flash.Store(&self.Controller)

		return

	}

	if usrinfo, err := model.GetUser(sess_userid); usrinfo != nil && err == nil {

		usrinfo.Username = username
		usrinfo.Email = email

		usrinfo.Nickname = nickname
		usrinfo.Realname = realname
		usrinfo.Content = content
		usrinfo.Birth = time.Now()
		usrinfo.Province = province
		usrinfo.City = city
		usrinfo.Company = company
		usrinfo.Address = address
		usrinfo.Postcode = postcode
		usrinfo.Mobile = mobile
		usrinfo.Website = website
		usrinfo.Sex = sex
		usrinfo.Qq = qq
		usrinfo.Msn = msn
		usrinfo.Weibo = weibo

		if _, err := model.PutUser(usrinfo.Id, usrinfo); err == nil {

			//更新session
			self.SetSession("userid", usrinfo.Id)
			self.SetSession("username", usrinfo.Username)
			self.SetSession("userrole", usrinfo.Role)
			self.SetSession("useremail", usrinfo.Email)
			self.SetSession("usercontent", usrinfo.Content)

			flash.Notice("设置个人信息成功~")
		} else {
			flash.Error("设置个人信息失败~")
		}
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	} else {

		flash.Error("该账号不存在~")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

}

func (self *Settings) Password() {

	flash := beego.NewFlash()
	self.TplNames = "settings.html"

	sess_userid, _ := self.GetSession("userid").(int64)
	curpass := self.GetString("curpass")
	newpassword := self.GetString("password")
	newrepassword := self.GetString("repassword")

	if curpass == "" {
		flash.Error("当前密码不能为空!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if newpassword == "" {
		flash.Error("设置密码不能为空!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if newrepassword == "" {
		flash.Error("重验设置密码不能为空!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if newpassword != newrepassword {
		flash.Error("两次密码不一致!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

	if helper.CheckPassword(curpass) == false {
		flash.Error("当前密码含有非法字符或当前密码过短(至少4~30位密码)!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return

	}

	if helper.CheckPassword(newpassword) == false {
		flash.Error("设置密码含有非法字符或设置密码过短(至少4~30位密码)!")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return

	}

	if usrinfo, err := model.GetUser(sess_userid); usrinfo != nil && err == nil {

		if helper.Validate_hash(usrinfo.Password, curpass) {
			usrinfo.Password = helper.Encrypt_hash(newpassword, nil)

			if _, err := model.PutUser(usrinfo.Id, usrinfo); err == nil {
				flash.Notice("设置密码成功~")
			} else {
				flash.Error("设置密码失败~")
			}
			flash.Store(&self.Controller)

			self.Redirect("/settings/", 302)
			return
		} else {

			flash.Error("密码无法通过校验~")
			flash.Store(&self.Controller)

			self.Redirect("/settings/", 302)
			return
		}
	} else {

		flash.Error("该账号不存在~")
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
	}

}

func (self *Settings) Avatar() {

	flash := beego.NewFlash()
	self.TplNames = "settings.html"

	targetFolder := "/file/"
	file, handler, e := self.GetFile("avatar")
	uid := self.GetSession("userid").(int64)

	if e != nil {
		flash.Error("SettingsHandler获取文件错误1," + fmt.Sprint(e))
		flash.Store(&self.Controller)

		self.Redirect("/settings/", 302)
		return
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
				flash.Error("SettingsHandler获取文件错误2!")
				flash.Store(&self.Controller)

				self.Redirect("/settings/", 302)
				return
			} else {
				io.Copy(f, file)
				defer file.Close()
				defer f.Close()
				input_file := "." + path
				output_file := "." + path
				output_size := "72x72"
				output_align := "center"
				background := "#f0f0f0"
				newpath := ""
				//所有上传的图片都会被缩略处理
				if err := helper.Thumbnail("crop", input_file, output_file, output_size, output_align, background); err != nil {

					flash.Error(fmt.Sprint(err))
					flash.Store(&self.Controller)

					if e := os.Remove(path); e != nil {
						fmt.Println("SettingsHandler清除残余文件出错:", e)
					}

					self.Redirect("/settings/", 302)
					return
				} else {
					f.Close() //手动关闭  不然下面会导致重命名文件出错~

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

						//收到的头像图片存储都设置ctype为 10 与其他图片类型区分开
						if _, err := model.AddImage(helper.Url2local(newpath), 0, 10, uid); err != nil {
							fmt.Print("model.AddImage:", err)
						}

						usr, _ := model.GetUser(uid)
						if usr.Avatar != "" {
							os.Remove(helper.Url2local(usr.Avatar))
						}
						if usr.AvatarLarge != "" {
							os.Remove(helper.Url2local(usr.AvatarLarge))
						}
						if usr.AvatarMedium != "" {
							os.Remove(helper.Url2local(usr.AvatarMedium))
						}
						if usr.AvatarSmall != "" {
							os.Remove(helper.Url2local(usr.AvatarSmall))
						}
						usr.Avatar = newpath
						usr.AvatarLarge = newpath
						usr.AvatarMedium = newpath
						usr.AvatarSmall = newpath
						model.PutUser(uid, usr)

						//hash, _ := utils.Filehash(output_file)
						//fileInfo, err := os.Stat(output_file)
						//var fsize int64 = 0
						//if err == nil {
						//	fsize = fileInfo.Size() / 1024
						//}

						flash.Notice("成功设置头像为:", handler.Filename)
						flash.Store(&self.Controller)
						self.Redirect("/settings/", 302)
						return
					} else {

						flash.Error("SettingsHandler图片添加水印失败!")
						flash.Store(&self.Controller)

						if e := os.Remove(helper.Url2local(newpath)); e != nil {
							fmt.Println("SettingsHandler清除错误水印遗留文件", newpath, "出错:", e)
						}

						self.Redirect("/settings/", 302)
						return
					}

				}

			}

		} else {

			flash.Error("SettingsHandler获取文件错误3!")
			flash.Store(&self.Controller)

			self.Redirect("/settings/", 302)
			return
		}
	}
}
