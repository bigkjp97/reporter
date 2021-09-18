package utils

import (
	"github.com/signintech/gopdf"
	"github.com/tidwall/gjson"
	"log"
)

/**
遍历获取列表里的每一条json
*/
func list(j string, n string) gjson.Result {
	//var l []string
	r := gjson.Get(j, n)
	//r.ForEach(func(key, value gjson.Result) bool {
	//	fmt.Println(value.String())
	//	l = append(l, value.String())
	//	return true
	//})
	//return l
	return r
}

func PDF(pdf string, font string) {
	file := gopdf.GoPdf{}
	file.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	file.AddPage()
	err := file.AddTTFFont("sarasa", font)
	if err != nil {
		log.Printf("字体文件加载异常 [Error]: %s\n\n", err.Error())
	}

	err = file.SetFont("sarasa", "", 14)
	if err != nil {
		log.Printf("字体文件加载异常 [Error]: %s\n\n", err.Error())
	}

	file.Cell(nil, "hello")
	file.SetLineWidth(2)
	file.SetLineType("dashed")
	file.Line(10, 30, 585, 30)
	file.Cell(nil, "test")
	file.WritePdf(pdf)
}
