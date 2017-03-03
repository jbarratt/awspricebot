package priceplugin

import (
	"fmt"
	"log"
	"regexp"

	"github.com/abourget/slick"
	"github.com/jbarratt/awsprice"
)

// PricePlugin stores the state for Price Plugin
type PricePlugin struct {
	pricer awsprice.Pricer
	bot    *slick.Bot
}

func init() {
	slick.RegisterPlugin(&PricePlugin{})
}

var reAtMention = regexp.MustCompile(`<@([A-Z0-9]+)(|([^>]+))>\s*`)

// InitPlugin initializes the plugin, including downloading AWS pricing data
// if needed
func (price *PricePlugin) InitPlugin(bot *slick.Bot) {
	price.bot = bot
	awsprice.FetchJSON()
	awsprice.ProcessJSON()
	pricer, err := awsprice.LoadPriceDB()
	if err != nil {
		panic(err)
	}
	price.pricer = pricer
	bot.Listen(&slick.Listener{
		MessageHandlerFunc: price.ChatHandler,
	})
}

// ChatHandler gets triggered when any matching messages come in
func (price *PricePlugin) ChatHandler(listen *slick.Listener, msg *slick.Message) {
	if msg.MentionsMe {

		log.Printf("Message: '%s'\n", msg.Text)
		payload := reAtMention.ReplaceAllString(msg.Text, "")
		log.Printf("After filter: '%s'\n", payload)
		value, err := awsprice.ParseInput(price.pricer, payload)

		if err != nil {
			msg.Reply("Unable to find a price for '%s'\n", payload)
		} else {
			msg.Reply(fmt.Sprintf("```\n%s```", value))
		}

	}
}
