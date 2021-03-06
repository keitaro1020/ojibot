package ojibot

import (
	"bytes"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/greymd/ojichat/generator"
	"github.com/nlopes/slack"
	"github.com/nlopes/slack/slackevents"
)

// Ojibot is handler for cloud function
func Ojibot(w http.ResponseWriter, r *http.Request) {

	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()
	events, err := slackevents.ParseEvent(
		json.RawMessage(body),
		slackevents.OptionVerifyToken(
			&slackevents.TokenComparator{
				VerificationToken: os.Getenv("SLACK_VERIFICATION_TOKEN"),
			},
		))
	if err != nil {
		log.Printf("error: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	log.Printf("eventType: %v", events.Type)
	switch events.Type {
	case slackevents.URLVerification:
		var r *slackevents.ChallengeResponse
		err := json.Unmarshal([]byte(body), &r)
		if err != nil {
			log.Printf("error: %v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text")
		w.Write([]byte(r.Challenge))
		return

	case slackevents.CallbackEvent:
		token := os.Getenv("SLACK_TOKEN")
		api := slack.New(token)

		innerEvent := events.InnerEvent
		switch ev := innerEvent.Data.(type) {
		case *slackevents.AppMentionEvent:
			mev := innerEvent.Data.(*slackevents.AppMentionEvent)
			user, err := api.GetUserInfo(mev.User)
			if err != nil {
				log.Printf("error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			tname := user.Profile.DisplayName
			if len(tname) == 0 {
				tname = user.Name
			}

			en := rand.Intn(5)
			if en < 1 {
				en = 1
			}
			pn := rand.Intn(3)

			message, err := generator.Start(generator.Config{
				TargetName:        tname,
				EmojiNum:          en,
				PunctiuationLebel: pn,
			})
			if err != nil {
				log.Printf("error: %v", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			api.PostMessage(ev.Channel, slack.MsgOptionText(message, false))
		}
	}
}
