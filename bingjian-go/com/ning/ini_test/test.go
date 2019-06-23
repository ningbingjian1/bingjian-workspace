package main

import (
	"fmt"
	"github.com/tietang/props/ini"
	"github.com/tietang/props/kvs"
	"time"
)
func main() {
	file := kvs.GetCurrentFilePath("conf/config.ini",1)
	conf := ini.NewIniFileConfigSource(file)
	port := conf.GetIntDefault("app.server.port",10901)
	name := conf.GetDefault("app.name","unknow")
	enabled := conf.GetBoolDefault("app.enabled",false)
	duration := conf.GetDurationDefault("app.time",time.Second)
	fmt.Println(port)
	fmt.Println(name)
	fmt.Println(enabled)
	fmt.Println(duration)
}
