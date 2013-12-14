package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	client := &http.Client{}

	request, err := http.NewRequest("Get", "http://veryhour.com/signup/", nil)
	if err != nil {
		log.Println(err)
	}

	resp, err := client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}

	var xsrf string
	func() {
		x := []rune(strings.Split(string(body), `name="_xsrf" value="`)[1])
		for _, v := range x {
			if v == []rune("\"")[0] {
				break
			}
			xsrf += string(v)
		}
	}()

	URL := url.Values{
		"_xsrf":      {xsrf},
		"email":      {"11@ss.cc"},
		"username":   {"zhuantou"},
		"password":   {"zhuantou"},
		"repassword": {"zhuantou"},
	}

	urlLink := "http://veryhour.com/signup/"
	urlStr, _ := url.Parse(urlLink)
	postData := ioutil.NopCloser(strings.NewReader(URL.Encode()))
	request, err = http.NewRequest("POST", urlStr.String(), postData)

	resp, err = client.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	log.Println(resp.StatusCode)
	log.Println(xsrf)
}
