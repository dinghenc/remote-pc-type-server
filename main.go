package main

import (
	"flag"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	var appPath string
	flag.StringVar(&appPath, "app", "./app", "app path")
	flag.Parse()
	if appPath == "" {
		log.Fatalln("app path is empty")
	}

	s := new(Server)

	r := gin.Default()
	r.POST("/server/onChangeText", s.OnChangeText)
	r.POST("/server/onSearch", s.OnSearch)
	r.Static("/", appPath)
	log.Fatalln(r.Run(":8272"))
}
