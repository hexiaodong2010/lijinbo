package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"lijinbo/config"
	"net/http"
	"os"
)

var (
	runMode string
)

func init() {
	flag.StringVar(&runMode, "runMode", "prod", "run mode such as `dev` `prod` `test`")
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `%s version: %s/%s
Usage: %s [-hvVtTq] [-s signal] [-c filename] [-p prefix] [-g directives]

Options:
`, config.AppName, config.AppName, config.Version, config.AppName)
		flag.PrintDefaults()
	}
}
func main() {
	flag.Parse()
	switch runMode {
	case config.DevEnv:
		gin.SetMode(gin.DebugMode)
	case config.TestEnv:
		gin.SetMode(gin.TestMode)
	case config.ProEnv, "":
		gin.SetMode(gin.ReleaseMode)
	default:
		panic("run mode unknown: " + runMode)
	}
	config.InitConfig(runMode)
	route := gin.Default()
	route.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "welcome!")
	})
	_ = route.Run()
}
