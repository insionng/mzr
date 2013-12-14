package main

import (
	"fmt"
	"strings"
	"veryhour/helper"
)

var p = `户名 整的用户名，按照常规的@insio 表现形式，一般是以@开头，以：结尾，中间为用户的名称。
匹配表达式就可写为：   是在微博之后@http://www.84420.com/ 还有
下载并加入知乎，随时随地提问解惑分享知识，发现更大的世 知乎阅  • © 2013 @知乎`

func main() {
	/*
		p, err := helper.GetPage("http://www.8mei.cc")
		fmt.Println("P>>>>>>>>>>>", p, err)
	*/
	//model.AtLinksPostImagesOnTopic(p)
	values := []string{}
	links, _ := helper.AtPages(p)
	for k, v := range links {
		fmt.Println("link #", k)
		content, _ := helper.GetPage(v)
		imgs, n := helper.GetImages(content)

		if n > 0 {
			for k, vv := range imgs {
				//vv为单图url 相对路径的处理较为复杂,这里暂时抛弃相对路径的图片 后续再修正
				if strings.HasPrefix(vv, "/") {
					vv = v + vv[1:]
				}
				fmt.Println("image #", k)
				if !helper.ContainsSets(values, vv) {
					values = append(values, vv)
				}
			}
		}
	}
	fmt.Println("values:", values)

}
