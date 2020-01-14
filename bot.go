package shipbot

import (
	"github.com/nlopes/slack"
	"github.com/object88/shipbot/log"
)

type Bot struct {
	Slacktoken string
	Log        *log.Log
}

func (b *Bot) Run() {
	api := slack.New(b.Slacktoken)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	b.eventloop(rtm)
}

func (b *Bot) eventloop(rtm *slack.RTM) {
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			b.Log.Infof("Event received\n")
			switch ev := msg.Data.(type) {

			case *slack.RTMError:
				b.Log.Errorf("Error: %s\n", ev.Error())

			case *slack.InvalidAuthEvent:
				b.Log.Errorf("Invalid credentials")
				return

			default:
				// Take no action
			}
		}
	}
}
