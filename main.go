package main

import (
	"month-report/utils"
)

func main() {
	//json := config.Load("./page.json")
	utils.PDF("hello.pdf", "utils/sarasa-term-sc-regular.ttf")
}
