package wolfx

import (
	"encoding/json"
	"fmt"
	"hayai/config"
	"hayai/constants"
	"hayai/seismo"
	"hayai/utils"
	"log"
	"net/url"
	"os"
	"time"

	"embed"

	"github.com/faiface/beep"
	"github.com/faiface/beep/flac"
	"github.com/faiface/beep/speaker"
	"github.com/gen2brain/beeep"
	"github.com/gorilla/websocket"
	"github.com/sqweek/dialog"
)

//go:embed assets/alertv3-sat.flac
var alertSoundFile embed.FS

type TypeMessage struct {
	Type string
}

type Issue struct {
	Source string
	Status string
}

type Accuracy struct {
	Epicenter string
	Depth     string
	Magnitude string
}

type MaxIntChange struct {
	String string
	Reason string
}

type WarnArea struct {
	Chiiki  string
	Shindo1 string
	Shindo2 string
	Time    string
	Type    string
	Arrive  bool
}

type JMAEEW struct {
	Type          string
	Title         string
	CodeType      string
	Issue         Issue
	EventID       string
	Serial        int
	AnnouncedTime string
	OriginTime    string
	Hypocenter    string
	Latitude      float64
	Longitude     float64
	Magunitude    float64
	Depth         int
	MaxIntensity  string
	Accuracy      Accuracy
	MaxIntChange  MaxIntChange
	WarnArea      []WarnArea
	IsSea         bool
	IsTraining    bool
	IsAssumption  bool
	IsWarn        bool
	IsFinal       bool
	IsCancel      bool
	OriginalText  string
	Pond          string
}

var LastRetry time.Time

func Listen() {
	// Init sound
	alertSound, err := alertSoundFile.Open("assets/alertv3-sat.flac")
	utils.CheckError(err)
	defer alertSound.Close()
	streamer, format, err := flac.Decode(alertSound)
	utils.CheckError(err)
	defer streamer.Close()
	// Listen
	u := url.URL{Scheme: "wss", Host: constants.WSHost, Path: "/all_eew"}
	for {
		log.Printf("connecting to %s", u.String())
		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			log.Print("dial:", err)
			time.Sleep(time.Duration(config.Config.RetryConnectionEveryXS) * time.Second)
			continue
		}

		defer c.Close()

		done := make(chan struct{})

		defer close(done)

		log.Printf("connected.")
		for {
			_, message, err := c.ReadMessage()
			utils.CheckError(err)
			if message == nil {
				log.Printf("connection lost.")
				if time.Since(LastRetry) < time.Duration(config.Config.RetryConnectionEveryXS)*time.Second {
					time.Sleep(time.Duration(config.Config.RetryConnectionEveryXS) * time.Second)
					LastRetry = time.Now()
				}
				break
			}
			var typeMessage TypeMessage
			if config.Config.TestWarning {
				message = constants.TestMessage
			}
			json.Unmarshal(message, &typeMessage)
			if typeMessage.Type == "jma_eew" {
				log.Printf("recv: %s", message)
				var jmaeew JMAEEW
				json.Unmarshal(message, &jmaeew)
				if !jmaeew.IsWarn && config.Config.OnlyWarnings {
					continue
				}
				equivalentMagnitude := seismo.CalculateEquivalentMagnitude(jmaeew.Magunitude, jmaeew.Latitude, jmaeew.Longitude, config.Config.Latitude, config.Config.Longitude)
				if config.Config.IssueWarningAtAnyMagnitude || config.Config.IssueWarningAtEquivalentMagnitude < equivalentMagnitude {
					warnareas := ""
					if len(jmaeew.WarnArea) != 0 {
						warnareas += "\nWarnings issued for:"
					}
					for _, warnarea := range jmaeew.WarnArea {
						warnareas += fmt.Sprintf("\n - %s (at %s JST) [Shindo %s / %s] [%s]", warnarea.Chiiki, warnarea.Time, warnarea.Shindo1, warnarea.Shindo2, warnarea.Type)
					}
					alert_body := fmt.Sprintf("Epicenter: %s\nMagnitude %0.1f\nApproximately magnitude %0.1f at your location%s\n\nStrong shaking is expected soon.\nStay calm and seek shelter nearby.\n\nOrigin time: %s JST\nAnnouncement time: %s JST\nDepth: %dkm\nCoordinates: %0.1f, %0.1f\n\nSource: %s\nStatus: %s\n\n%s", jmaeew.Hypocenter, jmaeew.Magunitude, equivalentMagnitude, warnareas, jmaeew.OriginTime, jmaeew.AnnouncedTime, jmaeew.Depth, jmaeew.Latitude, jmaeew.Longitude, jmaeew.Issue.Source, jmaeew.Issue.Status, jmaeew.OriginalText)
					if config.Config.IssuePopup {
						go dialog.Message("%s", alert_body).Title("Early Earthquake Warning!").Info()
					}

					if config.Config.IssueNotification {
						homedir, err := os.UserHomeDir()
						utils.CheckError(err)
						err = beeep.Notify("Early Earthquake Warning!", alert_body, homedir+"/.icons/actions/scalable/dialog-warning.svg")
						utils.CheckError(err)
					}

					if config.Config.OpenWebPages {
						utils.OpenURL(constants.OpenURLA)
						utils.OpenURL(constants.OpenURLB)
					}

					if config.Config.IssueWarningSound {
						speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
						streamer.Seek(0)
						donep := make(chan bool)
						speaker.Play(beep.Seq(streamer, beep.Callback(func() {
							donep <- true
						})))
						<-donep
						speaker.Close()
					}
				}
			}
		}
	}
}
