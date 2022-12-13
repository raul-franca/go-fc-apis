package main

import (
	"github.com/raul-franca/go-fc-apis/configs"
)

func main() {

	conf, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(conf.DBName)
}
