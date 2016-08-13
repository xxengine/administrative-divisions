# 中国行政区划

## 介绍

最新中国行政区划，数据更新时间：2016.08.09。
数据来源：http://www.stats.gov.cn/tjsj/tjbz/xzqhdm/201608/t20160809_1386477.html

## 安装

go get -v github.com/orivil/administrative-divisions


> **Note:** 如果中国地区用户出现 godoc.org/golang.org/x/net 包不能下载的情况，可选择从 GitHub 下载(go get -v github.com/golang/net)，
然后将 $GOPATH/src/github.com/golang/net 目录复制到 $GOPATH/src/godoc.org/golang.org/x/net 目录下

## 示例

`main.go`:

```GO
package main

import (
	"github.com/orivil/administrative-divisions"
	"fmt"
)

func main() {
	provinces := administrative_divisions.Getprovinces()
	for _, province := range provinces {
		// 打印省名
		fmt.Printf("%s\n", province)

		// 根据省名获得市名
		cities := administrative_divisions.GetCities(province)
		for _, city := range cities {
			// 打印市名
			fmt.Printf("  %s\n", city)

			// 根据省名及市名获得区县名
			counties := administrative_divisions.GetCounties(province, city)

			// 打印区县
			for _, county := range counties {
				fmt.Printf("    %s\n", county)
			}
		}
	}
}
```

## Contributors

https://github.com/orivil/administrative-divisions/graphs/contributors

## License

Released under the [MIT License](https://github.com/orivil/administrative-divisions/blob/master/LICENSE).