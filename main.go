package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/bnch/lan/dispatcher"
	"github.com/bnch/lan/handler"
	"github.com/bnch/lan/handler/banchobot"
	"github.com/fatih/color"
	"github.com/thehowl/conf"
	"gopkg.in/redis.v5"
)

const version = "0.0.1"

var c = struct {
	Port          string
	AdminPassword string `description:"The plaintext password of the admin user. Make this very long! If empty, it will be automatically generated."`
	RedisAddr     string
}{
	Port: ":80",
}

func main() {
	conf.Load(&c, "lan.conf")
	color.Yellow("lan v%s", version)

	if c.AdminPassword == "" {
		c.AdminPassword = handler.RandString(20)
		fmt.Println("Admin password was not found. Set to", c.AdminPassword)
	}

	handler.Redis = redis.NewClient(&redis.Options{
		Addr: c.RedisAddr,
	})

	handler.AdminPassword = fmt.Sprintf("%x", md5.Sum([]byte(c.AdminPassword)))
	conf.Export(&c, "lan.conf")

	banchobot.Start()
	handler.Start()

	color.Green("Listening on %s", c.Port)
	panic(http.ListenAndServe(c.Port, dispatcher.Server{}))
}
