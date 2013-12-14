package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"veryhour/helper"
)

func main() {

	fmt.Println(GetImagesFromUrl("http://www.fengyun5.com/"))

}

func GetImagesFromUrl(url string) (urls []string) {

	//c, _ := helper.GetPage(url)
	x, _ := goquery.NewDocumentFromReader(strings.NewReader(`<a href="sfdsfd"></a>`))
	s := x.First()
	for {
		v, b := s.Attr("href")
		fmt.Println(v, b)
		if b == true {
			if !helper.ContainsSets(urls, v) {
				urls = append(urls, v)
			}
		}

		s = s.Next()
		if _, b := s.Attr("href"); b == false {
			break
		}
	}
	return urls
}

func Downloader(urls []string) {
	if len(urls) > 0 {
		fmt.Println("urlz:", len(urls))
		for k, v := range urls {

			res, _ := http.Get(v)
			file, _ := os.Create(strconv.Itoa(int(k)) + ".jpg")
			io.Copy(file, res.Body)
		}
	}
	fmt.Println("done!")
}
