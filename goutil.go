package goutil

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/PuerkitoBio/goquery"
)

type Data struct {
	Data []string
}

//从一个json文件读取数据，返回字符串数组，json文件格式如下：
//{"Data":["xx","xx","xx",...]}
func GetStringsFromJson(jsonName string) ([]string, error) {
	b, err := ioutil.ReadFile(jsonName)
	if err != nil {
		return nil, err
	}
	items := &Data{[]string{}}
	err = json.Unmarshal(b, items)
	if err != nil {
		return nil, err
	}
	return items.Data, nil

}

//判断文件是否存在
func IsExist(fileName string) bool {
	_, err := os.Stat(fileName)
	return err == nil || os.IsExist(err)
}

//获得一个网址的所有图片链接
func GetImagesLinks(url string) ([]string, error) {
	links := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find("img")

	for _, n := range sel.Nodes {

		for _, i := range n.Attr {
			if i.Key == "src" {
				links = append(links, i.Val)
			}
		}
	}
	return links, nil

}

//获得一个网址的所有超链接
func GetHyperlinks(url string) ([]string, error) {
	links := []string{}
	doc, err := goquery.NewDocument(url)
	if err != nil {
		return nil, err
	}
	sel := doc.Find("a")

	for _, n := range sel.Nodes {

		for _, i := range n.Attr {
			if i.Key == "href" {
				links = append(links, i.Val)
			}
		}
	}
	return links, nil

}
