package shipbot

import (
	"strings"

	"github.com/nlopes/slack"
	"github.com/object88/shipbot/log"
)

type Bot struct {
	Slacktoken string
	Log        *log.Log
}

func (b *Bot) Run() {
	if b.Slacktoken == "" {
		return
	}
	api := slack.New(b.Slacktoken)
	rtm := api.NewRTM()
	go rtm.ManageConnection()

	b.eventloop(rtm)
}

func (b *Bot) eventloop(rtm *slack.RTM) {
	for {
		select {
		case msg := <-rtm.IncomingEvents:
			b.Log.Infof("Event received:\n\t%#v\n", msg)

			switch ev := msg.Data.(type) {
			case *slack.MessageEvent:
				info := rtm.GetInfo()

				text := ev.Text
				text = strings.TrimSpace(text)
				text = strings.ToLower(text)
				b.Log.Infof("Message: %s\n", text)

				if ev.User != info.User.ID {
					rtm.SendMessage(rtm.NewOutgoingMessage("\\[T]/ Praise the Sun \\[T]/", ev.Channel))
				}

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
