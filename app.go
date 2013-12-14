package main

import (
	"github.com/astaxie/beego"
	"mzr/core"
	"mzr/handler"
	"mzr/helper"
	"runtime"
)

func main() {

	//未登录用户的数据 不用缓存，而是直接 用静态文件即可
	//主要存储静态化后的话题页面
	beego.SetStaticPath("/doc/", "doc/")
	beego.SetStaticPath("/file/", "file/")

	//URL定义规范:必须以/结尾

	//首页
	beego.Router("/", &handler.MainHandler{})
	//首页 ?page
	beego.Router("/page-:page([1-9]\\d*)/", &handler.MainHandler{})

	//首页 hotness类
	beego.Router("/:tab([A-Za-z]+)/", &handler.MainHandler{})
	//http://localhost/lastest/page-2/
	beego.Router("/:tab([A-Za-z]+)/page-:page([1-9]\\d*)/", &handler.MainHandler{})

	beego.Router("/category/:cid:int/", &handler.CategoryHandler{})
	beego.Router("/category/:cid:int/page-:page([1-9]\\d*)/", &handler.CategoryHandler{})

	beego.Router("/category/:tab([A-Za-z]+)/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/", &handler.CategoryHandler{})
	//http://localhost/category/lastest/page-2/
	beego.Router("/category/:tab([A-Za-z]+)/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/page-:page([1-9]\\d*)/", &handler.CategoryHandler{})

	beego.Router("/category/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/", &handler.CategoryHandler{})
	beego.Router("/category/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/page-:page([1-9]\\d*)/", &handler.CategoryHandler{})

	//浏览节点 "/node/:tab([A-Za-z]+)/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/"优先级必须高于"/node/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/"
	beego.Router("/node/:nid:int/", &handler.NodeHandler{})
	beego.Router("/node/:nid:int/page-:page([1-9]\\d*)/", &handler.NodeHandler{})

	beego.Router("/node/:tab([A-Za-z]+)/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/", &handler.NodeHandler{})
	beego.Router("/node/:tab([A-Za-z]+)/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/page-:page([1-9]\\d*)/", &handler.NodeHandler{})

	beego.Router("/node/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/", &handler.NodeHandler{})
	beego.Router("/node/:name([\\x{4e00}-\\x{9fa5}A-Za-z]+)/page-:page([1-9]\\d*)/", &handler.NodeHandler{})

	//详情页面
	beego.Router("/:tid([1-9]\\d*)/", &handler.TopicHandler{})

	//捕抓话题  (获取图片 获取文本 等等..)
	beego.Router("/catch/topic/", &handler.NewTopicHandler{})

	//搜索话题
	beego.Router("/search/", &handler.SearchHandler{})
	//同时支持page和keyword参数
	beego.Router("/search/:keyword([\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]+)/page-:page([1-9]\\d*)/", &handler.SearchHandler{})
	//支持keyword参数 为了兼容所有搜索条件 这里需要用:all
	beego.Router("/search/:keyword([\\x{4e00}-\\x{9fa5}A-Z0-9a-z_-]+)/", &handler.SearchHandler{})

	beego.Router("/timeline/", &handler.TimelineHandler{})
	beego.Router("/user/:username([A-Za-z]+)/", &handler.TimelineHandler{})
	beego.Router("/userid/:userid:int/", &handler.TimelineHandler{})
	//发布时光记录
	beego.Router("/new/timeline/", &handler.TimelineHandler{})
	//删除时光
	beego.Router("/delete/timeline/:lid:int/", &handler.DeleteTimelineHandler{})

	//发现话题 以汇总资讯为方向
	beego.Router("/discover/topic/", &handler.DiscoverHandler{})

	//浏览单图
	beego.Router("/image/:mid:int/", &handler.ImageHandler{})

	//创建分类
	beego.Router("/new/category/", &handler.NewCategoryHandler{})
	//创建话题
	beego.Router("/new/topic/", &handler.NewTopicHandler{})
	//创建节点
	beego.Router("/new/node/", &handler.NewNodeHandler{})

	//创建回应
	beego.Router("/new/reply/:tid:int/", &handler.NewReplyHandler{})
	//删除回应
	beego.Router("/delete/reply/:rid:int/:tid:int/", &handler.DeleteReplyHandler{})
	//beego.Router("/delete/reply/:rid([0-9]+)", &handler.DeleteReplyHandler{})

	//编辑分类
	beego.Router("/edit/category/", &handler.EditCategoryHandler{})
	//编辑节点
	beego.Router("/edit/node/", &handler.EditNodeHandler{})
	//编辑话题
	beego.Router("/edit/topic/:tid:int/", &handler.EditTopicHandler{})

	//删除话题
	beego.Router("/delete/topic/:tid:int/", &handler.DeleteTopicHandler{})

	/*
		beego.Router("/delete/node/:nid([0-9]+)", &handler.NodeDeleteHandler{})

	*/

	//个人设定
	beego.Router("/settings/", &handler.Settings{})
	beego.AutoRouter(&handler.Settings{})

	beego.Router("/avatar/:username([A-Za-z]+)/:filename([A-Za-z]+)/", &handler.AvatarHandler{})

	//登录
	beego.Router("/signin/", &handler.SigninHandler{})
	//退出
	beego.Router("/signout/", &handler.SignoutHandler{})
	//注册
	beego.Router("/signup/", &handler.SignupHandler{})

	//hotness
	beego.Router("/like/:name([A-Za-z]+)/:id:int/", &handler.LikeHandler{})
	beego.Router("/hate/:name([A-Za-z]+)/:id:int/", &handler.HateHandler{})

	//外部URL路由
	beego.Router("/url/", &handler.UrlHandler{})

	//上传文件
	beego.Router("/upload/", &handler.UploaderHandler{})

	//访问次数
	beego.Router("/view/:name([A-Za-z]+)/:id:int/", &handler.ViewHandler{})

	//核心接口 话题接口
	beego.RESTRouter("/core/topic", &core.TopicHandler{})
	//beego.Router("/core/node", &core.NodeHandler{})
	/*
		beego.Router("/root-login", &root.RLoginHandler{})
		beego.Router("/root/account", &root.RAccountHandler{})
	*/

	beego.Router("/debug/pprof/", &handler.PprofHandler{})
	beego.Router("/debug/pprof/:pp([\\w]+)", &handler.PprofHandler{})

	//模板函数
	beego.AddFuncMap("timesince", helper.TimeSince)

	beego.SessionOn = true
	beego.SessionName = "mzr"
	beego.SessionProvider = "file"
	beego.SessionSavePath = "./session"
	beego.AutoRender = true
	beego.CopyRequestBody = true //必须开启,不然core api部分会无法正常工作

	beego.Run()
	runtime.GOMAXPROCS(2)
}
