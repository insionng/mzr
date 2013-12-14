package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {

	c, _ := GetPage("http://localhost/topic/1/")
	ma := GetImages(c)

	cc, _ := GetPage("http://localhost/topic/2/")
	mb := GetImages(cc)

	var mc, md []string
	var k []int

	//循环ma集合
	for i, av := range ma {

		//av对比mb集合里的每个元素,判断av是否存在mb中
		for j, bv := range mb {
			//若av存在mb中
			if av == bv {
				fmt.Println(i, "==", j)
				//记录av的id号i
				k = append(k, i)
				mc = append(mc, string(av))
			}
			fmt.Println(i, av, "::", j, bv)

		}
	}

	for k, cv := range mc {
		for i, av := range ma {
			if cv == av {

			}
		}
	}

	fmt.Println("交集:", mc)
	//假设 A是更新提交新内容  B是旧内容
	fmt.Println("差集:", md) //差集=(ASET -BSET )= ASET - 交集
}

func GetPage(url string) (string, error) {

	if resp, err := http.Get(url); err != nil {
		return "", err
	} else {

		if body, err := ioutil.ReadAll(resp.Body); err != nil {
			return "", err
		} else {
			return string(body), err
		}
	}

}

/*
 *  返回 图片url列表集合
 */
func GetImages(content string) (img_re []string) {

	//去掉HTML消息体的空白字符,替换为空格。
	ren := regexp.MustCompile(`\s`)
	bodystr := ren.ReplaceAllString(content, " ")

	//查找<img src="xxx" />间内容
	rem := regexp.MustCompile(`<img src="(.+?)"`)
	img_urls := rem.FindAllSubmatch([]byte(bodystr), -1)

	for _, img_url := range img_urls {
		img_re = append(img_re, string(img_url[1]))
	}

	return img_re
}
