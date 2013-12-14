package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func main() {

	createtime := time.Now().UnixNano()
	s := "hgdfjhsdfhjsdhjfhjd-" + strconv.Itoa(int(createtime))
	p := strings.Split(s, "-")

	fmt.Println(p[0])
	fmt.Println(p[1])

}
