package main

import (
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/robertkowalski/graylog-golang"
	"github.com/satori/go.uuid"
)

// ToLogGelf GELF representation
type ToLogGelf struct {
	*ToLog
	Version      string `json:"version"`
	ShortMessage string `json:"short_message"`
	Line         int    `json:"line"`
	File         string `json:"file"`
}

// ToLog bidule à loguer
type ToLog struct {
	OvhToken     string `json:"_X-OVH-TOKEN"`
	TimeStamp    int64  `json:"timestamp"`
	MailID       string `json:"_mail_id"`
	ProtecmailID string `json:"_protecmail_id"`
	MailFrom     string `json:"_mail_from"`
	RcptTo       string `json:"_rcpt_to"`
	Size         int    `json:"_size"`
	ScanStatus   string `json:"_scanstatus"` // SPAM VIRUS CLEAN PH
	Msg          string `json:"full_message"`
	Host         string `json:"host"`
	Stage        string `json:"_stage"`
	Level        int    `json:"level"`
}

// Gelf serialize ToLog as GELF
func (t ToLog) Gelf() (string, error) {
	tGelf := ToLogGelf{
		&t,
		"1.1",
		"blabla",
		0,
		"",
	}
	j, err := json.Marshal(tGelf)
	if err != nil {
		return "", err
	}
	return string(j), nil
}

func main() {

	// Get ovh TOKEN
	ovhToken := os.Getenv("OVH_TOKEN")
	if ovhToken == "" {
		log.Fatalln("pas de de token en ENV")
	}

	//ovhToken := "ovhtoken"

	g := gelf.New(gelf.Config{
		GraylogHostname: "laas.runabove.com",
		GraylogPort:     2202,
		Connection:      "wan",
		MaxChunkSizeWan: 42,
		MaxChunkSizeLan: 1337,
	})

	for {

		msg := ToLog{
			OvhToken:   ovhToken,
			TimeStamp:  time.Now().Unix(),
			MailID:     GetRandString(),
			MailFrom:   "toorop@toorop.fr",
			RcptTo:     "toorop@tmail.io",
			Size:       25874,
			ScanStatus: "clean",
			Msg:        GetRandString(),
			Host:       "localhost",
			Stage:      GetRandString(),
			Level:      7,
		}

		out, err := msg.Gelf()
		if err != nil {
			log.Fatalln(err)
		}
		//log.Println(out)

		g.Send([]byte(out))
		log.Println("message envoyé")
		time.Sleep(10 * time.Millisecond)
	}

	//log.Println(err, "-", out)

}

func GetRandString() string {
	return uuid.NewV4().String()
}
