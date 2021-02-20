package main

import (
	beego "github.com/beego/beego/v2/server/web"
	_ "p_accounting/routers"
)

func main() {
	beego.Run()
}
