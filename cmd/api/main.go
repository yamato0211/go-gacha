package main

import (
	adaptorHTTP "go-gacha-system/pkg/adapter/http"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	r := adaptorHTTP.InitRouter()
	r.Run(":8080")
}
