package main

import (
	"github.com/elysiamori/finalproject1-msib5-hacktiv8/config"
	"github.com/elysiamori/finalproject1-msib5-hacktiv8/routers"
)

func main() {
	_, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	r := routers.StartRouter()
	r.Run(":3000")
}
