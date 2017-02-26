package main

import (
	"flag"
	"os"

	"github.com/abourget/slick"
	_ "github.com/jbarratt/awspricebot/priceplugin"
)

var configFile = flag.String("config", os.Getenv("HOME")+"/.awspricebot.conf", "config file")

func main() {
	flag.Parse()
	bot := slick.New(*configFile)
	bot.Run()
}
