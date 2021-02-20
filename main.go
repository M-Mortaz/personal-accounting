package main

import (
	_ "p_accounting/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}

