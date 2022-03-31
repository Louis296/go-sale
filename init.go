package main

import (
	"github.com/gin-gonic/gin"
	"github.com/louis296/go-sale/conf"
	"github.com/louis296/go-sale/dao"
	"github.com/louis296/go-sale/handler"
	"github.com/louis296/go-sale/pkg/jwt"
	"strconv"
)

func Init(r *gin.Engine) {

	r.Any("/v1", handler.MainHandler)

	// init configure
	configure, err := conf.GetConf()
	if err != nil {
		panic(err)
	}

	// init database
	err = dao.InitDB(
		configure.Database.URL,
		configure.Database.DatabaseName,
		configure.Database.UserName,
		configure.Database.Password,
		configure.Database.MaxConn,
		configure.Database.MaxOpen,
	)
	if err != nil {
		panic(err)
	}

	// init jwt secret
	jwt.Secret = configure.Jwt.Secret

	err = r.Run(":" + strconv.Itoa(configure.Server.Port))
	if err != nil {
		panic(err)
	}
}
