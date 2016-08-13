package administrative_divisions

import (
	"gopkg.in/orivil/orivil.v1"
	"path/filepath"
	"os"
	"github.com/PuerkitoBio/goquery"
	"strconv"
)

// 数据来源
var url = "http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201608/t20160809_1386477.html"

// 数据更新时间: 2016.08.09

// 数据文件，该文件从 url 地址扒来
var html = filepath.Join(orivil.DirBundle, "base", "data", "administrative-division.html")

// 省
var provinces []string

// 市
var cities map[string][]string

// 区县
var counties map[string]map[string][]string

func init() {
	cities = make(map[string][]string, 34)
	counties = make(map[string]map[string][]string, 34)

	// open static file
	f, err := os.Open(html)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// new html document
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}

	div := doc.Find("div")

	// 临时数据
	province := ""
	city := ""
	// 遍历 p 标签
	div.Find("p").Each(func(i int, p *goquery.Selection) {
		code := ""
		name := ""
		// 读取 p 标签中的两个 span 标签
		p.Find("span").Each(func(j int, span *goquery.Selection) {
			if code == "" {
				code = span.Text()
			} else {
				name = span.Text()
				c2, _ := strconv.Atoi(string(code[2:4]))
				c3, _ := strconv.Atoi(string(code[4:6]))

				if c2 == 0 && c3 == 0 { // 将当前数据视为省名

						province = name
						provinces = append(provinces, province)
				} else if c2 != 0 && c3 == 0 { // 将当前数据视为市名

					city = name
					cities[province] = append(cities[province], city)
				} else if c2 != 0 && c3 != 0 { // 将当前数据视为区县名

					if _, ok := counties[province]; !ok {
						counties[province] = map[string][]string{city:{name}}
					} else {
						counties[province][city] = append(counties[province][city], name)
					}
				}
			}
		})
	})
}

// 获得所有省名
func Getprovinces() []string {

	return provinces
}

// 根据省名获得市名
func GetCities(province string) []string {

	return cities[province]
}

// 根据市名获得区县名
func GetCounties(province, city string) []string {

	return counties[province][city]
}