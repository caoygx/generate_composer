package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"
)

var Root = "./"

func main() {

	var module string
	flag.StringVar(&module, "module", "news", "项目名称")
	flag.Parse()

	createTpl(module)

}

func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

func createDir(dirProject string) {

	//dirProject := "payment"
	saveDir := Root + "vendor/" + dirProject
	if _, err := os.Stat(saveDir); os.IsNotExist(err) {
		os.MkdirAll(saveDir, os.ModePerm)
	}

	var directories = [...]string{"controller", "model", "service", "validate"}
	for _, v := range directories {
		err := os.Mkdir(saveDir+"/"+v, os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		//fmt.Println(v)
	}
}

func createTpl(module string) {

	//module := "new"
	createDir(module)
	var directories = [...]string{"controller", "model", "service", "validate"}
	for _, v := range directories {
		content := getTpl("tpl" + "/" + v + ".php")
		content = replace(content, module)
		err := ioutil.WriteFile(Root+"vendor/"+module+"/"+v+"/"+module+".php", []byte(content), 0666)
		fmt.Println(err)
	}

}

func getTpl(filepath string) string {

	f, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}

	content, _ := ioutil.ReadAll(f)

	//fmt.Println(string(content))
	return string(content)
}

func replace(content string, module string) string {
	//content = strings.Replace(content, "{moduleUpper}", strings.ToUpper(module), -1)
	content = strings.Replace(content, "{moduleUpper}", Ucfirst(module), -1)
	content = strings.Replace(content, "{module}", strings.ToLower(module), -1)
	fmt.Println(content)
	return content
}
