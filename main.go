package main

import (
	"fmt"
	_ "snatch_ssc/routers"

	"github.com/astaxie/beego"
)

func init() {
	fmt.Println("====init")

}

func main() {
	beego.AddTemplateExt("htm")
	beego.Run()
}
