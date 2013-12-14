package helper

import (
	"github.com/astaxie/beego"
	"veryhour/plugin/goconfig"
)

func init() {
	theme := GetTheme()

	beego.SetStaticPath("/static", theme+"/static")
	beego.SetViewsPath(theme + "/views")

}

func GetTheme() string {
	var theme string

	if conf, err := goconfig.LoadConfigFile("./conf/config.conf"); err != nil {

		//如果不存在配置文件，则设为默认主题路径
		theme = "default"
	} else { //如果存在配置文件

		// 主题设置读取错误 即section不存在 或 字段为空 则重置主题为默认主题并保存到配置文件
		if theme, err = conf.GetValue("theme", "name"); err != nil {
			conf.SetValue("theme", "name", "default")
			goconfig.SaveConfigFile(conf, "./conf/config.conf")
			theme = "default"
		}
	}
	return "theme/" + theme
}
