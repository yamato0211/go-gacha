package main

import (
	adaptorHTTP "go-gacha-system/pkg/adapter/http"

	_ "github.com/go-sql-driver/mysql"
)

//	@title			GachaAPI
//	@version		1.0
//	@description	This is a simple gacha system api
//	@host			localhost:8080
func main() {
	r := adaptorHTTP.InitRouter()
	r.Run(":8080")
}
