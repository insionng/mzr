package main

import (
	"fmt"
	"regexp"
	"veryhour/helper"
)

func main() {

	s, e := helper.GetPage("http://www.aimmxiu.com/")
	fmt.Println(e)
	m, n := GetImages(s)
	fmt.Println("匹配", n, "张图:")
	fmt.Println(m)
}

//  返回 图片url列表集合
func GetImages(content string) (imgs []string, num int) {

	//替换HTML的空白字符为空格
	ren := regexp.MustCompile(`\s`) //ns*r
	bodystr := ren.ReplaceAllString(content, " ")

	rex := regexp.MustCompile("((http://)+([^ rn()^$!\"|[]{}<>]*)((.gif)|(.jpg)|(.bmp)|(.png)|(.GIF)|(.JPG)|(.PNG)|(.BMP)))")
	img_urlz := rex.FindAllSubmatch([]byte(bodystr), -1)

	//匹配所有图片
	//rem := regexp.MustCompile(`<img.+?src="(.+?)".*?`) //匹配最前面的图
	//img_urls := rem.FindAllSubmatch([]byte(bodystr), -1)

	for _, bv := range img_urlz {
		if m := string(bv[1]); m != "" {
			imgs = append(imgs, m)
		}
	}

	return imgs, len(imgs)
}
