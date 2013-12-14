package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/ioutil"
	"reflect"
)

var home = `
{{Call .A "Right?"}}
`

var templateFuncs = template.FuncMap{"Call": Call}

func Call(method interface{}, args ...interface{}) interface{} {
	fun := reflect.ValueOf(method)
	params := make([]reflect.Value, 0, len(args))
	for _, arg := range args {
		params = append(params, reflect.ValueOf(arg))
	}
	values := fun.Call(params)
	if len(values) > 0 {
		value := values[0]
		return value.Interface()
	}
	return ""
}

func main() {
	A := func(name string) string {
		return "Hello," + name
	}

	t := template.New("home.html").Funcs(templateFuncs)
	t, _ = t.New("home.html").Parse(home)
	buf := bytes.NewBufferString("")
	t.ExecuteTemplate(buf, "home.html", map[string]interface{}{"A": A})
	con, _ := ioutil.ReadAll(buf)
	fmt.Println("--------------------")
	fmt.Println(string(con))
	fmt.Println("--------------------")
}
