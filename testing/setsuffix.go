package main

import (
	"../utils"
	"fmt"
)

func main() {
	fmt.Println("#1", utils.SetSuffix("/file/2013/09/14/1404994819612896a8e8316c21d70f6b.jpg", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("#2", utils.SetSuffix("/file/2013/09/14/1404994819612896a8e8316c21d70f6b", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("#3", utils.SetSuffix("/file/2013/09/14/1404994819612896a8e8316c21d70f6b.", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("#4", utils.SetSuffix("./file/2013/09/14/1404994819612896a8e8316c21d70f6b", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("#5", utils.SetSuffix("./file/2013/09/14/1404994819612896a8e8316c21d70f6b.", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("#6", utils.SetSuffix("./file/2013/09/14/1404994819612896a8e8316c21d70f6b.png", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("WEB#1", utils.SetSuffix("http://bgimg1.meimei22.com/recpic/2013/27.jpg", "_banner.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("WEB#2", utils.SetSuffix("http://t12.baidu.com/it/u=118735160,777859196&fm=120", "_banner.jpg"))

	fmt.Println("LC#1", utils.IsLocal("./file/2013/09/14/1404994819612896a8e8316c21d70f6b.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LC#2", utils.IsLocal("../file/2013/09/14/1404994819612896a8e8316c21d70f6b"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LC#3", utils.IsLocal("/file/2013/09/14/1404994819612896a8e8316c21d70f6b."))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LC#4", utils.IsLocal("./file/2013/09/14/1404994819612896a8e8316c21d70f6b"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LC#5", utils.IsLocal("./file/2013/09/14/1404994819612896a8e8316c21d70f6b."))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LC#6", utils.IsLocal("./file/2013/09/14/1404994819612896a8e8316c21d70f6b.png"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LCWEB#1", utils.IsLocal(".http://bgimg1.meimei22.com/recpic/2013/27.jpg"))
	fmt.Println("----------------------------------------------------------------")
	fmt.Println("LCWEB#2", utils.IsLocal("http://t12.baidu.com/it/u=118735160,777859196&fm=120"))

}
